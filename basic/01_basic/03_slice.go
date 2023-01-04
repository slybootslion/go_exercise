package main

import "fmt"

func main() {
	var resultSlice []int
	for i := 2; i < 30; i++ {
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
	fmt.Println(resultSlice)
}
