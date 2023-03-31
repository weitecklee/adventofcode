package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	cells := parseInput(string(data))
	fmt.Println(part1(cells))
	fmt.Println(part2(cells))
}

func parseInput(input string) *[300][300]int {
	serial, _ := strconv.Atoi(input)
	cells := [300][300]int{}
	for i := 1; i <= 300; i++ {
		for j := 1; j <= 300; j++ {
			rackID := i + 10
			power := rackID*j + serial
			power *= rackID
			power /= 100
			power %= 10
			cells[j-1][i-1] = power - 5
		}
	}
	return &cells
}

func part1(cells *[300][300]int) string {
	maxPower := 0
	pos := [2]int{}
	for i := 0; i < 297; i++ {
		for j := 0; j < 297; j++ {
			power := 0
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					power += (*cells)[j+n][i+m]
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

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

func part2(cells *[300][300]int) string {
	defer duration(track("part2"))
	maxPower := 0
	dims := [3]int{}
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			for k := 1; k+i <= 300 && k+j <= 300; k++ {
				power := 0
				for m := 0; m < k; m++ {
					for n := 0; n < k; n++ {
						power += (*cells)[j+n][i+m]
					}
				}
				if power > maxPower {
					maxPower = power
					dims[0] = i + 1
					dims[1] = j + 1
					dims[2] = k
				}
			}
		}
	}
	return fmt.Sprintf("%d,%d,%d", dims[0], dims[1], dims[2])
}
