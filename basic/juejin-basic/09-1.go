package main

import "fmt"

func main() {
	var resultSlice []int
	for i := 0; i < 1000; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
			}
		}
		if flag {
			resultSlice = append(resultSlice, i)
		}
	}
	fmt.Println(len(resultSlice))
}
