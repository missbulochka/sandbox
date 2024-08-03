package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int, l, r int) {
	if l < r {
		q := partition(arr, l, r)
		quickSort(arr, l, q)
		quickSort(arr, q+1, r)
	}
}

func partition(arr []int, l, r int) int {
	pivot := arr[(l+r)/2]
	left := l
	right := r

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left >= right {
			break
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	return right
}

func main() {
	arr := newRandomSlice(15)
	fmt.Printf("sequence: %d\n", arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("sorted sequence: %d\n", arr)
}

func newRandomSlice(size int) []int {
	s := []int{}

	for len(s) != size {
		s = append(s, rand.Int()%100)
	}

	return s
}
