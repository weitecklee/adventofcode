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
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
}

func parseInput(input []string) (int, [2]int) {
	re := regexp.MustCompile(`\d+`)
	parts0 := re.FindAllString(input[0], -1)
	depth, _ := strconv.Atoi(parts0[0])
	parts1 := re.FindAllString(input[1], -1)
	target0, _ := strconv.Atoi(parts1[0])
	target1, _ := strconv.Atoi(parts1[1])
	return depth, [2]int{target0, target1}
}

func calculateGI(coord [2]int, geologicIndex *map[[2]int]int, erosionLevel *map[[2]int]int, depth int) {
	if coord[0] == 0 {
		(*geologicIndex)[coord] = coord[1] * 48271
	} else if coord[1] == 0 {
		(*geologicIndex)[coord] = coord[0] * 16807
	} else {
		coord1 := [2]int{coord[0] - 1, coord[1]}
		coord2 := [2]int{coord[0], coord[1] - 1}
		(*geologicIndex)[coord] = (*erosionLevel)[coord1] * (*erosionLevel)[coord2]
	}
	(*erosionLevel)[coord] = ((*geologicIndex)[coord] + depth) % 20183
}

func part1(depth int, target [2]int) int {
	geologicIndex := map[[2]int]int{}
	geologicIndex[[2]int{0, 0}] = 0
	erosionLevel := map[[2]int]int{}
	for i := 0; i <= target[0]; i++ {
		for j := 0; j <= target[1]; j++ {
			calculateGI([2]int{i, j}, &geologicIndex, &erosionLevel, depth)
		}
	}
	geologicIndex[target] = 0
	erosionLevel[target] = erosionLevel[[2]int{0, 0}]
	sum := 0
	for i := 0; i <= target[0]; i++ {
		for j := 0; j <= target[1]; j++ {
			sum += erosionLevel[[2]int{i, j}] % 3
		}
	}
	return sum
}
