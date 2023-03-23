package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
	fmt.Println(part2(parseInput(input)))
}

func parseInput(input []string) [][]int {
	parsed := [][]int{}
	re := regexp.MustCompile(`\d+`)
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		tmp := []int{}
		for _, match := range matches {
			n, _ := strconv.Atoi(match)
			tmp = append(tmp, n)
		}
		parsed = append(parsed, tmp)
	}
	return parsed
}

func part1(input [][]int) int {
	visited := map[int]bool{}
	visited[0] = true
	queue := []int{0}
	for len(queue) > 0 {
		pipe := queue[0]
		connectingPipes := input[pipe][1:]
		for _, connectingPipe := range connectingPipes {
			if !visited[connectingPipe] {
				queue = append(queue, connectingPipe)
				visited[connectingPipe] = true
			}
		}
		queue = queue[1:]
	}
	return len(visited)
}

func part2(input [][]int) int {
	visited := map[int]bool{}
	groups := 0
	for _, line := range input {
		if visited[line[0]] {
			continue
		}
		queue := []int{line[0]}
		visited[line[0]] = true
		for len(queue) > 0 {
			pipe := queue[0]
			connectingPipes := input[pipe][1:]
			for _, connectingPipe := range connectingPipes {
				if !visited[connectingPipe] {
					queue = append(queue, connectingPipe)
					visited[connectingPipe] = true
				}
			}
			queue = queue[1:]
		}
		groups++
	}
	return groups
}