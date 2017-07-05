// cmdb_client_windows
package main

import (
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/cpu"
	//	"github.com/shirou/gopsutil/mem"
)

func collectCpu() map[string]string {
	var cpuInfo map[string]string
	cpuInfo = make(map[string]string)

	cpuNum, err := cpu.Counts(false)
	if err != nil {
		fmt.Println("collect Cpu count happened a error:", err)
	}
	cpuInfo["cpu_count"] = strconv.Itoa(cpuNum)
	return cpuInfo
}

//func collectDisk() map[string]string {

//}

//func collectRam() map[string]string {

//}

//func collectNic() map[string]string {

//}

//func collectHost() map[string]string {

//}

func main() {
	fmt.Println(collectCpu())
}
