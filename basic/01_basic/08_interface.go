package main

import "fmt"

func main() {
	var imageLoader ImageDownloader
	imageLoader = new(fileCache)
	data := imageLoader.FetchImage("https://www.example.com/a.png")
	fmt.Println(data)
	if data == "" {
		var imageLoader2 ImageDownloader
		imageLoader2 = new(netFetch)
		data2 := imageLoader2.FetchImage("https://www.example.com/a.png")
		fmt.Println(data2)
	}
}

type ImageDownloader interface {
	FetchImage(url string) string
}

type fileCache struct {
}

func (f *fileCache) FetchImage(url string) string {
	return "从本地缓存中获取图片：" + url
}

type netFetch struct {
}

func (f *netFetch) FetchImage(url string) string {
	return "从网络下载图片：" + url
}
