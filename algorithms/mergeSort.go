package main

import (
	"fmt"
)

func mergeSort(list []uint8) []uint8 {
	if len(list) < 2 {
		return list
	}

	left := mergeSort(list[:len(list)/2])
	right := mergeSort(list[len(list)/2:])
	lIndex, rIndex, aIndex := 0, 0, 0
	answer := make([]uint8, len(list))

	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] > right[rIndex] {
			answer[aIndex] = right[rIndex]
			rIndex++
		} else {
			answer[aIndex] = left[lIndex]
			lIndex++
		}
		aIndex++
	}
	for lIndex < len(left) {
		answer[aIndex] = left[lIndex]
		aIndex++
		lIndex++
	}
	for rIndex < len(right) {
		answer[aIndex] = right[rIndex]
		aIndex++
		rIndex++
	}
	return answer
}

func main() {
	seq := []uint8{2, 6, 8, 3, 0, 9, 1, 4, 7, 5}

	fmt.Println(seq)
	fmt.Println(mergeSort(seq))
}
