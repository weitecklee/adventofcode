package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	positions := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(positions))
}

type Instruction struct {
	turnOn bool
	xRange [2]int
	yRange [2]int
	zRange [2]int
}

func parseInput(input []string) []Instruction {
	instructions := []Instruction{}
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range input {
		turnOn := true
		if line[1] == 'f' {
			turnOn = false
		}
		matches := re.FindAllString(line, -1)
		nums := []int{}
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			nums = append(nums, num)
		}
		instructions = append(instructions, Instruction{
			turnOn: turnOn,
			xRange: [2]int{nums[0], nums[1]},
			yRange: [2]int{nums[2], nums[3]},
			zRange: [2]int{nums[4], nums[5]},
		})
	}
	return instructions
}

func part1(instructions []Instruction) int {
	cubes := map[[3]int]bool{}
	for _, instruction := range instructions {
		xMin := -50
		if instruction.xRange[0] > xMin {
			xMin = instruction.xRange[0]
		}
		xMax := 50
		if instruction.xRange[1] < xMax {
			xMax = instruction.xRange[1]
		}
		yMin := -50
		if instruction.yRange[0] > yMin {
			yMin = instruction.yRange[0]
		}
		yMax := 50
		if instruction.yRange[1] < yMax {
			yMax = instruction.yRange[1]
		}
		zMin := -50
		if instruction.zRange[0] > zMin {
			zMin = instruction.zRange[0]
		}
		zMax := 50
		if instruction.zRange[1] < zMax {
			zMax = instruction.zRange[1]
		}
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				for k := zMin; k <= zMax; k++ {
					cubes[[3]int{i, j, k}] = instruction.turnOn
				}
			}
		}
	}
	count := 0
	for _, on := range cubes {
		if on {
			count++
		}
	}
	return count
}
