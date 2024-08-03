package main

import "fmt"

const (
	exist    = true
	notExist = false
)

func binarySearch(arr []int, search int) (int, bool) {
	countIterations := 0
	low := 0
	high := len(arr) - 1

	for high >= low {
		countIterations++
		mid := (high + low) / 2
		if arr[mid] < search {
			low = mid
		} else if arr[mid] > search {
			high = mid
		} else if arr[mid] == search {
			fmt.Printf("needed %d iterations\n", countIterations)
			return mid, exist
		}
	}

	return 0, notExist
}

func main() {
	arr := newSortedSlice(0, 2, 15)
	fmt.Printf("sequence: %d\n", arr)

	searchElement := 10

	index, ok := binarySearch(arr, searchElement)
	if ok {
		fmt.Printf("The element %d has index %d\n", searchElement, index)
	} else {
		fmt.Printf("The sequence doesn't have the element %d\n", searchElement)
	}
}

func newSortedSlice(firstVal, step, size int) []int {
	s := []int{firstVal}

	for i := 1; i < size; i++ {
		s = append(s, s[i-1]+step)
	}

	return s
}
