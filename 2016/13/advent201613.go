package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input, _ := strconv.Atoi(string(data))
	fmt.Println(part1(input))
}

func mapout(coord [2]int, layout *map[[2]int]bool, input int) bool {
	if _, ok := (*layout)[coord]; !ok {
		sum := coord[0]*coord[0] + 3*coord[0] + 2*coord[0]*coord[1] + coord[1] + coord[1]*coord[1] + input
		count := 0
		for sum > 0 {
			count += sum & 1
			sum >>= 1
		}
		if count%2 == 1 {
			(*layout)[coord] = false
		} else {
			(*layout)[coord] = true
		}
	}
	return (*layout)[coord]
}

func part1(input int) int {
	queue := [][3]int{{1, 1, 0}}
	steps := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	layout := map[[2]int]bool{}
	visited := map[[2]int]int{}
	layout[[2]int{1, 1}] = true
	visited[[2]int{1, 1}] = 0
	i := 0
	for i < len(queue) {
		current := queue[i]
		pos := [2]int{current[0], current[1]}
		dist := current[2]
		for _, step := range steps {
			coord := [2]int{pos[0] + step[0], pos[1] + step[1]}
			if coord[0] < 0 || coord[1] < 0 {
				continue
			}
			if coord[0] == 31 && coord[1] == 39 {
				return dist + 1
			}
			if mapout(coord, &layout, input) {
				d, ok := visited[coord]
				if !ok || dist+1 < d {
					queue = append(queue, [3]int{coord[0], coord[1], dist + 1})
					visited[coord] = dist + 1
				}
			}
		}
		i++
	}
	return -1
}
