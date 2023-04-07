package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
	// fmt.Println(parseInput(input))
}

func notOpenSpace(r rune) bool {
	return r != ' '
}

func parseInput(input []string) ([]string, []any, [][2]int, [][2]int) {
	board := input[:len(input)-2]
	rowRange := [][2]int{}
	colRange := [][2]int{}
	lastCol := 0
	for _, row := range board {
		if len(row)-1 > lastCol {
			lastCol = len(row) - 1
		}
		a := strings.IndexFunc(row, notOpenSpace)
		b := strings.LastIndexFunc(row, notOpenSpace)
		rng := [2]int{a, b}
		rowRange = append(rowRange, rng)
	}

	for i := 0; i <= lastCol; i++ {
		rng := [2]int{0, 0}
		j := 0
		for j < len(board) && i < len(board[j]) && board[j][i] == ' ' {
			j++
		}
		rng[0] = j
		for j < len(board) && i < len(board[j]) && board[j][i] != ' ' {
			j++
		}
		rng[1] = j - 1
		colRange = append(colRange, rng)
	}
	path := []any{}
	curr := 0
	for _, c := range input[len(input)-1] {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			path = append(path, curr, string(c))
			curr = 0
		} else {
			curr = curr*10 + n
		}
	}
	path = append(path, curr)
	return board, path, rowRange, colRange
}

func part1(board []string, path []any, rowRange [][2]int, colRange [][2]int) int {
	pos := [3]int{0, 0, 0}
	pos[0] = rowRange[0][0]
	for _, step := range path {
		if fmt.Sprintf("%T", step) == "string" {
			if step == "R" {
				pos[2]++
				if pos[2] == 4 {
					pos[2] = 0
				}
			} else {
				pos[2]--
				if pos[2] < 0 {
					pos[2] = 3
				}
			}
		} else {
			for i := 0; i < step.(int); i++ {
				next := pos
				switch pos[2] {
				case 0:
					next[0]++
					if next[0] > rowRange[next[1]][1] {
						next[0] = rowRange[next[1]][0]
					}
				case 1:
					next[1]++
					if next[1] > colRange[next[0]][1] {
						next[1] = colRange[next[0]][0]
					}
				case 2:
					next[0]--
					if next[0] < rowRange[next[1]][0] {
						next[0] = rowRange[next[1]][1]
					}
				case 3:
					next[1]--
					if next[1] < colRange[next[0]][0] {
						next[1] = colRange[next[0]][1]
					}
				}
				if board[next[1]][next[0]] == '#' {
					break
				}
				pos = next
			}
		}
	}
	return 1000*(pos[1]+1) + 4*(pos[0]+1) + pos[2]
}
