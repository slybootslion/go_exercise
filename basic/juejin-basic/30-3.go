package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var testInt1 int32 = 0
var syncWait5 sync.WaitGroup

func main() {
	syncWait5.Add(2)
	go testFunc5()
	go testFunc5()
	syncWait5.Wait()
	fmt.Println(testInt1)
}

func testFunc5() {
	defer syncWait5.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(&testInt1, 1)
	}
}
