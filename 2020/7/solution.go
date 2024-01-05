package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	re := regexp.MustCompile(`\w+ \w+ bag`)
	bags := map[string][]string{}
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		for _, bag := range matches[1:] {
			bags[bag] = append(bags[bag], matches[0])
		}
	}
	q := []string{}
	q = append(q, "shiny gold bag")
	checked := map[string]bool{}
	checked["shiny gold bag"] = true
	i := 0
	for i < len(q) {
		for _, bag := range bags[q[i]] {
			if !checked[bag] {
				q = append(q, bag)
				checked[bag] = true
			}
		}
		i++
	}
	return len(checked) - 1
}

func recur(bag string, contains map[string]int, bags map[string][]map[string]any) map[string]int {
	total := 0
	for _, bagMap := range bags[bag] {
		inColor := bagMap["color"].(string)
		inNum := bagMap["num"].(int)
		if _, ok := contains[inColor]; !ok {
			contains = recur(inColor, contains, bags)
		}
		total += (inNum * contains[inColor])
	}
	contains[bag] = total + 1
	return contains
}

func part2(input []string) int {
	re := regexp.MustCompile(`\w+ \w+ bag`)
	re2 := regexp.MustCompile(`\d+`)
	bags := map[string][]map[string]any{}
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		nums := re2.FindAllString(line, -1)
		if matches[1] == "no other bag" {
			bag := map[string]any{}
			bag["color"] = "no other bag"
			bag["num"] = 0
			bags[matches[0]] = append(bags[matches[0]], bag)
		} else {
			for i := range matches[1:] {
				bag := map[string]any{}
				bag["color"] = matches[i+1]
				bag["num"], _ = strconv.Atoi(nums[i])
				bags[matches[0]] = append(bags[matches[0]], bag)
			}
		}
	}
	contains := map[string]int{}
	contains["no other bag"] = 0
	contains = recur("shiny gold bag", contains, bags)
	// contains includes the bag itself with the count so you have to reduce by 1
	return contains["shiny gold bag"] - 1
}
