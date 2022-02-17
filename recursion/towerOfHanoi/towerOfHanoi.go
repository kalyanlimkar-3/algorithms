package main

import (
	"fmt"
)

type tower struct {
	name  string
	disks []int
}

func main() {
	source := &tower{name: "Source", disks: []int{3, 2, 1}}
	auxiliary := &tower{name: "Auxiliary", disks: []int{}}
	destination := &tower{name: "Destination", disks: []int{}}

	fmt.Println(source.name, ":", source.disks, auxiliary.name, ":", auxiliary.disks, destination.name, ":", destination.disks)
	fmt.Println()

	towerOfHanoi(len(source.disks), source, destination, auxiliary)

	fmt.Println()
	fmt.Println(source.name, ":", source.disks, auxiliary.name, ":", auxiliary.disks, destination.name, ":", destination.disks)
}

func towerOfHanoi(nDisks int, source *tower, destination *tower, auxiliary *tower) {
	if nDisks > 0 {
		towerOfHanoi(nDisks-1, source, auxiliary, destination)

		fmt.Println("Moving disk", source.disks[len(source.disks)-1], "from", source.name, "to", destination.name)
		destination.disks = append(destination.disks, source.disks[len(source.disks)-1])
		source.disks = source.disks[:len(source.disks)-1]

		towerOfHanoi(nDisks-1, auxiliary, destination, source)
	}
}
