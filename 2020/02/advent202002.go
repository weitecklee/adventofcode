package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	r, _ := regexp.Compile("(\\d+)-(\\d+) (\\w): (\\w+)")
	valid := 0
	for _, line := range input {
		match := r.FindStringSubmatch(line)
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		n := 0
		for i := range match[4] {
			if match[4][i] == match[3][0] {
				n++
			}
		}
		if n >= a && n <= b {
			valid++
		}
	}
	return valid
}

func part2(input []string) int {
	r, _ := regexp.Compile("(\\d+)-(\\d+) (\\w): (\\w+)")
	valid := 0
	for _, line := range input {
		match := r.FindStringSubmatch(line)
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		if (match[4][a-1] == match[3][0] && match[4][b-1] != match[3][0]) || (match[4][a-1] != match[3][0] && match[4][b-1] == match[3][0]) {
			valid++
		}
	}
	return valid

}
