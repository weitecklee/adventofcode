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
	input := strings.Split(string(data), ",")
	fmt.Println(part1(input))
}

func part1(input []string) int {
	pos := [2]float64{0, 0}
	for _, step := range input {
		switch step {
		case "n":
			pos[1] += 2
		case "ne":
			pos[0]++
			pos[1]++
		case "nw":
			pos[0]--
			pos[1]++
		case "s":
			pos[1] -= 2
		case "se":
			pos[0]++
			pos[1]--
		case "sw":
			pos[0]--
			pos[1]--
		default:
			panic("What is this? " + step)
		}
	}
	a := int(math.Abs(pos[0]))
	b := int(math.Abs(pos[1]))
	if a > b {
		return (a-b)/2 + b
	} else {
		return (b-a)/2 + a
	}
}
