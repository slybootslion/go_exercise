package main

import (
	"fmt"
	"sync"
	"time"
)

var syncWait2 sync.WaitGroup

const DaysOfWeek = 7

var intChan1 = make(chan int, DaysOfWeek)

func main() {
	syncWait2.Add(2)
	go layEggs1()
	go collectEggs(intChan1)
	syncWait2.Wait()
}

func layEggs1() {
	defer syncWait2.Done()
	for i := 0; i < DaysOfWeek; i++ {
		time.Sleep(time.Millisecond * 500)
		intChan1 <- 1
		fmt.Println("产鸡蛋了")
	}
}

func collectEggs(intChan1 chan int) {
	defer syncWait2.Done()
	var eggCount int
	for i := 0; i < DaysOfWeek; i++ {
		eggCount += <-intChan1
		fmt.Println("鸡蛋被收集了")
	}
	fmt.Printf("本周共产%d个鸡蛋\n", eggCount)
}
