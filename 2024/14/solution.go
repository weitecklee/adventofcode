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
	fmt.Println(part1(parseInput(strings.Split(string(data), "\n"))))
}

const (
	WIDTH  = 101
	HEIGHT = 103
	RMAX   = HEIGHT - 1
	CMAX   = WIDTH - 1
	RHALF  = RMAX / 2
	CHALF  = CMAX / 2
)

type Robot struct {
	p [2]int
	v [2]int
}

var robotRegex = regexp.MustCompile(`^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$`)

func NewRobot(s string) *Robot {
	match := robotRegex.FindStringSubmatch(s)
	if match == nil {
		panic(fmt.Sprintf("Error matching robot regex with: %s", s))
	}
	var nums [4]int
	for i := range 4 {
		n, err := strconv.Atoi(match[i+1])
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return &Robot{[2]int{nums[0], nums[1]}, [2]int{nums[2], nums[3]}}
}

func (r *Robot) Quadrant() int {
	if r.p[0] < CHALF {
		if r.p[1] < RHALF {
			return 1
		} else if r.p[1] > RHALF {
			return 2
		}
	} else if r.p[0] > CHALF {
		if r.p[1] < RHALF {
			return 3
		} else if r.p[1] > RHALF {
			return 4
		}
	}
	return 0
}

func (r *Robot) Simulate(k int) {
	r.p[0] = (r.p[0] + k*r.v[0]) % WIDTH
	if r.p[0] < 0 {
		r.p[0] += WIDTH
	}
	r.p[1] = (r.p[1] + k*r.v[1]) % HEIGHT
	if r.p[1] < 0 {
		r.p[1] += HEIGHT
	}
}

func parseInput(puzzleInput []string) []*Robot {
	robots := make([]*Robot, len(puzzleInput))
	for i, s := range puzzleInput {
		robots[i] = NewRobot(s)
	}
	return robots
}

func part1(robots []*Robot) int {
	var quadrants [5]int
	for _, robot := range robots {
		robot.Simulate(100)
		quadrants[robot.Quadrant()]++
	}
	res := 1
	for i := range 4 {
		res *= quadrants[i+1]
	}
	return res
}
