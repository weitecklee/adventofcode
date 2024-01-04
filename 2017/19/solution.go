package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
}

func part1(input []string) (string, int) {
	pos := [2]int{0, 0}
	pos[1] = strings.Index(input[0], "|")
	dir := [2]int{1, 0}
	res := ""
	re := regexp.MustCompile(`[A-Z]`)
	steps := 0
loop:
	for {
		steps++
		pos[0] += dir[0]
		pos[1] += dir[1]
		curr := string(input[pos[0]][pos[1]])
		if re.MatchString(curr) {
			res += curr
			continue
		}
		switch curr {
		case "+":
			if dir[0] == 0 {
				dir[1] = 0
				if pos[0] > 0 && pos[1] < len(input[pos[0]-1]) && input[pos[0]-1][pos[1]] == '|' {
					dir[0] = -1
				} else {
					dir[0] = 1
				}
			} else {
				dir[0] = 0
				if pos[1] > 0 && input[pos[0]][pos[1]-1] == '-' {
					dir[1] = -1
				} else {
					dir[1] = 1
				}
			}
		case " ":
			break loop
		}
	}
	return res, steps
}
