package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	dir  string
	dist int
}

func main() {
	data, err := os.ReadFile("input202012.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	re := regexp.MustCompile(`(\w)(\d+)`)
	parsedInput := []Instruction{}
	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		instruc := Instruction{}
		instruc.dir = matches[1]
		instruc.dist, _ = strconv.Atoi(matches[2])
		parsedInput = append(parsedInput, instruc)
	}
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
}

func turn(dir int, angle int, facing [2]int) [2]int {
	if angle == 90 {
		facing2 := [2]int{}
		facing2[0] = facing[1] * dir
		facing2[1] = -facing[0] * dir
		return facing2
	} else if angle == 180 {
		return [2]int{-facing[0], -facing[1]}
	} else {
		facing2 := [2]int{}
		facing2[0] = -facing[1] * dir
		facing2[1] = facing[0] * dir
		return facing2
	}
}

func part1(input []Instruction) int {
	pos := [2]int{0, 0}
	facing := [2]int{1, 0}
	for _, instruc := range input {
		switch instruc.dir {
		case "N":
			pos[1] += instruc.dist
		case "S":
			pos[1] -= instruc.dist
		case "E":
			pos[0] += instruc.dist
		case "W":
			pos[0] -= instruc.dist
		case "L":
			facing = turn(-1, instruc.dist, facing)
		case "R":
			facing = turn(1, instruc.dist, facing)
		case "F":
			pos[0] += instruc.dist * facing[0]
			pos[1] += instruc.dist * facing[1]
		}
	}
	return int(math.Abs(float64((pos[0]))) + math.Abs(float64(pos[1])))
}

func part2(input []Instruction) int {
	pos := [2]int{0, 0}
	waypoint := [2]int{10, 1}
	for _, instruc := range input {
		switch instruc.dir {
		case "N":
			waypoint[1] += instruc.dist
		case "S":
			waypoint[1] -= instruc.dist
		case "E":
			waypoint[0] += instruc.dist
		case "W":
			waypoint[0] -= instruc.dist
		case "L":
			waypoint = turn(-1, instruc.dist, waypoint)
		case "R":
			waypoint = turn(1, instruc.dist, waypoint)
		case "F":
			pos[0] += instruc.dist * waypoint[0]
			pos[1] += instruc.dist * waypoint[1]
		}
	}
	return int(math.Abs(float64((pos[0]))) + math.Abs(float64(pos[1])))
}
