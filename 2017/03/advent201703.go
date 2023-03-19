package main

import (
	"fmt"
	"math"
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
	fmt.Println(part2(input))
}

func part1(input int) int {
	if input == 1 {
		return 0
	}
	n := math.Ceil(math.Sqrt(float64(input)))
	ring := int(math.Ceil((n - 1) / 2))
	side := 2*ring + 1
	offset := side*side - input
	return ring + int(math.Abs(float64(ring-offset%(side-1))))
}

func coord(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func part2(input int) int {
	n := 1
	grid := map[string]int{}
	x := 1
	y := 0
	grid[coord(0, 0)] = n
	dir := [2]int{0, 1}
	ring := 1
	for n <= input {
		n2 := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				n2 += grid[coord(x+i, y+j)]
			}
		}
		n = n2
		grid[coord(x, y)] = n
		x += dir[0]
		y += dir[1]
		if y == ring {
			dir[0] = -1
			dir[1] = 0
		}
		if x == -ring {
			dir[0] = 0
			dir[1] = -1
		}
		if y == -ring {
			dir[0] = 1
			dir[1] = 0
		}
		if x == ring+1 {
			dir[0] = 0
			dir[1] = 1
			ring++
		}
	}
	return n
}
