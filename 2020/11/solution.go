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
	input := strings.Split(string(data), "\n")
	layout := [][]string{}
	for _, line := range input {
		layout = append(layout, strings.Split(line, ""))
	}
	fmt.Println(part1(layout))
	fmt.Println(part2(layout))
}

func isValidSeat(r int, c int, layout [][]string) bool {
	return r >= 0 && c >= 0 && r <= len(layout)-1 && c <= len(layout[0])-1
}

func occupiedAdjacent(r int, c int, layout [][]string) int {
	occupied := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if isValidSeat(r+i, c+j, layout) && layout[r+i][c+j] == "#" && !(i == 0 && j == 0) {
				occupied++
			}
		}
	}
	return occupied
}

func occupiedAdjacent2(r int, c int, layout [][]string) int {
	occupied := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			k := 1
			for isValidSeat(r+k*i, c+k*j, layout) {
				if layout[r+k*i][c+k*j] == "#" {
					occupied++
					break
				} else if layout[r+k*i][c+k*j] == "L" {
					break
				}
				k++
			}
		}
	}
	return occupied
}

func part1(layout [][]string) int {
	changes := 1
	for changes > 0 {
		changes = 0
		layout2 := [][]string{}
		for r, row := range layout {
			row2 := []string{}
			for c, seat := range row {
				if seat == "L" && occupiedAdjacent(r, c, layout) == 0 {
					row2 = append(row2, "#")
					changes++
				} else if seat == "#" && occupiedAdjacent(r, c, layout) >= 4 {
					row2 = append(row2, "L")
					changes++
				} else {
					row2 = append(row2, seat)
				}
			}
			layout2 = append(layout2, row2)
		}
		layout = layout2
	}
	occupied := 0
	for _, row := range layout {
		for _, seat := range row {
			if seat == "#" {
				occupied++
			}
		}
	}
	return occupied
}

func part2(layout [][]string) int {
	changes := 1
	for changes > 0 {
		changes = 0
		layout2 := [][]string{}
		for r, row := range layout {
			row2 := []string{}
			for c, seat := range row {
				if seat == "L" && occupiedAdjacent2(r, c, layout) == 0 {
					row2 = append(row2, "#")
					changes++
				} else if seat == "#" && occupiedAdjacent2(r, c, layout) >= 5 {
					row2 = append(row2, "L")
					changes++
				} else {
					row2 = append(row2, seat)
				}
			}
			layout2 = append(layout2, row2)
		}
		layout = layout2
	}
	occupied := 0
	for _, row := range layout {
		for _, seat := range row {
			if seat == "#" {
				occupied++
			}
		}
	}
	return occupied
}
