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
	puzzleInput := string(data)
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func parseMarker(s string) (int, int) {
	parts := strings.Split(s, "x")
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return n, k
}

func decompress(s string, recursive bool) int {
	res := 0
	i := 0
	for i < len(s) {
		if s[i] != '(' {
			res++
		} else {
			j := i + 1
			for s[j] != ')' {
				j++
			}
			n, k := parseMarker(s[i+1 : j])
			if recursive {
				res += decompress(s[j+1:j+n+1], recursive) * k
			} else {
				res += n * k
			}
			i = j + n
		}
		i++
	}
	return res
}

func part1(s string) int {
	return decompress(s, false)
}

func part2(s string) int {
	return decompress(s, true)
}
