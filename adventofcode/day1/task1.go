package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inventory := make([]int, 1)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Elves list")
	for flag, index := false, 0; ; {
		res, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		if res == "\n" {
			if flag {
				break
			} else {
				flag = true
				continue
			}
		}
		if flag {
			flag = false
			inventory = append(inventory, 0)
			index++
		}

		res = res[:len(res)-1]
		ires, err := strconv.Atoi(res)
		if err != nil {
			break
		}
		inventory[index] += ires
	}

	var iMax, max int
	for i := 0; i < len(inventory)-1; i++ {
		if inventory[i] > max {
			max = inventory[i]
			iMax = i
		}

	}
	fmt.Printf("\nElve %v carrying the most (%v) Calories\n", iMax, max)
}
