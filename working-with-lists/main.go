package main

import "fmt"

func main() {
	number := factorial(5)
	fmt.Println(number)
}

func factorial(number int) int {
	if number == 1 {
		return 1
	}

	return number * factorial(number-1)
}
