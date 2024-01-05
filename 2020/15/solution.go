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
	input := strings.Split(string(data), ",")
	fmt.Println(recite(input, 2020))
	fmt.Println(recite(input, 30000000))
}

func recite(input []string, turn int) int {
	numbers := map[int]int{}
	i := 0
	for _, s := range input {
		i++
		n, _ := strconv.Atoi(s)
		numbers[n] = i
	}
	next := 0
	for i < turn-1 {
		i++
		if numbers[next] > 0 {
			tmp := i - numbers[next]
			numbers[next] = i
			next = tmp
		} else {
			numbers[next] = i
			next = 0
		}
	}
	return next
}
