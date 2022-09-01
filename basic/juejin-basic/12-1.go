package main

import (
	"fmt"
	"time"
)

var percent = 0

func main() {
	var keepChecking = true
	fmt.Println("开始下载任务！")
	go download(" ", func() {
		keepChecking = false
		fmt.Println("下载完成")
	}, func(currentPercent int) {
		keepChecking = false
		fmt.Println("下载取消", currentPercent)
	}, func(currentPercent int) {
		keepChecking = false
		fmt.Println("下载失败", currentPercent)
	})
	for {
		if keepChecking {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("当前进度", getPercent())
		} else {
			break
		}
	}
}

func getPercent() int {
	return percent
}

func download(
	url string, downloadSuccess func(),
	downloadCancelled func(int),
	downloadFailed func(int),
) {
	for {
		time.Sleep(1 * time.Second)
		percent += 7
		//if percent >= 30 {
		//	downloadFailed(percent)
		//	break
		//}
		//if percent >= 10 {
		//	downloadFailed(percent)
		//	break
		//}
		if percent >= 100 {
			downloadSuccess()
			break
		}
	}
}
