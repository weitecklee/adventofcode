package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
	puzzleInput := strings.Split(string(data), "\n")
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

type Keypad struct {
	layout []string
	pos    [2]int
}

func (k *Keypad) CurrentKey() byte {
	return k.layout[k.pos[0]][k.pos[1]]
}

func (k *Keypad) Move(r rune) {
	pos2 := k.pos
	switch r {
	case 'L':
		pos2[1]--
	case 'R':
		pos2[1]++
	case 'U':
		pos2[0]--
	case 'D':
		pos2[0]++
	}
	if pos2[0] >= 0 && pos2[1] >= 0 && pos2[0] < len(k.layout) && pos2[1] < len(k.layout[pos2[0]]) && k.layout[pos2[0]][pos2[1]] != '_' {
		k.pos = pos2
	}
}

func solve(puzzleInput []string, keypad *Keypad) string {
	var sb strings.Builder
	for _, line := range puzzleInput {
		for _, r := range line {
			keypad.Move(r)
		}
		sb.WriteByte(keypad.CurrentKey())
	}
	return sb.String()
}

func part1(puzzleInput []string) string {
	layout := []string{"123", "456", "789"}
	pos := [2]int{1, 1}
	keypad := Keypad{layout, pos}
	return solve(puzzleInput, &keypad)
}

func part2(puzzleInput []string) string {
	layout := []string{"__1__", "_234_", "56789", "_ABC_", "__D__"}
	pos := [2]int{2, 1}
	keypad := Keypad{layout, pos}
	return solve(puzzleInput, &keypad)
}
