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

	var iminMax, minMax int
	var imidMax, midMax int
	var imaxMax, maxMax int
	for i := 0; i < len(inventory); i++ {
		if inventory[i] > minMax {
			minMax = inventory[i]
			iminMax = i + 1
		} else {
			continue
		}
		if inventory[i] > midMax {
			minMax = midMax
			midMax = inventory[i]
			iminMax = imidMax
			imidMax = i + 1
		} else {
			continue
		}
		if inventory[i] > maxMax {
			midMax = maxMax
			maxMax = inventory[i]
			imidMax = imaxMax
			imaxMax = i + 1
		}

	}
	fmt.Printf("\nElves %v, %v and %v are carrying %v Calories\n", iminMax, imidMax, imaxMax, minMax+midMax+maxMax)
}
