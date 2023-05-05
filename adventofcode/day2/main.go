package main

import (
	"fmt"
)

func roundScore(round string) (int, bool) {
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
		return res, true
	}
	fmt.Println("Something go wrong. Try to enter the round again")
	return 0, false
}

func main() {
	var res, N int
	fmt.Println("How many rounds were played?")
	_, err := fmt.Scanln(&N)
	if err != nil {
		return
	}
	fmt.Println("Enter your rounds:")
	for i := 0; i < N; i++ {
		var player1, player2 string
		_, err1 := fmt.Scan(&player1)
		_, err2 := fmt.Scan(&player2)
		if err1 != nil || err2 != nil {
			return
		}

		score, flag := roundScore(player1 + " " + player2)
		if flag {
			res += score
		} else {
			i--
		}
	}
	fmt.Println(res)
}
