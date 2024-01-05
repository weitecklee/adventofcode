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
	parsedInput := parseInput(input)
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
}

func parseInput(input []string) *[][]int {
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
	return &parsed
}

func bfs(queue []int, input *[][]int, visited *map[int]bool) {
	for len(queue) > 0 {
		pipe := queue[0]
		connectingPipes := (*input)[pipe][1:]
		for _, connectingPipe := range connectingPipes {
			if !(*visited)[connectingPipe] {
				queue = append(queue, connectingPipe)
				(*visited)[connectingPipe] = true
			}
		}
		queue = queue[1:]
	}
}

func part1(input *[][]int) int {
	visited := map[int]bool{}
	visited[0] = true
	queue := []int{0}
	bfs(queue, input, &visited)
	return len(visited)
}

func part2(input *[][]int) int {
	visited := map[int]bool{}
	groups := 0
	for _, line := range *input {
		if visited[line[0]] {
			continue
		}
		queue := []int{line[0]}
		visited[line[0]] = true
		bfs(queue, input, &visited)
		groups++
	}
	return groups
}
