package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(string(data)))
	fmt.Println(part2(data))
}

func part1(input string) int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(input, -1)
	res := 0
	for _, match := range matches {
		n, _ := strconv.Atoi(match)
		res += n
	}
	return res
}

func recurMap(thing map[string]any) float64 {
	sum := 0.0
	for _, v := range thing {
		if v == "red" {
			return 0
		} else if m, ok := v.(map[string]any); ok {
			sum += recurMap(m)
		} else if a, ok := v.([]any); ok {
			sum += recurArr(a)
		} else if n, ok := v.(float64); ok {
			sum += n
		}
	}
	return sum
}

func recurArr(thing []any) float64 {
	sum := 0.0
	for _, v := range thing {
		if m, ok := v.(map[string]any); ok {
			sum += recurMap(m)
		} else if a, ok := v.([]any); ok {
			sum += recurArr(a)
		} else if n, ok := v.(float64); ok {
			sum += n
		}
	}
	return sum
}

func part2(data []byte) float64 {
	var store map[string]any
	json.Unmarshal(data, &store)
	return recurMap(store)
}
