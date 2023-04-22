package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(solve(input))
}

type Location struct {
	pos       [2]int
	distances map[string]int
}

func solve(input []string) (int, int) {
	locations := map[string]Location{}
	for j, line := range input {
		for i, char := range line {
			if char != '#' && char != '.' {
				locations[string(char)] = Location{
					pos:       [2]int{i, j},
					distances: map[string]int{},
				}
			}
		}
	}
	ht := len(input)
	wd := len(input[0])
	visit := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, loc := range locations {
		visited := map[[2]int]bool{}
		visited[loc.pos] = true
		queue := [][3]int{{loc.pos[0], loc.pos[1], 0}}
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			for _, vis := range visit {
				toVisit := [2]int{curr[0] + vis[0], curr[1] + vis[1]}
				if toVisit[0] >= 0 && toVisit[0] < wd && toVisit[1] >= 0 && toVisit[1] < ht && input[toVisit[1]][toVisit[0]] != '#' && !visited[toVisit] {
					visited[toVisit] = true
					if input[toVisit[1]][toVisit[0]] != '.' {
						loc.distances[string(input[toVisit[1]][toVisit[0]])] = curr[2] + 1
					}
					queue = append(queue, [3]int{toVisit[0], toVisit[1], curr[2] + 1})
				}
			}
		}
	}
	minDist := math.MaxInt
	minDist2 := math.MaxInt
	visited := map[string]bool{}
	visited["0"] = true
	travelingSalesman(&minDist, &minDist2, &visited, 0, "0", &locations)
	return minDist, minDist2
}

func travelingSalesman(minDist *int, minDist2 *int, visited *map[string]bool, currDist int, route string, locations *map[string]Location) {
	currLocation := (*locations)[string(route[len(route)-1])]
	if len(route) == len(*locations) {
		if currDist < *minDist {
			*minDist = currDist
		}
		distanceToZero := currLocation.distances["0"]
		if currDist+distanceToZero < *minDist2 {
			*minDist2 = currDist + distanceToZero
		}
		return
	}
	for loc, dist := range currLocation.distances {
		if !(*visited)[loc] {
			(*visited)[loc] = true
			travelingSalesman(minDist, minDist2, visited, currDist+dist, route+loc, locations)
			(*visited)[loc] = false
		}
	}
}
