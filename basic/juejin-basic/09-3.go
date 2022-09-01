package main

import "fmt"

func main() {
	var studentInfos = make(map[string]string)
	studentInfos["0001"] = "王小红"
	studentInfos["0002"] = "李小明"
	studentInfos["0003"] = "张三丰"
	studentInfos["0004"] = "孙小贝"
	studentInfos["0005"] = "何明明"

	for key, value := range studentInfos {
		fmt.Println("学号：", key, "姓名：", value)
	}
}
