package main

import (
	"bufio"
	"fmt"
	"os"
)

func roundScore(round string) int {
	var pairs = map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
	if res, ok := pairs[round]; ok {
		return res
	}
	fmt.Println("Something go wrong. Try to enter the round again")
	return 0
}

func main() {
	var res int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your rounds:")
	for {
		scanner.Scan()
		round := scanner.Text()
		if len(round) != 0 {
			res += roundScore(round)
		} else {
			break
		}
	}
	fmt.Println(res)
}
