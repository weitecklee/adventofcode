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
	fishMap := parseInput(strings.Split(string(data), ","))
	fmt.Println(simulate(fishMap, 80))
	fmt.Println(simulate(fishMap, 256))
}

func parseInput(data []string) map[int]int {
	fishMap := make(map[int]int, 9)
	for _, s := range data {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		fishMap[n]++
	}
	return fishMap
}

func simulate(fishMap map[int]int, days int) int {
	for range days {
		fishMap2 := make(map[int]int, 9)
		for i := range 8 {
			fishMap2[i] = fishMap[i+1]
		}
		fishMap2[8] = fishMap[0]
		fishMap2[6] += fishMap[0]
		fishMap = fishMap2
	}
	res := 0
	for _, n := range fishMap {
		res += n
	}
	return res
}
