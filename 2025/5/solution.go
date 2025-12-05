package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
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
	ranges, ids := parseInput(strings.Split(string(data), "\n\n"))
	fmt.Println(part1(ranges, ids))
	fmt.Println(part2(ranges))
}

func parseInput(data []string) ([][2]int, []int) {
	data1 := strings.Split(data[0], "\n")
	ranges := make([][2]int, len(data1))
	for i, line := range data1 {
		ss := strings.Split(line, "-")
		n1, _ := strconv.Atoi(ss[0])
		n2, _ := strconv.Atoi(ss[1])
		ranges[i] = [2]int{n1, n2}
	}
	data2 := strings.Split(data[1], "\n")
	ids := make([]int, len(data2))
	for i, s := range data2 {
		n, _ := strconv.Atoi(s)
		ids[i] = n
	}
	slices.SortFunc(ranges, func(a, b [2]int) int {
		return a[0] - b[0]
	})
	return ranges, ids
}

func part1(ranges [][2]int, ids []int) int {
	res := 0
	for _, id := range ids {
		for _, rnge := range ranges {
			if id >= rnge[0] && id <= rnge[1] {
				res++
				break
			}
		}
	}
	return res
}

func part2(ranges [][2]int) int {
	res := 0
	tmp := ranges[0]
	for _, rnge := range ranges {
		if rnge[0] <= tmp[1] {
			tmp[1] = utils.MaxInt(tmp[1], rnge[1])
		} else {
			res += tmp[1] - tmp[0] + 1
			tmp = rnge
		}
	}
	return res + tmp[1] - tmp[0] + 1
}
