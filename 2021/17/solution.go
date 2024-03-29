package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	parsedInput := parseInput(string(data))
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
}

func parseInput(input string) []int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(input, -1)
	nums := []int{}
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		nums = append(nums, num)
	}
	return nums
}

func fireAway(xVel int, yVel int, xMin int, xMax int, yMin int, yMax int) int {
	x := 0
	y := 0
	maxHt := 0
	for {
		x += xVel
		y += yVel
		if y > maxHt {
			maxHt = y
		}
		if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
			return maxHt
		}
		if x > xMax {
			return -1
		}
		if xVel == 0 && y < yMin {
			return -1
		}
		if xVel > 0 {
			xVel--
		} else if xVel < 0 {
			xVel++
		}
		yVel--
	}
}

func part1(input []int) int {
	xMin := input[0]
	xMax := input[1]
	yMin := input[2]
	yMax := input[3]
	maxHt := -1
	// wasted too much time pre-optimizing, just search the area with bounds of 1000
	for xVel := 1; xVel < 1000; xVel++ {
		for yVel := 1; yVel < 1000; yVel++ {
			currMaxHt := fireAway(xVel, yVel, xMin, xMax, yMin, yMax)
			// if currMaxHt >= 0 {
			// 	fmt.Println(xVel, yVel, currMaxHt)
			// }
			if currMaxHt > maxHt {
				maxHt = currMaxHt
			}
		}
	}
	return maxHt
}

func part2(input []int) int {
	xMin := input[0]
	xMax := input[1]
	yMin := input[2]
	yMax := input[3]
	count := 0
	// do the same thing as part 1 but yVel can be negative now
	for xVel := 1; xVel < 1000; xVel++ {
		for yVel := -1000; yVel < 1000; yVel++ {
			currMaxHt := fireAway(xVel, yVel, xMin, xMax, yMin, yMax)
			if currMaxHt >= 0 {
				// fmt.Println(xVel, yVel, currMaxHt)
				count++
			}
		}
	}
	return count
}
