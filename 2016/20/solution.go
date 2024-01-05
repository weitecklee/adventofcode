package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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
	input := strings.Split(string(data), "\n")
	fmt.Println(solve(input))
}

func solve(input []string) (int, int) {
	ranges := [][2]int{}
	for _, line := range input {
		nums := strings.Split(line, "-")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		ranges = append(ranges, [2]int{n1, n2})
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	merged := [][2]int{}
	tmp := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if tmp[1]+1 >= ranges[i][0] {
			if ranges[i][1] > tmp[1] {
				tmp[1] = ranges[i][1]
			}
		} else {
			merged = append(merged, tmp)
			tmp = ranges[i]
		}
	}
	merged = append(merged, tmp)
	n := 0
	for i := 0; i < len(merged)-1; i++ {
		n += merged[i+1][0] - merged[i][1] - 1
	}
	n += 4294967295 - merged[len(merged)-1][1]
	return merged[0][1] + 1, n
}
