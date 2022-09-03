package main

import (
	"fmt"
	"time"
)

func main() {
	go testFunc()
	time.Sleep(time.Second * 5)
	fmt.Println("程序运行结束")
}

func testFunc() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("第%d次运行\n", i)
		time.Sleep(time.Second)
	}
}
