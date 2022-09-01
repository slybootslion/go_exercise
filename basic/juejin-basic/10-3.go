package main

import "fmt"

func main() {
	defer fmt.Print("素数")
	defer fmt.Print("查找")
	var resultSlice []int
	FindPrimeNumber2(&resultSlice, 10)
	fmt.Println(resultSlice)
}

func FindPrimeNumber2(result *[]int, max int) {
	for i := 2; i < max; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
			}
		}
		if flag {
			*result = append(*result, i)
		}
	}
}
