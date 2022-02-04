package main

import (
	"fmt"
)

func main() {
	n := 5
	fmt.Println(factorial(n))
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
