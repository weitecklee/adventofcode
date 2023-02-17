package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input202015.txt")
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
