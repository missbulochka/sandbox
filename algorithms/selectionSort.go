package main

import "fmt"

func selectionSort(list []uint8) []uint8 {
	for i := 0; i < len(list)-1; i++ {
		indexOfMin := i
		for j := i; j < len(list); j++ {
			if list[j] < list[indexOfMin] {
				indexOfMin = j
			}
		}
		list[i], list[indexOfMin] = list[indexOfMin], list[i]
	}
	return list
}

func main() {
	seq := []uint8{9, 5, 7, 1, 4, 6, 2, 3, 8}

	fmt.Println(seq)
	fmt.Println(selectionSort(seq))
}
