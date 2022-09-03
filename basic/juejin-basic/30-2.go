package main

import (
	"fmt"
	"sync"
	"time"
)

var syncWait4 sync.WaitGroup
var locker1 sync.RWMutex

func main() {
	syncWait4.Add(6)
	go write1Sec()
	time.Sleep(time.Second * 1)
	go read5Sec()
	time.Sleep(time.Second * 2)
	go read3Sec()
	time.Sleep(time.Millisecond * 500)
	go write3Sec()
	time.Sleep(time.Millisecond * 500)
	go read1Sec()
	go write5Sec()
	syncWait4.Wait()
	fmt.Println("程序运行结束", time.Now().Format("15:04:05"))

}

func read5Sec() {
	defer syncWait4.Done()
	defer locker1.RUnlock()
	locker1.RLock()
	fmt.Println("读文件耗时5秒 开始", time.Now().Format("15:04:05"))
	time.Sleep(time.Second * 5)
	fmt.Println("读文件耗时5秒 结束", time.Now().Format("15:04:05"))
}
func read3Sec() {
	defer syncWait4.Done()
	defer locker1.RUnlock()
	locker1.RLock()
	fmt.Println("读文件耗时3秒 开始", time.Now().Format("15:04:05"))
	time.Sleep(time.Second * 3)
	fmt.Println("读文件耗时3秒 结束", time.Now().Format("15:04:05"))
}
func read1Sec() {
	defer syncWait4.Done()
	defer locker1.RUnlock()
	locker1.RLock()
	fmt.Println("读文件耗时1秒 开始", time.Now().Format("15:04:05"))
	time.Sleep(time.Second * 1)
	fmt.Println("读文件耗时1秒 结束", time.Now().Format("15:04:05"))
}

func write5Sec() {
	defer syncWait4.Done()
	defer locker1.Unlock()
	locker1.Lock()
	fmt.Println("写文件耗时5秒 开始", time.Now().Format("15:04:05"))
	time.Sleep(time.Second * 5)
	fmt.Println("写文件耗时5秒 开始", time.Now().Format("15:04:05"))
}

func write3Sec() {
	defer syncWait4.Done()
	defer locker1.Unlock()
	locker1.Lock()
	fmt.Println("写文件耗时3秒 开始", time.Now().Format("15:04:05"))
	time.Sleep(time.Second * 3)
	fmt.Println("写文件耗时3秒 结束", time.Now().Format("15:04:05"))
}

func write1Sec() {
	defer syncWait4.Done()
	defer locker1.Unlock()
	locker1.Lock()
	fmt.Println("写文件耗时1秒 开始", time.Now().Format("15:04:05"))
	time.Sleep(time.Second * 1)
	fmt.Println("写文件耗时1秒 结束", time.Now().Format("15:04:05"))
}
