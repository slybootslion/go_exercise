package main

import "fmt"

type ImageDownloader interface {
	FetchImage(url string) string
}

type fileCache struct {
}

func (f *fileCache) FetchImage(url string) string {
	if url == "" {
		return ""
	}
	return "从本地缓存中获取图片：" + url
}

type netFetch struct {
}

func (n *netFetch) FetchImage(url string) string {
	return "从网络下载图片：" + url
}

func main() {
	var imageLoader ImageDownloader
	imageLoader = new(fileCache)
	data := imageLoader.FetchImage("")
	if data != "" {
		fmt.Println(data)
	}
	if data == "" {
		var imageLoader2 ImageDownloader = new(netFetch)
		data2 := imageLoader2.FetchImage("https://www.example.com/a.png")
		fmt.Println(data2)
	}
}
