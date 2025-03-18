package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
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

var hypernetRegex = regexp.MustCompile(`\[(.*?)\]`)

func supportsAbba(s string) bool {
	for i := range len(s) - 3 {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

func supportsTls(s string) bool {
	hypernets := hypernetRegex.FindAllString(s, -1)
	if slices.ContainsFunc(hypernets, func(s string) bool { return supportsAbba(s) }) {
		return false
	}
	leftover := hypernetRegex.ReplaceAllString(s, "-")
	return supportsAbba(string(leftover))
}

func findAbaSequences(s string) []string {
	var res []string
	for i := range len(s) - 2 {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			res = append(res, s[i:i+3])
		}
	}
	return res
}

func hasBab(s, aba string) bool {
	for i := range len(s) - 2 {
		if s[i] == aba[1] && s[i+1] == aba[0] && s[i+2] == aba[1] {
			return true
		}
	}
	return false
}

func supportsSsl(s string) bool {
	supernets := hypernetRegex.ReplaceAllString(s, "-+_=")
	abaSequences := findAbaSequences(supernets)
	var sb strings.Builder
	matches := hypernetRegex.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		sb.WriteString(match[1])
		sb.WriteString("-+_=")
	}
	hypernets := sb.String()
	for _, aba := range abaSequences {
		if hasBab(hypernets, aba) {
			return true
		}
	}
	return false
}

func part1(puzzleInput []string) int {
	res := 0
	for _, s := range puzzleInput {
		if supportsTls(s) {
			res++
		}
	}
	return res
}

func part2(puzzleInput []string) int {
	res := 0
	for _, s := range puzzleInput {
		if supportsSsl(s) {
			res++
		}
	}
	return res
}
