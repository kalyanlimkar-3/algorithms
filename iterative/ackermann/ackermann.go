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
	stack := []int{}

	for {
		if m == 0 {
			n = n + 1
			if len(stack) == 0 {
				return n
			}

			m = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if n == 0 {
			m = m - 1
			n = 1
		} else {
			stack = append(stack, m-1)
			n = n - 1
		}
	}
}
