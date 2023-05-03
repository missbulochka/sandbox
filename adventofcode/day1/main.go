package main

import "fmt"

func main() {
	var N int

	fmt.Println("How many elves on the team have inventory?")
	_, err := fmt.Scanln(&N)
	if err != nil {
		return
	}

	inventory := make([]int, N)
	fmt.Println("Enter Elves list")
	for i := 0; i < N; i++ {
		var res int
		for err2 := err; err2 == nil; {
			inventory[i] += res
			_, err2 = fmt.Scanln(&res)
		}
	}

	var iMax, max int

	for i := 0; i < N; i++ {
		if inventory[i] > max {
			max = inventory[i]
			iMax = i
		}
	}

	fmt.Printf("Elve %v carrying the most (%v) Calories\n", iMax, max)
}
