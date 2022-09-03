package main

import (
	"fmt"
	"sync"
)

var testInt = 0
var syncWait3 sync.WaitGroup
var locker sync.Mutex

func main() {
	syncWait3.Add(2)
	go testFunc4()
	go testFunc4()
	syncWait3.Wait()
	fmt.Println(testInt)
}

func testFunc4() {
	defer syncWait3.Done()
	defer locker.Unlock()
	locker.Lock()
	for i := 0; i < 1000; i++ {
		testInt++
	}
}
