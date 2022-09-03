package main

import (
	"fmt"
	"sync"
	"time"
)

var goRoutineWait sync.WaitGroup

func main() {
	goRoutineWait.Add(1)
	go testFunc1()
	goRoutineWait.Wait()
	fmt.Println("程序运行结果")
}

func testFunc1() {
	defer goRoutineWait.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("第%d次运行\n", i+1)
		time.Sleep(time.Second)
	}
}
