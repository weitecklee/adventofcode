package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	points := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(solve(points))
}

func parseInput(data []string) [][4]int {
	points := make([][4]int, len(data))
	for i, line := range data {
		points[i] = [4]int(utils.ExtractIntsSigned(line))
	}
	return points
}

func distanceBetweenPoints(p1, p2 [4]int) int {
	res := 0
	for i := range 4 {
		res += utils.AbsInt(p1[i] - p2[i])
	}
	return res
}

func solve(points [][4]int) int {
	uf := utils.NewUnionFind(len(points))

	for i, p1 := range points {
		for j, p2 := range points[i+1:] {
			if distanceBetweenPoints(p1, p2) <= 3 {
				uf.Union(i, i+j+1)
			}
		}
	}

	return uf.Count()
}
