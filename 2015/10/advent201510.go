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
	fmt.Println(part1(string(data)))
}

func lookAndSay(input string) string {
	var curr byte
	i := 0
	res := ""
	for i < len(input) {
		curr = input[i]
		n := 0
		for i < len(input) && input[i] == curr {
			n++
			i++
		}
		res += strconv.Itoa(n) + string(curr)
	}
	return res
}

func part1(input string) int {
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}
	return len(input)
}
