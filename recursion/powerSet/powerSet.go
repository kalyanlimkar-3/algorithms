package main

import "fmt"

type tower struct {
	name  string
	disks []int
}

func main() {
	var s string = "abc"
	powerSet(s, -1, "")
}

func powerSet(s string, index int, c string) {
	n := len(s)

	if index == n {
		return
	}

	fmt.Println(c)

	for i := index + 1; i < len(s); i++ {
		c += string(s[i])
		powerSet(s, i, c)
		c = c[:len(c)-1]
	}
}
