package main

import "fmt"

func main() {
	fmt.Println(calcFactorial(5))
	fmt.Println(calcFibonacci(8))
}

func calcFactorial(n int) (result int) {
	if n > 0 {
		result = n * calcFactorial(n-1)
		return result
	}
	return 1
}

func calcFibonacci(n int) (result int) {
	if n <= 2 {
		return 1
	} else {
		return calcFibonacci(n-1) + calcFibonacci(n-2)
	}
}
