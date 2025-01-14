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

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	cuboids := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(solve(cuboids, -50, 50, -50, 50, -50, 50))
	fmt.Println(solve(cuboids))
}

type Cuboid struct {
	state  bool
	xRange [2]int
	yRange [2]int
	zRange [2]int
}

func parseInput(input []string) []Cuboid {
	cuboids := []Cuboid{}
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range input {
		state := true
		if line[1] == 'f' {
			state = false
		}
		matches := re.FindAllString(line, -1)
		nums := []int{}
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			nums = append(nums, num)
		}
		cuboids = append(cuboids, Cuboid{
			state:  state,
			xRange: [2]int{nums[0], nums[1]},
			yRange: [2]int{nums[2], nums[3]},
			zRange: [2]int{nums[4], nums[5]},
		})
	}
	return cuboids
}

func overlapCuboid(cuboid1, cuboid2 Cuboid) (bool, Cuboid) {
	xMin := max(cuboid1.xRange[0], cuboid2.xRange[0])
	xMax := min(cuboid1.xRange[1], cuboid2.xRange[1])
	if xMin > xMax {
		return false, Cuboid{}
	}
	yMin := max(cuboid1.yRange[0], cuboid2.yRange[0])
	yMax := min(cuboid1.yRange[1], cuboid2.yRange[1])
	if yMin > yMax {
		return false, Cuboid{}
	}
	zMin := max(cuboid1.zRange[0], cuboid2.zRange[0])
	zMax := min(cuboid1.zRange[1], cuboid2.zRange[1])
	if zMin > zMax {
		return false, Cuboid{}
	}
	return true, Cuboid{
		state:  !cuboid2.state,
		xRange: [2]int{xMin, xMax},
		yRange: [2]int{yMin, yMax},
		zRange: [2]int{zMin, zMax},
	}
}

func solve(cuboids []Cuboid, ranges ...int) int {
	xRangeMin, xRangeMax := math.MinInt, math.MaxInt
	yRangeMin, yRangeMax := math.MinInt, math.MaxInt
	zRangeMin, zRangeMax := math.MinInt, math.MaxInt
	if len(ranges) > 0 {
		xRangeMin, xRangeMax = ranges[0], ranges[1]
		yRangeMin, yRangeMax = ranges[2], ranges[3]
		zRangeMin, zRangeMax = ranges[4], ranges[5]
	}
	spaceCuboid := Cuboid{
		state:  false,
		xRange: [2]int{xRangeMin, xRangeMax},
		yRange: [2]int{yRangeMin, yRangeMax},
		zRange: [2]int{zRangeMin, zRangeMax},
	}
	var overlaps []Cuboid
	for _, cuboid1 := range cuboids {
		var tmp []Cuboid
		if cuboid1.state {
			if isOverlap, cbd := overlapCuboid(cuboid1, spaceCuboid); isOverlap {
				tmp = append(tmp, cbd)
			}
		}
		for _, cuboid2 := range overlaps {
			if isOverlap, cbd := overlapCuboid(cuboid1, cuboid2); isOverlap {
				tmp = append(tmp, cbd)
			}
		}
		overlaps = append(overlaps, tmp...)
	}

	res := 0
	for _, cuboid := range overlaps {
		size := (cuboid.xRange[1] - cuboid.xRange[0] + 1) * (cuboid.yRange[1] - cuboid.yRange[0] + 1) * (cuboid.zRange[1] - cuboid.zRange[0] + 1)
		if cuboid.state {
			res += size
		} else {
			res -= size
		}
	}
	return res
}
