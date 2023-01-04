package main

import "fmt"

func main() {
	var resultArray [4]int
	arrayIndex := 0
	for i := 2; i < 10; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
			}
		}
		if flag {
			resultArray[arrayIndex] = i
			arrayIndex++
		}
	}
	fmt.Println(resultArray)
}
