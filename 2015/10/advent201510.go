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
	fmt.Println(part1(string(data)))
}

func lookAndSay(input string) string {
	var curr byte
	i := 0
	var res strings.Builder
	for i < len(input) {
		curr = input[i]
		n := 0
		for i < len(input) && input[i] == curr {
			n++
			i++
		}
		res.WriteString(strconv.Itoa(n))
		res.WriteByte(curr)
	}
	return res.String()
}

func part1(input string) int {
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}
	return len(input)
}
