package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [5]int

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(200)	// Returns random number between 0 to 200.
	}

	fmt.Println("Input array:", arr)

	for i := 0; i < len(arr); i++ {
		min := i

		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	fmt.Println("Array after sorting:", arr)
}
