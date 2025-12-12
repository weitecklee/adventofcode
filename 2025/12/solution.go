package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/weitecklee/adventofcode/utils"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	presents, regions := parseInput(strings.Split(string(data), "\n\n"))
	fmt.Println(solve(presents, regions))
}

const MAX_PRESENT_AREA = 9

type Present struct {
	shape []string
	area  int
}

func NewPresent(line []string) *Present {
	shape := line[1:]
	var size int
	for _, line := range shape {
		for _, ch := range line {
			if ch == '#' {
				size += 1
			}
		}
	}
	return &Present{
		shape,
		size,
	}
}

type Region struct {
	dimensions    [2]int
	area          int
	presentCounts []int
}

func NewRegion(line string) *Region {
	nums := utils.ExtractInts(line)
	dimensions := [2]int(nums[:2])
	area := dimensions[0] * dimensions[1]
	presentCounts := nums[2:]
	return &Region{
		dimensions,
		area,
		presentCounts,
	}
}

func parseInput(data []string) ([]*Present, []*Region) {
	presentRegex := regexp.MustCompile(`^\d+:`)
	var presentPart []string
	var regionPart []string
	for _, part := range data {
		if presentRegex.MatchString(part) {
			presentPart = append(presentPart, part)
		} else {
			regionPart = append(regionPart, strings.Split(part, "\n")...)
		}
	}

	presents := make([]*Present, len(presentPart))
	regions := make([]*Region, len(regionPart))
	for i, part := range presentPart {
		presents[i] = NewPresent(strings.Split(part, "\n"))
	}
	for i, part := range regionPart {
		regions[i] = NewRegion(part)
	}

	return presents, regions
}

func solve(presents []*Present, regions []*Region) int {
	res := 0
	for _, region := range regions {
		var minimumAreaNeeded, minimumTrivialArea int
		for i, count := range region.presentCounts {
			minimumAreaNeeded += presents[i].area * count
			minimumTrivialArea += count * MAX_PRESENT_AREA
		}
		if minimumAreaNeeded > region.area {
			continue
		}
		if region.area >= minimumTrivialArea {
			res += 1
		} else {
			panic("Packing problem!")
		}
	}
	return res
}

/*
	So I looked at this problem when it first came up at midnight local time,
	read through the problem, thought "Do they really want me to solve a packing
	problem? It's the last day of AoC! This is supposed to be a victory lap!
	Forget it, I'm going to sleep."
	The next morning, I saw all the memes on the subreddit, got disappointed in
	myself, "you got me again, AoC". Anyway it turns out all the regions are
	trivial: 	either too small to ever fit all the pieces, or are so big that
	they'll	easily fit all the pieces.
*/
