package main

import "fmt"

func main() {
	outputMode := "全"
	n := 7
	if outputMode == "上" || outputMode == "全" {
		for i := 1; i <= n; i++ {
			for j := 0; j < n-i; j++ {
				fmt.Print(" ")
			}
			for k := 0; k < 2*i-1; k++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
	if outputMode == "下" || outputMode == "全" {
		for i := 1; i <= n; i++ {
			for j := 0; j < i; j++ {
				fmt.Print(" ")
			}
			for k := 0; k < 2*n-1-2*i; k++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}
