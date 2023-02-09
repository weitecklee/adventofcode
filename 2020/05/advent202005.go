package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	data, err := os.ReadFile("input202005.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	seatIDs := []int{}
	for _, line := range input {
		seatIDs = append(seatIDs, calcSeatID((line)))
	}
	sort.Ints(seatIDs)
	fmt.Println(part1(seatIDs))
	fmt.Println(part2(seatIDs))
}

func calcSeatID(line string) int {
	row, col := 0.0, 0.0
	for i := 0; i < 7; i++ {
		if line[i] == "B"[0] {
			row += math.Pow(2, float64(6-i))
		}
	}
	for i := 7; i < 10; i++ {
		if line[i] == "R"[0] {
			col += math.Pow(2, float64(9-i))
		}
	}
	return int(8*row + col)
}

func part1(seatIDs []int) int {
	return seatIDs[len(seatIDs)-1]
}

func part2(seatIDs []int) int {
	for i, n := range seatIDs {
		if seatIDs[i+1] == n+2 {
			return n + 1
		}
	}
	return -1
}
