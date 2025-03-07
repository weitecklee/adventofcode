package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
	area := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(solve(area))
}

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func parseInput(data []string) [][]int {
	area := make([][]int, len(data))
	for i, line := range data {
		nums := make([]int, len(line))
		for j, s := range line {
			n := s - '0'
			nums[j] = int(n)
		}
		area[i] = nums
	}
	return area
}

func calcBasinSize(area [][]int, lowPoint [2]int) int {
	visited := make(map[[2]int]struct{})
	visited[lowPoint] = struct{}{}
	var queue [][2]int
	var pos, pos2 [2]int
	queue = append(queue, lowPoint)
	for len(queue) > 0 {
		pos = queue[0]
		queue = queue[1:]
		for _, d := range directions {
			pos2[0], pos2[1] = pos[0]+d[0], pos[1]+d[1]
			if pos2[0] < 0 || pos2[1] < 0 || pos2[0] >= len(area) || pos2[1] >= len(area[0]) {
				continue
			}
			if area[pos2[0]][pos2[1]] == 9 {
				continue
			}
			if _, ok := visited[pos2]; ok {
				continue
			}
			visited[pos2] = struct{}{}
			queue = append(queue, pos2)
		}
	}
	return len(visited)
}

func updateTop3(top3 *[3]int, n int) {
	for i := range top3 {
		if n > top3[i] {
			copy(top3[i+1:], top3[i:])
			top3[i] = n
			break
		}
	}
}

func solve(area [][]int) (int, int) {
	var isLowPoint bool
	var r2, c2, part1 int
	var basinSizesTop3 [3]int
	for r, row := range area {
		for c, ht := range row {
			isLowPoint = true
			for _, d := range directions {
				r2, c2 = r+d[0], c+d[1]
				if r2 < 0 || c2 < 0 || r2 >= len(area) || c2 >= len(row) {
					continue
				}
				if area[r2][c2] <= ht {
					isLowPoint = false
					break
				}
			}
			if isLowPoint {
				part1 += ht + 1
				updateTop3(&basinSizesTop3, calcBasinSize(area, [2]int{r, c}))
			}
		}
	}
	part2 := 1
	for _, size := range basinSizesTop3 {
		part2 *= size
	}
	return part1, part2
}
