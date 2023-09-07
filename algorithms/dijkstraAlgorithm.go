package main

import "fmt"

func makeWeightedGraph() map[uint8]map[uint8]uint8 {
	graph := make(map[uint8]map[uint8]uint8)

	graph[1] = make(map[uint8]uint8)
	graph[1][2] = 7
	graph[1][3] = 9
	graph[1][6] = 14
	graph[2] = make(map[uint8]uint8)
	graph[2][3] = 10
	graph[2][4] = 15
	graph[3] = make(map[uint8]uint8)
	graph[3][4] = 11
	graph[3][6] = 2
	graph[6] = make(map[uint8]uint8)
	graph[6][5] = 9
	graph[5] = make(map[uint8]uint8)
	graph[5][4] = 6

	return graph
}

func makeParentsAndCosts(graph *map[uint8]map[uint8]uint8) (map[uint8]uint8, map[uint8]uint8) {
	var parents = make(map[uint8]uint8)
	var costs = make(map[uint8]uint8)

	for i, node := range *graph {
		if i == 1 {
			for child, weight := range node {
				parents[child] = i
				costs[child] = weight
			}
		} else {
			_, ok := costs[i]
			if !ok {
				costs[i] = 255
				parents[i] = 0
			}
		}
	}
	parents[4] = 0
	costs[4] = 255
	return parents, costs
}

func findMinWeight(costs map[uint8]uint8, searched map[uint8]bool) uint8 {
	var minNode uint8 = 255
	for node, weight := range costs {
		if weight < minNode && !searched[node] {
			minNode = node
		}
	}
	return minNode
}

func dijkstra(graph map[uint8]map[uint8]uint8, parents, costs map[uint8]uint8) uint8 {
	var minWeight uint8
	searched := make(map[uint8]bool)
	curNode := findMinWeight(costs, searched)
	node := graph[curNode]

	for node != nil {
		fixCost := costs[curNode]
		for neighbour, weight := range node {
			curCost := fixCost + weight

			if curCost < costs[neighbour] {
				costs[neighbour] = curCost
				parents[neighbour] = curNode
			}
		}

		searched[curNode] = true
		curNode = findMinWeight(costs, searched)
		node = graph[curNode]
	}
	minWeight = costs[4]
	return minWeight
}

func main() {
	graph := makeWeightedGraph()
	parents, costs := makeParentsAndCosts(&graph)

	fmt.Println(dijkstra(graph, parents, costs))
}
