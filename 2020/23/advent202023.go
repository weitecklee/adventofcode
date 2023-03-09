package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parsedInput := parseInput(string(data))
	fmt.Println(part1(parsedInput))
}

func parseInput(input string) []int {
	arr := []int{}
	for _, c := range input {
		n, _ := strconv.Atoi(string(c))
		arr = append(arr, n)
	}
	return arr
}

func contains(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}

func move(cups []int) []int {
	curr := cups[0]
	pickup := cups[1:4]
	dest := curr - 1
	if dest == 0 {
		dest = 9
	}
	for contains(pickup, dest) {
		dest--
		if dest == 0 {
			dest = 9
		}
	}
	cups2 := []int{}
	for _, n := range cups[4:] {
		cups2 = append(cups2, n)
		if n == dest {
			cups2 = append(cups2, pickup...)
		}
	}
	return append(cups2, curr)
}

func part1(input []int) string {
	for i := 0; i < 100; i++ {
		input = move(input)
	}
	idx := 0
	for input[idx] != 1 {
		idx++
	}
	ret := ""
	for i := idx + 1; i < len(input); i++ {
		ret += strconv.Itoa(input[i])
	}
	for i := 0; i < idx; i++ {
		ret += strconv.Itoa(input[i])
	}
	return ret
}
