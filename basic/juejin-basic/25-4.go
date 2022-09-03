package main

import (
	"fmt"
	"sync"
	"time"
)

var goRoutineWait2 sync.WaitGroup

func main() {
	goRoutineWait2.Add(1)
	go func() {
		defer goRoutineWait2.Done()
		for i := 0; i < 3; i++ {
			fmt.Printf("第%d次运行\n", i+1)
			time.Sleep(time.Second)
		}
	}()
	goRoutineWait2.Wait()
	fmt.Println("程序运行结束")
}
