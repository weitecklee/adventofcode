package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
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
	nodeMap := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(nodeMap))

}

type Node struct {
	name    string
	outputs []*Node
}

func parseInput(data []string) map[string]*Node {
	nodeMap := make(map[string]*Node, len(data))
	nodeRegex := regexp.MustCompile(`\w+`)
	for _, line := range data {
		matches := nodeRegex.FindAllString(line, -1)
		name := matches[0]
		if _, ok := nodeMap[name]; !ok {
			node := Node{
				name,
				make([]*Node, 0),
			}
			nodeMap[name] = &node
		}
		node := nodeMap[name]
		node.outputs = make([]*Node, 0, len(matches)-1)
		for _, match := range matches[1:] {
			if _, ok := nodeMap[match]; !ok {
				output := Node{
					match,
					make([]*Node, 0),
				}
				nodeMap[match] = &output
			}
			node.outputs = append(node.outputs, nodeMap[match])
		}
	}
	return nodeMap
}

func part1(nodeMap map[string]*Node) int {
	youNode := nodeMap["you"]
	outNode := nodeMap["out"]
	queue := []*Node{youNode}
	res := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr == outNode {
			res += 1
			continue
		}

		queue = append(queue, curr.outputs...)
	}

	return res
}
