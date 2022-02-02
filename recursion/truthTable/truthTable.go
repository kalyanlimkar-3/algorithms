package main

import (
	"fmt"
)

func main() {
	permutations("", 3)
}

func permutations(b string, size int) {
	if size == 0 {
		fmt.Println(string(b))
		return
	}

	permutations(b+"T", size-1)
	permutations(b+"F", size-1)
}
