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
