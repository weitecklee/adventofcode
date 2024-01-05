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
	fmt.Println(solve(string(data)))
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

func solve(input string) (int, int) {
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}
	part1 := len(input)
	for i := 40; i < 50; i++ {
		input = lookAndSay(input)
	}
	return part1, len(input)
}
