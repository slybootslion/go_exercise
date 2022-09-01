package main

import "fmt"

func main() {
	fmt.Println(calcFibonacci(10))
}

func calcFibonacci(n int) (result int) {
	if n <= 2 {
		result = 1
	} else {
		result = calcFibonacci(n-1) + calcFibonacci(n-2)
	}

	return result
}
