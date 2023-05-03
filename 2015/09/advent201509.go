package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	distances := parseInput(input)
	fmt.Println(solve(distances))
}

func parseInput(input []string) *map[string]map[string]int {
	distances := map[string]map[string]int{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		if _, ok := distances[parts[0]]; !ok {
			distances[parts[0]] = map[string]int{}
		}
		if _, ok := distances[parts[2]]; !ok {
			distances[parts[2]] = map[string]int{}
		}
		n, _ := strconv.Atoi(parts[4])
		distances[parts[0]][parts[2]] = n
		distances[parts[2]][parts[0]] = n
	}
	return &distances
}

func travelingSalesman(shortest *int, longest *int, visited *map[string]bool, distances *map[string]map[string]int, currentLocation string, currentDist int, nVisited int) {
	if nVisited == len(*distances) {
		if currentDist < *shortest {
			*shortest = currentDist
		}
		if currentDist > *longest {
			*longest = currentDist
		}
		return
	}
	for nextLocation, dist := range (*distances)[currentLocation] {
		if !(*visited)[nextLocation] {
			(*visited)[nextLocation] = true
			travelingSalesman(shortest, longest, visited, distances, nextLocation, currentDist+dist, nVisited+1)
			(*visited)[nextLocation] = false
		}
	}

}

func solve(distances *map[string]map[string]int) (int, int) {
	shortest := math.MaxInt
	longest := 0
	visited := map[string]bool{}
	for location := range *distances {
		visited[location] = true
		travelingSalesman(&shortest, &longest, &visited, distances, location, 0, 1)
		visited[location] = false
	}
	return shortest, longest
}
