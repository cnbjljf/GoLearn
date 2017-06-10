// handle_data
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var p chan map[string]map[string]int // saving data of timepoint and value
var timeList chan []string           // saving data of timepoint

// saveRetList and saveRetMap  DataValueMap是用来格式化为json数据的,
type saveRetList struct {
	DataList []saveRetMap
}

type saveRetMap struct {
	DataName  string
	DataValue []DataValueMap
}

type DataValueMap struct {
	Datetime string
	Value    int
}

type record struct {
	// 这个struct必须和你要从mysql取值出来的字段和数值类型一致
	timepoint  string
	metervalue int
}

func saveData(saveRet []map[string]map[string]int) int64 {
	// save data to redis

	//formating these data as json format
	var tJS []DataValueMap
	//	var resultList saveRetList
	var retMapList []saveRetMap

	for _, v := range saveRet {
		for name, v1 := range v {
			//			DataValueMap{v1[}
			for k, v := range v1 {
				jsData1 := DataValueMap{k, v}
				tJS = append(tJS, jsData1)
			}
			r := saveRetMap{name, tJS}
			retMapList = append(retMapList, r)
		}
	}
	resultList := saveRetList{retMapList}
	jsSaveData, _ := json.Marshal(resultList)

	// begin to save data to redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln("Happend a error when connecting,", err)
	}
	conn.Do("AUTH", "password")

	rt, err := conn.Do("SET", "JsonSaveData", jsSaveData)
	if rt == "OK" {
		return 1
	} else if err != nil {
		return 0
	}
	return 0
}

func handleData(sql_name, sql_cmd string, timeOffset int, p chan map[string]map[string]int, timeList chan []string) {
	// 从mysql取数据后进行清洗，根据时间跨度区间有三种不同的清洗方式，时间跨度为天
	/* <=5:按5分钟间隔处理
	5<x<25:按照1小时间隔处理
	>25：按照一天的间隔处理
	其中按照五分种间隔处理的清洗方式较多一步，需要把该时间点不存在的数值通过前后两个值的平均数给凑上
	*/
	var recordDict map[string]int
	recordDict = make(map[string]int, 100000)
	var timepointList []string

	//	log.Println("begin to connenct db")
	db, err := sql.Open("mysql", "username:password@tcp(host:port)/automatic_new")
	if err != nil {
		log.Fatal("can't connect mysql db", err)
	}
	//	log.Println("connected db was successfully!!")
	row, err := db.Query(sql_cmd)
	if err != nil {
		log.Fatal("sql command(", sql_cmd, ") is error ", err)
	}
	if row == nil {
		log.Fatal("no data from the db")
	} else {
		if timeOffset <= 5 {
			for i := 0; row.Next(); i++ {
				var r record
				row.Scan(&r.timepoint, &r.metervalue)
				value, ok := recordDict[r.timepoint]
				if ok == true { //means the map has the key ,so we need add the value
					recordDict[r.timepoint] = value + r.metervalue
				} else {
					recordDict[r.timepoint] = r.metervalue
					timepointList = append(timepointList, r.timepoint)
				}
			}
		}

		if 25 > timeOffset && timeOffset > 5 {
			for i := 0; row.Next(); i++ {
				var r record
				row.Scan(&r.timepoint, &r.metervalue)
				timepoint := strings.Split(r.timepoint, ":")[0:1]
				datetime := ""
				for _, i := range timepoint {
					datetime = datetime + i
				}
				value, ok := recordDict[datetime]
				if ok == true { // means the map has the key
					recordDict[datetime] = value + r.metervalue
				} else {

					recordDict[datetime] = r.metervalue // according to a hour as time Point
					//					timepointList = append(timepointList, datetime)
				}
			}
		}
		if timeOffset >= 25 {
			for i := 0; row.Next(); i++ {
				var r record
				row.Scan(&r.timepoint, &r.metervalue)
				timepoint := strings.Split(r.timepoint, " ")[0]
				value, ok := recordDict[timepoint]
				if ok == true {
					recordDict[timepoint] = value + r.metervalue
				} else {
					recordDict[timepoint] = r.metervalue
					//					timepointList = append(timepointList, timepoint)
				}
			}
		}
	}
	defer db.Close()
	sort.Strings(timepointList)
	timeList <- timepointList

	var dataDict map[string]map[string]int // 数据格式为 {"数据名字":{"时间点":数值}}
	dataDict = make(map[string]map[string]int, 10)
	dataDict[sql_name] = recordDict

	p <- dataDict

}

func getDateData(dict map[string]map[string]int, timeList []string, p chan map[string]map[string]int) {
	// this is first time to clear the data of date ,because these data they have different datetime

	for k, _ := range dict { // k is a name of select sql

		for _, value := range timeList { // 循环这个时间数组
			_, ok := dict[k][value] // 如果循环出来的时间不等于tTmp指定下标的时间,那么就绪要把dict里这个时间点的数据进行清洗,也就是把这个值设置为0
			if ok == false {
				dict[k][value] = 0 // No data,so set the key's Value as 0
			}
		}
		p <- dict
	}

}

func clearData(timeList chan []string, len_sql_list int) []string {
	// 把时间去重，返回一个去重后的时间轴
	var t []string
	for i := 0; i < len_sql_list; i++ {
		rt := <-timeList // 把time管道里面的值取出来,一个select语句对应一个结果
		for _, v := range rt {
			t = append(t, v) // 把时间追加到这个列表上去
		}
	}
	sort.Strings(t) // 排序时间
	a_len := len(t)
	var ret []string
	for i := 0; i < a_len; i++ { // 循环这个 时间数组，如前后两个值不相等或者这个值不为空，那么就添加到ret数组里，实现去重的效果
		if (i > 0 && t[i-1] == t[i]) || len(t[i]) == 0 {
			continue
		}
		ret = append(ret, t[i])
	}

	return ret
}

//export mainRun
func mainRun(timeOffset int, start_time, end_time string) int64 {
	voip_select_sql := " select timepoint,metervalue from tb_meter where category like '%-voip' and meterkey='minutes' and timepoint>='" + start_time + " 00:00:00' and timepoint<='" + end_time + " 23:59:59' "
	pstn_select_sql := " select timepoint,metervalue from tb_meter where category like 'summit_%' and meterkey='minutes' and timepoint>='" + start_time + " 00:00:00' and timepoint<='" + end_time + " 23:59:59' "
	meet_number_select_sql := " select timepoint,metervalue from tb_meter where ( category like 'tang4_%' or category like 'tangpc2_2%' or category like 'tang_econf%' or  category like 'tang3_%' ) and meterkey='conferences' and timepoint>='" + start_time + " 00:00:00' and timepoint<='" + end_time + " 23:59:59' "
	people_number_select_sql := " select timepoint,metervalue from tb_meter where ( category like 'tang4_%' or category like 'tangpc2_2%' or category like 'tang_econf%' or  category like 'tang3_%' ) and meterkey='times' and timepoint>='" + start_time + " 00:00:00' and timepoint<='" + end_time + " 23:59:59' "
	conn_number_select_sql := " select timepoint,metervalue from tb_meter where ( category like 'tang4_%' or category like 'tangpc2_2%' or category like 'tang_econf%' or  category like 'tang3_%' ) and meterkey='concurrent' and timepoint>='" + start_time + " 00:00:00' and timepoint<='" + end_time + " 23:59:59' "

	sql_cmd_list := map[string]string{"voip_data": voip_select_sql, "pstn_data": pstn_select_sql, "meet_data": meet_number_select_sql,
		"people_data": people_number_select_sql, "conn_data": conn_number_select_sql}

	p = make(chan map[string]map[string]int, 10) // this channel is use for saving data after first process
	timeList = make(chan []string, 100000)       // this channel is use for saving time point after first process
	len_sql_list := len(sql_cmd_list)

	for sql_name, sql_cmd := range sql_cmd_list {
		go handleData(sql_name, sql_cmd, timeOffset, p, timeList)
	}

	if timeOffset <= 5 {
		var pipe chan map[string]map[string]int // 这个channel用来存储第二次清洗后的数据，两个channel来存放是避免数据之间互相干扰
		pipe = make(chan map[string]map[string]int, 10)
		ret := clearData(timeList, len_sql_list) // 清洗的时间轴

		for i := 0; i < len_sql_list; i++ {
			v := <-p
			go getDateData(v, ret, pipe)
		}
		s1 := <-pipe
		s2 := <-pipe
		s3 := <-pipe
		s4 := <-pipe
		s5 := <-pipe

		saveRet := []map[string]map[string]int{s1, s2, s3, s4, s5}
		ok := saveData(saveRet)
		return ok

	} else {
		s1 := <-p
		s2 := <-p
		s3 := <-p
		s4 := <-p
		s5 := <-p
		saveRet := []map[string]map[string]int{s1, s2, s3, s4, s5}
		ok := saveData(saveRet)
		return ok
	}

}

func main() {
	//	timeOffset := flag.Int("flagname", 1, "help message for flagname")
	//	starTime := flag.String("s", "startime", "the begin time of query")
	//	endTime := flag.String("e", "endTime", "the end time of query")
	timeOffset, _ := strconv.Atoi(os.Args[1]) // string to int
	starTime := os.Args[2]
	endTime := os.Args[3]
	fmt.Println(mainRun(timeOffset, starTime, endTime))
	//	fmt.Println(mainRun(5, "2017-05-01", "2017-05-06"))

}
