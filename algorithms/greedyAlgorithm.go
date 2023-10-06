package main

import "fmt"

func makeTask() (map[string]bool, map[string][]string) {
	set := map[string]bool{
		"mt": false,
		"wa": false,
		"or": false,
		"id": false,
		"nv": false,
		"ut": false,
		"са": false,
		"az": false,
	}

	stations := make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "са"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}

	return set, stations
}

func countCovered(need map[string]bool, curStation []string) []string {
	var covered []string
	for _, st := range curStation {
		if _, ok := need[st]; ok {
			covered = append(covered, st)
		}
	}
	return covered
}

func greedyAlgorythm(stations map[string][]string, need map[string]bool) []string {
	var result []string

	for len(need) > 0 {
		var bestStation string
		var statesOfBestStation []string

		for station, states := range stations {
			covered := countCovered(need, states)
			if len(covered) > len(bestStation) {
				bestStation = station
				statesOfBestStation = covered
			}
		}

		result = append(result, bestStation)
		for _, st := range statesOfBestStation {
			delete(need, st)
		}
	}

	return result
}

func main() {
	set, stations := makeTask()
	fmt.Println(stations)

	bestRes := greedyAlgorythm(stations, set)
	for _, el := range bestRes {
		fmt.Println(el)
	}
}
