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

func part1(input int) string {
	cells := [300][300]int{}
	for i := 1; i <= 300; i++ {
		for j := 1; j <= 300; j++ {
			rackID := i + 10
			power := rackID*j + input
			power *= rackID
			power /= 100
			power %= 10
			cells[j-1][i-1] = power - 5
		}
	}
	maxPower := 0
	pos := [2]int{}
	for i := 0; i < 297; i++ {
		for j := 0; j < 297; j++ {
			power := 0
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					power += cells[j+n][i+m]
				}
			}
			if power > maxPower {
				maxPower = power
				pos[0] = i + 1
				pos[1] = j + 1
			}
		}
	}
	return fmt.Sprintf("%d,%d", pos[0], pos[1])
}
