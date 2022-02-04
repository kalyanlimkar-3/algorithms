package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Println(factorial(n))
}

func factorial(n int) int {
	fact := 1

	for i := 1; i <= n; i++ {
		fact = fact * i
	}

	return fact
}
