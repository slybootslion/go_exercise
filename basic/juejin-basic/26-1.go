package main

import (
	"fmt"
	"runtime"
)

func main() {
	go fmt.Println("Hello World")
	runtime.Gosched()
	fmt.Println("程序运行结束")
}
