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
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

var directions = [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

type Point struct {
	x        int
	y        int
	closest  int
	infinite bool
}

func (p *Point) distanceTo(x, y int) int {
	return utils.AbsInt(p.x-x) + utils.AbsInt(p.y-y)
}

func parseInput(data []string) [][2]int {
	points := make([][2]int, len(data))
	pointsRegex := regexp.MustCompile(`(\d+), (\d+)`)
	for i, line := range data {
		var point [2]int
		match := pointsRegex.FindStringSubmatch(line)
		for j := range 2 {
			if n, err := strconv.Atoi(match[j+1]); err != nil {
				panic(err)
			} else {
				point[j] = n
			}
		}
		points[i] = point
	}
	return points
}

func part1(puzzleInput [][2]int) int {
	xMin := math.MaxInt
	xMax := math.MinInt
	yMin := xMin
	yMax := xMax
	points := make([]*Point, len(puzzleInput))
	for i, point := range puzzleInput {
		if point[0] < xMin {
			xMin = point[0]
		}
		if point[0] > xMax {
			xMax = point[0]
		}
		if point[1] < yMin {
			yMin = point[1]
		}
		if point[1] > yMax {
			yMax = point[1]
		}
		points[i] = &Point{point[0], point[1], 0, false}
	}
	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			d := math.MaxInt
			isTied := false
			var winner *Point
			for _, point := range points {
				curr := point.distanceTo(x, y)
				if curr == d {
					isTied = true
				} else if curr < d {
					d = curr
					isTied = false
					winner = point
				}
			}
			if !isTied {
				winner.closest++
				if x == xMin || x == xMax || y == yMin || y == yMax {
					winner.infinite = true
				}
			}
		}
	}
	res := 0
	for _, point := range points {
		if !point.infinite && point.closest > res {
			res = point.closest
		}
	}
	return res
}

func withinRegion(x, y int, points *[]*Point, limit int) bool {
	sum := 0
	for _, point := range *points {
		sum += point.distanceTo(x, y)
		if sum >= limit {
			return false
		}
	}
	return true
}

func part2(puzzleInput [][2]int) int {
	xAvg := 0
	yAvg := 0
	points := make([]*Point, len(puzzleInput))
	limit := 10000
	for i, point := range puzzleInput {
		xAvg += point[0]
		yAvg += point[1]
		points[i] = &Point{point[0], point[1], 0, false}
	}
	xAvg /= len(puzzleInput)
	yAvg /= len(puzzleInput)
	res := 0
	queue := [][2]int{{xAvg, yAvg}}
	visited := make(map[[2]int]struct{})
	visited[[2]int{xAvg, yAvg}] = struct{}{}
	for len(queue) > 0 {
		coord := queue[0]
		queue = queue[1:]
		if withinRegion(coord[0], coord[1], &points, limit) {
			res++
			for _, d := range directions {
				point := [2]int{coord[0] + d[0], coord[1] + d[1]}
				if _, ok := visited[point]; ok {
					continue
				}
				visited[point] = struct{}{}
				queue = append(queue, point)
			}
		}
	}
	return res
}
