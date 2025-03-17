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
	puzzleInput := strings.Split(string(data), "\n")
	screen := makeScreen()
	fmt.Println(part1(puzzleInput, screen))
	part2(screen)
}

const (
	ScreenHeight = 6
	ScreenWidth  = 50
)

func rect(screen *[][]bool, n1, n2 int) {
	for i := range n1 {
		for j := range n2 {
			(*screen)[j][i] = true
		}
	}
}

func rotateRow(screen *[][]bool, n1, n2 int) {
	(*screen)[n1] = append((*screen)[n1][ScreenWidth-n2:], (*screen)[n1][:ScreenWidth-n2]...)

}

func rotateColumn(screen *[][]bool, n1, n2 int) {
	tmp := make([]bool, ScreenHeight)
	for row := range ScreenHeight {
		tmp[(row+n2)%ScreenHeight] = (*screen)[row][n1]
	}
	for row := range ScreenHeight {
		(*screen)[row][n1] = tmp[row]
	}
}

func makeScreen() *[][]bool {
	screen := make([][]bool, ScreenHeight)
	for row := range ScreenHeight {
		screen[row] = make([]bool, ScreenWidth)
	}
	return &screen
}

func printScreen(screen *[][]bool) {
	var sb strings.Builder
	for _, row := range *screen {
		sb.Reset()
		for _, ch := range row {
			if ch {
				sb.WriteRune('█')
			} else {
				sb.WriteRune(' ')
			}
		}
		fmt.Println(sb.String())
	}
}

func part1(puzzleInput []string, screen *[][]bool) int {
	instructionRegex := regexp.MustCompile(`^([a-z ]+)\D+(\d+)\D+(\d+)$`)
	for _, line := range puzzleInput {
		match := instructionRegex.FindStringSubmatch(line)
		if match == nil {
			panic(fmt.Sprintf("Error matching instruction regex with : %s", line))
		}
		n1, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(match[3])
		if err != nil {
			panic(err)
		}
		switch match[1] {
		case "rect":
			rect(screen, n1, n2)
		case "rotate row y":
			rotateRow(screen, n1, n2)
		case "rotate column x":
			rotateColumn(screen, n1, n2)
		default:
			panic(fmt.Sprintf("Unrecognized instruction: %s ", match[1]))
		}
	}
	res := 0
	for _, row := range *screen {
		for _, p := range row {
			if p {
				res++
			}
		}
	}
	return res
}

func part2(screen *[][]bool) {
	printScreen(screen)
}
