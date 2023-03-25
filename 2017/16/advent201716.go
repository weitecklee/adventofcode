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
	input := strings.Split(string(data), ",")
	fmt.Println(part1(input))
}

func part1(input []string) string {
	n := 16
	dancers := ""
	for i := 0; i < n; i++ {
		dancers += string(i + 97)
	}
	for _, step := range input {
		dancers2 := ""
		if step[0] == 's' {
			spin, _ := strconv.Atoi(step[1:])
			dancers2 = dancers[(n-spin):] + dancers[:(n-spin)]
		} else if step[0] == 'x' {
			swaps := strings.Split(step[1:], "/")
			swap1, _ := strconv.Atoi(swaps[0])
			swap2, _ := strconv.Atoi(swaps[1])
			for i := 0; i < n; i++ {
				if i == swap1 {
					dancers2 += string(dancers[swap2])
				} else if i == swap2 {
					dancers2 += string(dancers[swap1])
				} else {
					dancers2 += string(dancers[i])
				}
			}
		} else {
			swaps := strings.Split(step[1:], "/")
			for i := 0; i < n; i++ {
				if string(dancers[i]) == swaps[0] {
					dancers2 += swaps[1]
				} else if string(dancers[i]) == swaps[1] {
					dancers2 += swaps[0]
				} else {
					dancers2 += string(dancers[i])
				}
			}
		}
		dancers = dancers2
	}
	return dancers
}
