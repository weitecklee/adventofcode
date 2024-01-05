package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	board, path, rowRange, colRange := parseInput(input)
	fmt.Println(part1(board, path, rowRange, colRange))
	fmt.Println(part2(board, path, rowRange, colRange))
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

func nextOnCube(pos [3]int, rowRange [][2]int, colRange [][2]int, sideLen int) [3]int {
	if pos[0] < 0 {
		if pos[1] < 3*sideLen {
			pos[1] = 3*sideLen - pos[1] - 1
			pos[0] = sideLen
			pos[2] = 0
		} else {
			pos[0] = pos[1] - 2*sideLen
			pos[1] = 0
			pos[2] = 1
		}
	} else if pos[1] < 0 {
		if pos[0] < 2*sideLen {
			pos[1] = pos[0] + 2*sideLen
			pos[0] = 0
			pos[2] = 0
		} else {
			pos[0] = pos[0] - 2*sideLen
			pos[1] = 4*sideLen - 1
			pos[2] = 3
		}
	} else if pos[0] >= 3*sideLen {
		pos[0] = 2*sideLen - 1
		pos[1] = 3*sideLen - pos[1] - 1
		pos[2] = 2
	} else if pos[1] >= 4*sideLen {
		pos[1] = 0
		pos[0] = pos[0] + 2*sideLen
		pos[2] = 1
	} else if pos[0] < rowRange[pos[1]][0] && pos[2] == 2 {
		if pos[1] < sideLen {
			pos[0] = 0
			pos[1] = 3*sideLen - pos[1] - 1
			pos[2] = 0
		} else {
			pos[0] = pos[1] - sideLen
			pos[1] = 2 * sideLen
			pos[2] = 1
		}
	} else if pos[0] > rowRange[pos[1]][1] && pos[2] == 0 {
		if pos[1] < 2*sideLen {
			pos[0] = pos[1] + sideLen
			pos[1] = sideLen - 1
			pos[2] = 3
		} else if pos[1] < 3*sideLen {
			pos[0] = 3*sideLen - 1
			pos[1] = 3*sideLen - pos[1] - 1
			pos[2] = 2
		} else {
			pos[0] = pos[1] - 2*sideLen
			pos[1] = 3*sideLen - 1
			pos[2] = 3
		}
	} else if pos[1] < colRange[pos[0]][0] && pos[2] == 3 {
		pos[1] = pos[0] + sideLen
		pos[0] = sideLen
		pos[2] = 0
	} else if pos[1] > colRange[pos[0]][1] && pos[2] == 1 {
		if pos[0] < 2*sideLen {
			pos[1] = pos[0] + 2*sideLen
			pos[0] = sideLen - 1
			pos[2] = 2
		} else {
			pos[1] = pos[0] - sideLen
			pos[0] = 2*sideLen - 1
			pos[2] = 2
		}
	}
	return pos
}

func part2(board []string, path []any, rowRange [][2]int, colRange [][2]int) int {
	sideLen := len(board[0]) / 3
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
				case 1:
					next[1]++
				case 2:
					next[0]--
				case 3:
					next[1]--
				}
				next = nextOnCube(next, rowRange, colRange, sideLen)
				if board[next[1]][next[0]] == '#' {
					break
				}
				pos = next
			}
		}
	}
	return 1000*(pos[1]+1) + 4*(pos[0]+1) + pos[2]
}
