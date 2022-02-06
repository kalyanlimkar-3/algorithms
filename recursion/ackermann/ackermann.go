package main

import (
	"fmt"
)

func main() {
	m := 2
	n := 1
	fmt.Println(ackermann(m, n))
}

func ackermann(m int, n int) int {
	if m == 0 {
		return n + 1
	}

	if n == 0 {
		return ackermann(m-1, 1)
	}

	return ackermann(m-1, ackermann(m, n-1))
}
