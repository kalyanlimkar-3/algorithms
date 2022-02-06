package main

import (
	"fmt"
)

func main() {
	n := 7
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
