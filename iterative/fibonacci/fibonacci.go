package main

import (
	"fmt"
)

func main() {
  n := 7
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) (sum int) {
	if n == 0 || n == 1 {
		return n
	}

	i := 1
	a := 0
	b := 1

	for i < n {
		sum = a + b

		a = b
		b = sum

		i++
	}

	return sum
}
