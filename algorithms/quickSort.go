package main

import (
	"fmt"
)

func quickSort(list []uint8) []uint8 {
	if len(list) < 2 {
		return list
	} else {
		pivot := list[len(list)/2]
		var less, greater []uint8
		for i := 0; i < len(list); i++ {
			if list[i] < pivot {
				less = append(less, list[i])
			} else if list[i] > pivot {
				greater = append(greater, list[i])
			}
		}
		return append(append(quickSort(less), pivot), quickSort(greater)...)
	}
}

func main() {
	seq := []uint8{9, 5, 7, 1, 4, 6, 2, 3, 8}

	fmt.Println(seq)
	fmt.Println(quickSort(seq))
}
