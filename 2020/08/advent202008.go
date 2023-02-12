package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input202008.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func runCode(input []string, lineToSkip int) (int, bool, []int) {
	re, _ := regexp.Compile(`(\w{3}) ([+-])(\d+)`)
	acc := 0
	executed := map[int]bool{}
	line := 0
	linesToSkip := []int{}
	for line < len(input) && !executed[line] {
		executed[line] = true
		linesToSkip = append(linesToSkip, line)
		instruc := input[line]
		matches := re.FindStringSubmatch(instruc)
		if line == lineToSkip {
			if matches[1] == "nop" {
				matches[1] = "jmp"
			} else {
				matches[1] = "nop"
			}
		}
		if matches[1] == "nop" {
			line++
		} else if matches[1] == "acc" {
			i, _ := strconv.Atoi(matches[3])
			if matches[2] == "+" {
				acc += i
			} else {
				acc -= i
			}
			line++
		} else {
			i, _ := strconv.Atoi(matches[3])
			if matches[2] == "+" {
				line += i
			} else {
				line -= i
			}
		}
	}
	return acc, line < len(input), linesToSkip
}

func part1(input []string) int {
	acc, _, _ := runCode(input, len(input))
	return acc
}

func part2(input []string) int {
	_, _, linesToSkip := runCode(input, len(input))
	for _, lineToSkip := range linesToSkip {
		if !strings.Contains(input[lineToSkip], "acc") {
			acc, repeat, _ := runCode(input, lineToSkip)
			if !repeat {
				return acc
			}
		}
	}
	return -1
}
