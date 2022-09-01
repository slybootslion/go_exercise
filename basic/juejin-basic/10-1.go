package main

import "fmt"

func main() {
	var resultSlice []int
	resultSlice = findPrimeNumber(10)
	fmt.Println(resultSlice)
}

func findPrimeNumber(n int) []int {
	var resultSlice []int

	for i := 2; i < n; i++ {
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

	return resultSlice
}
