package main

import (
	"fmt"
	"runtime"
	"sync"
)

var syncWait sync.WaitGroup

func main() {
	syncWait.Add(1)
	go testFunc3()
	syncWait.Wait()
	fmt.Println("程序运行结束")
}

func testFunc3() {
	defer syncWait.Done()
	for i := 0; i < 101; i++ {
		fmt.Println(i)
		if i >= 15 {
			runtime.Goexit()
		}
	}
}
