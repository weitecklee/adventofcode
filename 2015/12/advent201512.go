package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(string(data)))
}

func part1(input string) int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(input, -1)
	res := 0
	for _, match := range matches {
		n, _ := strconv.Atoi(match)
		res += n
	}
	return res
}
