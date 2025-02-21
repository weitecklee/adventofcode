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
	puzzleInput := strings.Split((string(data)), "\n\n")
	rules, pageUpdates := parseInput(puzzleInput)
	part1Answer, incorrectPages := part1(rules, pageUpdates)
	fmt.Println(part1Answer)
	fmt.Println(part2(rules, incorrectPages))
}

func convertToInts(data []string) [][]int {
	res := make([][]int, len(data))
	numRegex := regexp.MustCompile(`\d+`)
	for i, s := range data {
		nums := numRegex.FindAllString(s, -1)
		for _, num := range nums {
			if n, err := strconv.Atoi(num); err != nil {
				panic(err)
			} else {
				res[i] = append(res[i], n)
			}
		}
	}
	return res
}

func parseInput(data []string) (map[int]map[int]struct{}, [][]int) {
	ruleInput := strings.Split(data[0], "\n")
	pageInput := strings.Split(data[1], "\n")
	ruleList := convertToInts(ruleInput)
	rules := make(map[int]map[int]struct{})
	for _, row := range ruleList {
		if _, ok := rules[row[0]]; !ok {
			rules[row[0]] = make(map[int]struct{})
		}
		rules[row[0]][row[1]] = struct{}{}
	}
	return rules, convertToInts(pageInput)
}

func isCorrectOrder(pages []int, rules map[int]map[int]struct{}) bool {
	pageMap := make(map[int]int, len(pages))
	for i, j := range pages {
		pageMap[j] = i
	}
	for i, page1 := range pages {
		if rule, ok := rules[page1]; ok {
			for page2 := range rule {
				if j, ok := pageMap[page2]; ok && j < i {
					return false
				}
			}
		}
	}
	return true
}

func part1(rules map[int]map[int]struct{}, pageUpdates [][]int) (int, [][]int) {
	res := 0
	incorrectPages := make([][]int, 0, len(pageUpdates))
	for _, pages := range pageUpdates {
		if isCorrectOrder(pages, rules) {
			res += pages[len(pages)/2]
		} else {
			incorrectPages = append(incorrectPages, pages)
		}
	}
	return res, incorrectPages
}

func part2(rules map[int]map[int]struct{}, pageUpdates [][]int) int {
	res := 0
	for _, pages := range pageUpdates {
		slices.SortFunc(pages, func(a, b int) int {
			if _, ok := rules[a]; ok {
				if _, ok := rules[a][b]; ok {
					return -1
				}
			}
			if _, ok := rules[b]; ok {
				if _, ok := rules[b][a]; ok {
					return 1
				}
			}
			return 0
		})
		res += pages[len(pages)/2]
	}
	return res
}
