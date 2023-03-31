package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(solve(parseInput(input)))
}

type Point struct {
	pos [2]int
	vel [2]int
}

func (p *Point) move() {
	p.pos[0] += p.vel[0]
	p.pos[1] += p.vel[1]
}

func parseInput(input []string) *[]*Point {
	points := []*Point{}
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range input {
		nums := re.FindAllString(line, -1)
		pos1, _ := strconv.Atoi(nums[0])
		pos2, _ := strconv.Atoi(nums[1])
		vel1, _ := strconv.Atoi(nums[2])
		vel2, _ := strconv.Atoi(nums[3])
		point := Point{
			pos: [2]int{pos1, pos2},
			vel: [2]int{vel1, vel2},
		}
		points = append(points, &point)
	}
	return &points
}

func countIslands(pointMap *map[[2]int]bool) int {
	islands := 0
	for pos, notCounted := range *pointMap {
		if notCounted {
			islands++
			(*pointMap)[pos] = false
			mapIsland(pointMap, pos)
		}
	}
	return islands
}

func mapIsland(pointMap *map[[2]int]bool, pos [2]int) {
	p1 := [2]int{pos[0] - 1, pos[1]}
	p2 := [2]int{pos[0] + 1, pos[1]}
	p3 := [2]int{pos[0], pos[1] - 1}
	p4 := [2]int{pos[0], pos[1] + 1}
	for _, point := range [][2]int{p1, p2, p3, p4} {
		if (*pointMap)[point] {
			(*pointMap)[point] = false
			mapIsland(pointMap, point)
		}
	}
}

func printIslands(pointMap *map[[2]int]bool) {
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for pos := range *pointMap {
		if pos[0] < minX {
			minX = pos[0]
		}
		if pos[0] > maxX {
			maxX = pos[0]
		}
		if pos[1] < minY {
			minY = pos[1]
		}
		if pos[1] > maxY {
			maxY = pos[1]
		}
	}
	ht := maxY - minY
	wd := maxX - minX
	message := [][]string{}
	for i := 0; i <= ht; i++ {
		row := []string{}
		for j := 0; j <= wd; j++ {
			row = append(row, ".")
		}
		message = append(message, row)
	}
	for pos := range *pointMap {
		message[pos[1]-minY][pos[0]-minX] = "#"
	}
	for _, row := range message {
		fmt.Println(strings.Join(row, ""))
	}
}

func solve(points *[]*Point) int {
	minIslands := len(*points)
	timeTaken := 0
	for i := 0; i < 1000000; i++ {
		pointMap := map[[2]int]bool{}
		for _, point := range *points {
			point.move()
			pointMap[point.pos] = true
		}
		islands := countIslands(&pointMap)
		if islands < minIslands {
			minIslands = islands
			timeTaken = i + 1
			if minIslands < 50 { // kinda arbitrary break point, my input had 29 islands
				printIslands(&pointMap)
				break
			}
		}
	}
	return timeTaken
}
