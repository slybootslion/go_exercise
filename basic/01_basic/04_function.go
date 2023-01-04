package main

import "fmt"

func main() {
	var resultSlice []int
	findPrimeNumber(&resultSlice, 10)
	fmt.Println(resultSlice)
}

func findPrimeNumber(resultSlice *[]int, max int) {
	for i := 2; i < max; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
			}
		}
		if flag {
			*resultSlice = append(*resultSlice, i)
		}
	}
}
