package main

import "fmt"

func main() {
	var arr [5]int

	for i, j := 0, 5; i < 5; i, j = i+1, j-1 {
		arr[i] = j
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
