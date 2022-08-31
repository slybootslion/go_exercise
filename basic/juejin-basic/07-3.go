package main

import "fmt"

func main() {
	for i := 2; i > 0; i++ {
		if i > 10 {
			break
		}

		if i == 2 {
			fmt.Println(i)
			continue
		}
		//假定i为素数
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				//当i能被某个整数整除时，不是素数
				flag = false
				break
			}
		}
		//如果依旧为true，则i为素数
		if flag {
			fmt.Println(i)
		}
	}
}
