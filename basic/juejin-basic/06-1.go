package main

import "fmt"

func main() {
	var exampleNumberA int = 10
	exampleNumberAPtr := &exampleNumberA
	fmt.Println(exampleNumberAPtr)

	exampleNumberAPtrValue := *exampleNumberAPtr
	fmt.Println(exampleNumberAPtrValue)

	exampleNumberAPtr2 := new(int64)
	*exampleNumberAPtr2 = 100
	fmt.Println(*exampleNumberAPtr2)
}
