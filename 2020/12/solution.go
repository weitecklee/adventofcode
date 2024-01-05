package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type Instruction struct {
	dir  string
	dist float64
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
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
		instruc.dist, _ = strconv.ParseFloat(matches[2], 64)
		parsedInput = append(parsedInput, instruc)
	}
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
}

func turn(dir float64, angle float64, facing [2]float64) [2]float64 {
	facing2 := [2]float64{}
	facing2[0] = facing[0]*math.Cos(angle*dir/180*math.Pi) + facing[1]*math.Sin(angle*dir/180*math.Pi)
	facing2[1] = facing[1]*math.Cos(angle*dir/180*math.Pi) - facing[0]*math.Sin(angle*dir/180*math.Pi)
	return facing2
}

func part1(input []Instruction) int {
	pos := [2]float64{0, 0}
	facing := [2]float64{1, 0}
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
	return int(math.Abs((pos[0])) + math.Abs(pos[1]))
}

func part2(input []Instruction) int {
	pos := [2]float64{0, 0}
	waypoint := [2]float64{10, 1}
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
	return int(math.Abs((pos[0])) + math.Abs(pos[1]))
}
