package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/weitecklee/adventofcode/utils"
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
	list1, list2 := parseInput(&puzzleInput)
	fmt.Println(part1(list1, list2))
	fmt.Println(part2(list1, list2))
}

func parseInput(puzzleInput *[]string) (*[]int, *[]int) {
	list1 := make([]int, len(*puzzleInput))
	list2 := make([]int, len(*puzzleInput))
	numRegex := regexp.MustCompile(`^(\d+)\s+(\d+)$`)
	for i, s := range *puzzleInput {
		matches := numRegex.FindStringSubmatch(s)
		if n, err := strconv.Atoi(matches[1]); err != nil {
			panic(err)
		} else {
			list1[i] = n
		}
		if n, err := strconv.Atoi(matches[2]); err != nil {
			panic(err)
		} else {
			list2[i] = n
		}
	}
	sort.Ints(list1)
	sort.Ints(list2)
	return &list1, &list2
}

func part1(list1, list2 *[]int) int {
	res := 0
	for i := range *list1 {
		res += utils.AbsInt((*list1)[i] - (*list2)[i])
	}
	return res
}

func part2(list1, list2 *[]int) int {
	res := 0
	i := 0
	j := 0
	for i < len(*list1) && j < len(*list2) {
		for i < len(*list1) && (*list1)[i] < (*list2)[j] {
			i++
		}
		for j < len(*list2) && (*list1)[i] > (*list2)[j] {
			j++
		}
		for j < len(*list2) && (*list1)[i] == (*list2)[j] {
			res += (*list1)[i]
			j++
		}
	}
	return res
}
