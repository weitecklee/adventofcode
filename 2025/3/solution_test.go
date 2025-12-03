package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func BenchmarkLinkedList(b *testing.B) {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	for i := 0; i < b.N; i++ {
		part1(puzzleInput)
		part2(puzzleInput)
	}
}

func BenchmarkSlidingWindow(b *testing.B) {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	for i := 0; i < b.N; i++ {
		part1_2(puzzleInput)
		part2_2(puzzleInput)
	}
}
