package main

import (
	"fmt"
	"time"
)

type process struct {
	current int
	total   int
}

func main() {
	chan1 := make(chan process)
	chan2 := make(chan int)
	go recvFunc(chan1, chan2)
	go sendFunc1(chan1)
	go sendFunc2(chan2)
	time.Sleep(10 * time.Second)
	fmt.Println("下载完成")
}

func sendFunc1(chan1 chan process) {
	for i := 0; i < 10; i++ {
		chan1 <- process{
			current: i,
			total:   10,
		}
		time.Sleep(1 * time.Second)
	}
}

func sendFunc2(chan2 chan int) {
	time.Sleep(2 * time.Second)
	chan2 <- 1
	time.Sleep(2 * time.Second)
	chan2 <- 1
}

func recvFunc(chan1 chan process, chan2 chan int) {
	for {
		select {
		case processInfo := <-chan1:
			fmt.Printf("当前任务进度：%d\n", 100.0*processInfo.current/processInfo.total)
		case <-chan2:
			fmt.Println("添加了新任务")
		}
	}
}
