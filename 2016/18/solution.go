package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	input := string(data)
	fmt.Println(solve(input))
}

func trapCheck(tiles string) string {
	switch tiles {
	case "^^.":
		return "^"
	case ".^^":
		return "^"
	case "^..":
		return "^"
	case "..^":
		return "^"
	default:
		return "."
	}
}

func safeCheck(tiles string) int {
	n := 0
	for _, c := range tiles {
		if c == '.' {
			n++
		}
	}
	return n
}

func solve(row string) (int, int) {
	safeTiles := safeCheck(row)
	for i := 1; i < 40; i++ {
		row2 := ""
		rowCheck := "." + row + "."
		for j := 0; j < len(rowCheck)-2; j++ {
			row2 += trapCheck(rowCheck[j : j+3])
		}
		row = row2
		safeTiles += safeCheck(row)
	}
	safeTiles2 := safeTiles
	for i := 40; i < 400000; i++ {
		row2 := ""
		rowCheck := "." + row + "."
		for j := 0; j < len(rowCheck)-2; j++ {
			row2 += trapCheck(rowCheck[j : j+3])
		}
		row = row2
		safeTiles2 += safeCheck(row)
	}
	return safeTiles, safeTiles2
}
