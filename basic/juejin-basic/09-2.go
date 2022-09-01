package main

import "fmt"

func main() {
	var resultArray [11]bool
	for i := 0; i < 11; i++ {
		resultArray[i] = false
	}
	for i := 2; i < 10; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
			}
		}
		if flag {
			resultArray[i] = true
		}
	}
	for i := 0; i < 11; i++ {
		fmt.Println(i, resultArray[i])
	}
}
