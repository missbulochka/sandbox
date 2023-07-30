package main

import "fmt"

func binarySearch(list []uint8, item uint8) int { // О(log(N))
	low, high, mid := 0, len(list)-1, -1

	for low <= high {
		mid = (low + high) / 2
		if item == list[mid] {
			return mid
		} else if item > list[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	seq := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Print(binarySearch(seq, 3))
}
