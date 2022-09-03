package main

import (
	"fmt"
	"sync"
	"time"
)

var goRoutineWait1 sync.WaitGroup

func main() {
	goRoutineWait1.Add(2)
	go testFunc2()
	go testFunc2()
	goRoutineWait1.Wait()
	fmt.Println("程序运行结束")
}

func testFunc2() {
	defer goRoutineWait1.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("第%d次运行\n", i+1)
		time.Sleep(time.Second)
	}
}
