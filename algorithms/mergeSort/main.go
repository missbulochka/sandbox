package main

import (
	"fmt"
	"math/rand"
)

func mergeSort(arr []int, l, r int) {
	if l < r {
		mid := (l + r) / 2
		mergeSort(arr, l, mid)
		mergeSort(arr, mid+1, r)
		merge(arr, l, mid, r)
	}
}

func merge(arr []int, l, mid, r int) {
	// f*cking slices
	left := make([]int, (mid-l)+1)
	right := make([]int, (r-mid)+1)
	copy(left, arr[l:mid+1])
	copy(right, arr[mid+1:r])

	i, j, p := 0, 0, 0
	for l+i <= mid && mid+1+j < r {
		if left[i] < right[j] {
			arr[l+p] = left[i]
			i++
		} else {
			arr[l+p] = right[j]
			j++
		}
		p++
	}

	for l+i <= mid {
		arr[l+p] = left[i]
		i++
		p++
	}

	for mid+1+j < r {
		arr[l+p] = right[j]
		j++
		p++
	}
}

func main() {
	//arr := newRandomSlice(15)
	arr := []int{11, 5, 98, 92, 38, 18, 39, 26, 36, 81, 5, 30, 47, 35, 63}
	fmt.Printf("sequence: %d\n", arr)

	mergeSort(arr, 0, len(arr)-1)
	fmt.Printf("sorted sequence: %d\n", arr)
}

func newRandomSlice(size int) []int {
	s := []int{}

	for len(s) != size {
		s = append(s, rand.Int()%100)
	}

	return s
}
