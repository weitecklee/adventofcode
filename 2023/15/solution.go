package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
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
	puzzleInput := strings.Split(string(data), ",")
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

type Lens struct {
	label       string
	focalLength int
}

var lensRegex = regexp.MustCompile(`(\w+)([\-=])(\d?)`)

func hashString(s string) int {
	curr := 0
	for _, c := range s {
		curr += int(c)
		// curr *= 17
		// curr %= 256
		curr = ((curr << 4) + curr) & 255
	}
	return curr
}

func part1(puzzleInput []string) int {
	res := 0
	for _, s := range puzzleInput {
		res += hashString(s)
	}
	return res
}

func parseStep(s string) (string, byte, int) {
	match := lensRegex.FindStringSubmatch(s)
	if match == nil {
		panic(fmt.Sprintf("Error matching regex with step: %s", s))
	}
	if len(match[3]) > 0 {
		n, err := strconv.Atoi(match[3])
		if err != nil {
			panic(err)
		}
		return match[1], match[2][0], n
	}
	return match[1], match[2][0], -1
}

func part2(puzzleInput []string) int {
	var hashmap [256][]*Lens
	var box []*Lens
	var label string
	var op byte
	var boxNum, focalLength int
loop:
	for _, s := range puzzleInput {
		label, op, focalLength = parseStep(s)
		boxNum = hashString(label)
		box = hashmap[boxNum]
		if op == '-' {
			for i, lens := range box {
				if lens.label == label {
					hashmap[boxNum] = slices.Delete(box, i, i+1)
					break
				}
			}
		} else if op == '=' {
			for _, lens := range box {
				if lens.label == label {
					lens.focalLength = focalLength
					continue loop
				}
			}
			hashmap[boxNum] = append(box, &Lens{label, focalLength})
		}
	}

	res := 0
	for i, box := range hashmap {
		for j, lens := range box {
			res += (i + 1) * (j + 1) * lens.focalLength
		}
	}
	return res
}
