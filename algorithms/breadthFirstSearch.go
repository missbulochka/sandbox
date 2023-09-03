package main

import (
	"fmt"
)

func makeGraph() map[uint8][]uint8 {
	graph := make(map[uint8][]uint8)

	graph[1] = []uint8{3, 4, 6}
	graph[2] = []uint8{1}
	graph[3] = []uint8{2, 5}
	graph[4] = []uint8{6}
	graph[5] = []uint8{}
	graph[6] = []uint8{4}
	graph[7] = []uint8{4}

	return graph
}

func BFS(graph map[uint8][]uint8, start, end uint8) bool {
	var queue []uint8
	searched := make(map[uint8]bool)
	queue = append(queue, graph[start]...)

	for len(queue) > 0 {
		if !searched[queue[0]] {
			if queue[0] == end {
				return true
			} else {
				searched[queue[0]] = true
				queue = append(queue, graph[queue[0]]...)
			}
		}
		queue = queue[1:]
	}
	return false
}

func main() {
	graph := makeGraph()

	fmt.Println(graph)
	if BFS(graph, 1, 5) {
		fmt.Println("Path found!")
	} else {
		fmt.Println("Path not found")
	}
}
