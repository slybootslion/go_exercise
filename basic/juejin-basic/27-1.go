package main

import (
	"fmt"
	"sync"
)

var syncWait1 sync.WaitGroup

var intChan = make(chan int)

func main() {
	syncWait1.Add(2)
	go layEggs()

	go eatEggs(intChan)

	syncWait1.Wait()
}

func layEggs() {
	defer syncWait1.Done()
	intChan <- 1
	close(intChan)
}

func eatEggs(intChan chan int) {
	defer syncWait1.Done()
	eggCounts := <-intChan
	fmt.Printf("吃%d个荷包蛋", eggCounts)
}
