package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("an error occurred while reading input. please try again", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. please enter an integer value")
			continue
		}
		fmt.Println("you guess is", guess)
		if guess > secretNumber {
			fmt.Println("your guess is bigger than the secret number. please try again")
		} else if guess < secretNumber {
			fmt.Println("your guess is smaller than the secret number. please try again")
		} else {
			fmt.Println("correct, you legend!")
			break
		}
	}
}

// https://github.com/wangkechun/go-by-example
