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
	puzzleInput := strings.Split(string(data), "\n")
	fmt.Println(part1(puzzleInput))
}

func part1(puzzleInput []string) int {
	antinodeMap := make(map[[2]int]struct{})
	antennaMap := make(map[byte][][2]int)
	rMax := len(puzzleInput) - 1
	cMax := len(puzzleInput[0]) - 1
	for r := range puzzleInput {
		for c := range puzzleInput[r] {
			if puzzleInput[r][c] != '.' {
				antennaMap[puzzleInput[r][c]] = append(antennaMap[puzzleInput[r][c]], [2]int{r, c})
			}
		}
	}
	for _, pairs := range antennaMap {
		for i, antenna1 := range pairs {
			for _, antenna2 := range pairs[i+1:] {
				dr := antenna2[0] - antenna1[0]
				dc := antenna2[1] - antenna1[1]
				r1, c1 := antenna1[0]-dr, antenna1[1]-dc
				r2, c2 := antenna2[0]+dr, antenna2[1]+dc
				if r1 >= 0 && c1 >= 0 && r1 <= rMax && c1 <= cMax {
					antinodeMap[[2]int{r1, c1}] = struct{}{}
				}
				if r2 >= 0 && c2 >= 0 && r2 <= rMax && c2 <= cMax {
					antinodeMap[[2]int{r2, c2}] = struct{}{}
				}
			}
		}
	}
	return len(antinodeMap)
}
