// cmdb_client_windows
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/toolkits/file"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	asset_type              = "server"
	server                  = "192.168.5.100"
	port                    = 80
	request_timeout         = 30                               //发起http请求的超时时间
	asset_report_with_no_id = "/cmdb/report_with_no_asset_id/" // 第一次汇报的uRL
	asset_report            = "/cmdb/report/"                  // 第二次汇报的地址
	asset_id                = `c:\cmdb\asset_id`               // 资产ID
	user                    = "jiafa.liao@quanshi.com"         // 汇报的用户
	token                   = "todayisgoodtoday"               // 汇报的token
)

var (
	collectInfo                             map[string]interface{}
	reportAssetWithoutId, reportAssetWithId string
	isCollect                               string
	reportData                              string
	responseInfo                            map[string]interface{}
	tt                                      string
)

func collectCpu() bool {
	infoStat, err := cpu.Info()
	if err != nil { // 在虚拟机下面有时候用这个模块采集不到CPU信息，那么我们就使用命令来采集
		log.Println("collecting Cpu info happened a error:", err)
		log.Println("begin to collect CPU info by [ wmic cpu list brief ]")
		// 因为这个wmic命令直接使用wmic cpu list brief是无法使用的在go里面。
		// 所以我们需要和她交互下才可以使用这个命令
		cpucmd := "cpu list brief\n"
		a := exec.Command("wmic")
		inpipe, err := a.StdinPipe()
		if err != nil {
			log.Fatalf("executing command wmic happend a error:", err)
			return false
		}

		go func() {
			defer inpipe.Close()
			io.WriteString(inpipe, cpucmd)
		}()

		out, err := a.CombinedOutput()
		if err != nil {
			log.Println("collecting CPU info by wmic happend a error", err)
			return false
		}
		rt := string(out)
		strRt := strings.Split(rt, "\n")[1:]

		for i, line := range strRt {
			collectInfo["cpu_model"] = strings.Fields(line)[4]
			collectInfo["cpu_count"] = i
		}
		return true
	}
	if len(infoStat) <= 1 {
		for _, m := range infoStat {
			collectInfo["cpu_model"] = m.ModelName
			collectInfo["cpu_count"] = int(m.CPU) + 1
			collectInfo["cpu_core_count"] = m.Cores
		}
	} else {
		for _, m := range infoStat {
			collectInfo["cpu_model"] = m.ModelName
			v, ok := collectInfo["cpu_count"]
			if ok {
				core_v, okcore := collectInfo["cpu_core_count"]
				if okcore {
					collectInfo["cpu_core_count"] = int(m.Cores) + int(core_v.(int32)) // 把cpu核心相加
				} else {
					collectInfo["cpu_core_count"] = m.Cores
				}
				collectInfo["cpu_count"] = int(m.CPU) + int(v.(int32))
			} else {
				collectInfo["cpu_core_count"] = m.Cores
				collectInfo["cpu_count"] = int(m.CPU) + 1
			}
		}
	}
	return true
}

func collectDisk() {
	part, err := disk.Partitions(true)
	if err != nil {
		log.Println("collecting disk info happend a error:", err)
		return
	}
	var disk_info_array []interface{}

	for _, d := range part {
		var disk_map map[string]string
		disk_map = make(map[string]string)
		disk_map["name"] = d.Device
		info, err := disk.Usage(d.Device)
		if err != nil {
			log.Fatalf("collecting the %s happend a error:%s\n ", d.Device, err)
			continue
		}
		disk_map["capacity"] = strconv.Itoa(int(info.Total / (1024 * 1024 * 1024))) // 字节，需要换算到GB
		disk_map["sn"] = "None"
		disk_info_array = append(disk_info_array, disk_map)
	}
	collectInfo["local_disk_driver"] = disk_info_array
	var physical_disk []string
	collectInfo["physical_disk_driver"] = physical_disk
}

func collectRam() {
	/*
		采集内存的方法，由于模块返回的只有一个总内存，那么所以我们就按照一根内存条来处理了
	*/
	virtualMem, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalln("collect ram info happend a error: ", err)
	}
	var ramInfo []interface{}
	var perRamInfo map[string]string
	perRamInfo = make(map[string]string)

	perRamInfo["slot"] = "RAMslot#0"
	perRamInfo["manufactory"] = "NotSpecified"
	perRamInfo["asset_tag"] = "NotSpecified"
	perRamInfo["sn"] = "NotSpecified"
	perRamInfo["model"] = "DRAM"
	perRamInfo["capacity"] = strconv.Itoa(int(virtualMem.Total / (1024 * 1024))) // 字节转为兆
	ramInfo = append(ramInfo, perRamInfo)
	collectInfo["ram"] = ramInfo
}

func collectNic() {
	ifstat, err := net.Interfaces()
	if err != nil {
		log.Fatalln("collect nic info happend a error:", err)
	}
	var nic_array []interface{}

	for _, i := range ifstat {
		var nic_map map[string]string
		nic_map = make(map[string]string)
		nic_map["network"] = "NotSpecified"
		nic_map["bonding"] = "0"
		nic_map["mac_address"] = i.HardwareAddr
		nic_map["model"] = "unknow"
		for _, ip := range i.Addrs {
			if strings.Count(ip.Addr, ".") == 3 { // 意味着是IPv4
				if !strings.Contains(i.Name, "Loopback") { //回环地址去掉
					ip_netmask := strings.Split(ip.Addr, "/")
					nic_map["name"] = i.Name
					nic_map["ip_address"] = ip_netmask[0]
					// 下面两行是取子网掩码然后转换的
					netmask_num := ip_netmask[1]
					n, _ := strconv.Atoi(netmask_num)
					netmask := calcNetmask(n)
					nic_map["netmask"] = netmask
					nic_array = append(nic_array, nic_map)
				}
			}
		}
	}
	collectInfo["nic"] = nic_array
}

func Gbk2Utf(s string) (string, error) {
	/* 由gbk ---> utf8  */
	bs := []byte(s)
	reader := transform.NewReader(bytes.NewReader(bs), simplifiedchinese.GBK.NewDecoder())
	res, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func Utf2Gbk(s string) (string, error) {
	/* utf8 ---> gbk  */
	bs := []byte(s)
	reader := transform.NewReader(bytes.NewReader(bs), simplifiedchinese.GBK.NewEncoder())
	res, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func collectHost() {
	h, err := host.Info()
	if err != nil {
		log.Fatalln("collecting host info happend a error :", err)
		return
	}
	collectInfo["hostname"] = h.Hostname
	collectInfo["os_distribution"] = h.OS
	collectInfo["os_type"] = h.OS
	collectInfo["os_release"] = h.Platform
	collectInfo["kernel_release"] = h.PlatformVersion
	collectInfo["uuid"] = h.HostID
	sncmd := "bios get serialnumber\n"
	a := exec.Command("wmic")
	inpipe, err := a.StdinPipe()
	if err != nil {
		log.Println("executing command wmic happend a error:", err)
	}

	go func() {
		defer inpipe.Close()
		io.WriteString(inpipe, sncmd)
	}()

	out, err := a.CombinedOutput()
	c := strings.Split(string(out), "SerialNumber")
	c1 := c[1]
	d := strings.Split(c1, "\n")
	collectInfo["sn"] = strings.TrimSpace(d[1])

}

func addStr(num int, substring string) (str string) {
	/* 拼接字符串的！ */
	for i := 0; i < num; i++ {
		str = str + substring
	}
	return
}

func calcNetmask(netmask int) (return_netmask string) {
	/*
	   计算子网掩码，由数字类型转为xx.xxx.xx.xx
	   算法是这样的： 得到的数字先除以8，得到的商就是有多少个255，余数就需要再计算，余数等于1，那么最后一位子网眼码就是2**(8-1)，8是一段子网眼码长度，
	   为8个1，1111111,二进制计算。最后一段眼码计算方式如下：
	   余数为1，即2**7，
	   余数为2，即2**7+2**6
	   余数为3，即2**7+2**6+2**5
	   依次类推
	   :param netmask: tmp netmask
	   :return:
	*/
	factor := netmask / 8    // 商
	remainder := netmask % 8 // 余数
	mi := 8 - remainder      //  计算这个数字的幂
	var tmp_last_mask float64
	for mi <= 7 { // 判断掩码长度是否超过了7，超过了长度跳出循环，因为掩码的长度最多是8.
		tmp_last_mask = math.Pow(float64(2), float64(mi)) + tmp_last_mask
		mi = mi + 1
	}
	last_mask := int(tmp_last_mask)
	switch {
	case factor == 1 && last_mask == 0: // 意味着是8位的子网掩码
		return_netmask = fmt.Sprintf("%s%s", addStr(factor, "255."), "0.0.0")
	case (factor == 1 && last_mask != 0): //意味着8-24之间的子网掩码
		return_netmask = fmt.Sprintf("%s%d.%s", addStr(factor, "255."), last_mask, "0.0")
	case (factor == 2 && last_mask == 0): //意味着16位整的子网掩码
		return_netmask = fmt.Sprintf("%s0.0", addStr(factor, "255."))
	case (factor == 3 && last_mask == 0): // 意味着24位整的掩码
		return_netmask = fmt.Sprintf("%s%s", addStr(factor, "255."), "0")
	case (factor == 3 && last_mask != 0): // 意味着24-32之间的子网掩码
		return_netmask = fmt.Sprintf("%s%d", addStr(factor, "255."), last_mask)
	case (factor == 2 && last_mask != 0): // 意味着16-24位之间子网掩码
		return_netmask = fmt.Sprintf("%s%d.%s", addStr(factor, "255."), last_mask, "0")
	case factor == 4: // 意味着4个255
		return_netmask = fmt.Sprintf("%s255", addStr(factor-1, "255."))
	case factor == 0: // 小于8位的掩码
		return_netmask = fmt.Sprintf("%d.0.0.0", last_mask)
	}
	return return_netmask
}

func collectOther() {
	/* 采集工厂，型号，CPU等信息*/
	flag := collectCpu() // 之所以放在这里是因为2003系统兼容性不好，采集CPU信息挺多问题的
	a := exec.Command("systeminfo")
	b, _ := a.Output()
	bb := string(b)
	sb, _ := Gbk2Utf(bb)
	result := strings.Split(sb, "\n")
	for _, line := range result {
		switch {
		case strings.HasPrefix(line, "系统制造商:") || strings.HasPrefix(line, "System Model:"):
			fa := strings.Split(line, ":")[1]
			fatory := strings.TrimSpace(fa)
			collectInfo["manufactory"] = fatory
		case strings.HasPrefix(line, "系统型号:") || strings.HasPrefix(line, "System Manufacturer:"):
			mod := strings.Split(line, ":")[1]
			model := strings.TrimSpace(mod)
			collectInfo["model"] = model

		case !flag && strings.HasPrefix(line, "处理器:"):
			c := strings.Split(line, ":")[1]
			collectInfo["cpu_count"] = strings.Fields(c)[1]
			collectInfo["cpu_core_count"] = 1
			collectInfo["cpu_model"] = "unknow(go)"

		}
	}
}

func CollectAllData() string {
	assetId, _ := getAssetId()
	collectInfo = make(map[string]interface{})
	collectInfo["wake_up_type"] = "Power Switch"
	collectInfo["asset_type"] = asset_type
	collectInfo["asset_type"] = asset_type
	if len(assetId) > 0 {
		collectInfo["asset_id"] = assetId
	} else {
		collectInfo["asset_id"] = nil
	}
	collectHost()
	collectDisk()
	collectRam()
	collectNic()
	collectOther()
	//	var postData map[string]map[string]interface{}
	//	postData = make(map[string]map[string]interface{})
	//	postData["asset_data"] = make(map[string]interface{})
	//	postData["asset_data"] = collectInfo
	jsondata, err := json.Marshal(collectInfo)
	if err != nil {
		log.Fatalln("encoding json format data happend a error:", err)
		os.Exit(0)
	}
	return string(jsondata)
}

func connectCmdbServer(url, method string, jsondata string) (result []byte) {
	// 链接到cmdb服务器的，然后把采集到的信息提交上去的
	if method == "post" {
		//		res, err := http.Post(url, )
		//		postData := bytes.NewReader(jsondata)
		fmt.Printf("begin to post data to this url[%s]", url)
		resp, err := http.Post(url,
			"application/x-www-form-urlencoded",
			strings.NewReader(fmt.Sprintf("asset_data=%s", jsondata)))
		if err != nil {
			fmt.Println("connect cmdb server was error:", err)
			return
		}
		defer resp.Body.Close()
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ioutil read happend a error:", err)
		}
		return result
	}
	return
}

func ReportData() {
	/* 汇报资产数据 */
	jsondata := CollectAllData()
	responseInfo = make(map[string]interface{})
	fmt.Println(jsondata)
	if len(collectInfo) > 0 {
		assetId, _ := getAssetId()
		fmt.Println(len(assetId), assetId)
		if len(assetId) > 0 {
			method := "post"
			result := connectCmdbServer(reportAssetWithId, method, jsondata)
			//json.Unmarshal(result, &responseInfo)
			fmt.Println(string(result))
		} else {
			method := "post"
			result := connectCmdbServer(reportAssetWithoutId, method, jsondata)
			json.Unmarshal(result, &responseInfo)
			v, ok := responseInfo["asset_id"]
			if ok {
				if saveAssetId(v) {
					log.Println("already saved asset id :", responseInfo["asset_id"])
				}
			} else {
				fmt.Println(string(result))
			}
		}
	} else {
		log.Fatalln("No Data after collecting, so won't report ")
	}
}

func saveAssetId(assetId interface{}) bool {
	/* 保存资产ID */
	fmt.Println("assetId", assetId)
	fileObj, err := os.OpenFile(asset_id, os.O_WRONLY|os.O_CREATE, 666)
	if err != nil {
		fmt.Println("open file happend a error:", err)
		return false
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	d, ok := assetId.(float64)
	if ok == false {
		log.Println("can't  convert the asset id from interface to float64")
		return false
	}
	writer.WriteString(strconv.Itoa(int(d)))
	writer.Flush()
	return true
}

func getAssetId() (string, error) {
	// 从文件里获取资产ID的
	if !file.IsExist(asset_id) {
		errMsg := fmt.Sprintf("not found the asset_id(%s)", asset_id)
		log.Printf(errMsg)
		return "", errors.New(errMsg)
	}
	assetIdFileObj, err := os.Open(asset_id)
	if err != nil {
		errMsg := fmt.Sprintf("open asset id(%s) happend a error:%s", asset_id, err)
		log.Printf(errMsg)
		return "", errors.New(errMsg)
	}
	defer assetIdFileObj.Close()
	assetIdFileReader := bufio.NewReader(assetIdFileObj)
	for {
		assetId, readErr := assetIdFileReader.ReadString('\n')
		if readErr == io.EOF {
			return assetId, nil
		}
		return assetId, nil
	}
	return "", errors.New("not asset id return")
}

func main() {
	reportAssetWithoutId = fmt.Sprintf("http://%s:%d%s", server, port, asset_report_with_no_id)
	reportAssetWithId = fmt.Sprintf("http://%s:%d%s", server, port, asset_report)
	flag.StringVar(&isCollect, "c", "", "collect the machine's data")
	flag.StringVar(&reportData, "r", "", "report the machine's data(include collect data)")
	switch {
	case len(isCollect) != 0:
		CollectAllData()
	case len(reportData) != 0:
		ReportData()
	default:
		ReportData()
	}
	fmt.Scanln(&tt)
}
