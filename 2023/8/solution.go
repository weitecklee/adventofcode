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
	instructions, nodeMap := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(instructions, nodeMap))
	fmt.Println(part2(instructions, nodeMap))
}

type Node struct {
	name  string
	left  *Node
	right *Node
}

func parseInput(data []string) (string, map[string]*Node) {
	nodeMap := make(map[string]*Node, len(data)-2)
	nodeRegex := regexp.MustCompile(`\w{3}`)
	for _, line := range data[2:] {
		matches := nodeRegex.FindAllString(line, -1)
		if _, ok := nodeMap[matches[0]]; !ok {
			nodeMap[matches[0]] = &Node{matches[0], nil, nil}
		}
		if _, ok := nodeMap[matches[1]]; !ok {
			nodeMap[matches[1]] = &Node{matches[1], nil, nil}
		}
		if _, ok := nodeMap[matches[2]]; !ok {
			nodeMap[matches[2]] = &Node{matches[2], nil, nil}
		}
		nodeMap[matches[0]].left = nodeMap[matches[1]]
		nodeMap[matches[0]].right = nodeMap[matches[2]]
	}
	return data[0], nodeMap
}

func part1(instructions string, nodeMap map[string]*Node) int {
	curr := nodeMap["AAA"]
	i := 0
	for {
		switch instructions[i%len(instructions)] {
		case 'L':
			curr = curr.left
		case 'R':
			curr = curr.right
		}
		i++
		if curr.name == "ZZZ" {
			return i
		}
	}
}

func stepsToNodeZAndLoop(instructions string, startingNode *Node) (int, int) {
	curr := startingNode
	i := 0
	loop1 := 0
	for {
		switch instructions[i%len(instructions)] {
		case 'L':
			curr = curr.left
		case 'R':
			curr = curr.right
		}
		i++
		if loop1 == 0 && curr.name[2] == 'Z' {
			loop1 = i
		} else if curr.name[2] == 'Z' {
			return loop1, i - loop1
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func part2(instructions string, nodeMap map[string]*Node) int {
	// Upon inspection, turns out that steps to a 'Z' node is equal to steps
	// to another 'Z' node (i.e., it'll loop once it gets to a 'Z' node)
	// Just need to calculate LCM of the loop sizes (This is only Day 8 after all)

	var aNodes []*Node
	for name, node := range nodeMap {
		if name[2] == 'A' {
			aNodes = append(aNodes, node)
		}
	}
	loops := make([]int, len(aNodes))
	var loop1, loop2 int
	for i, node := range aNodes {
		loop1, loop2 = stepsToNodeZAndLoop(instructions, node)
		if loop1 != loop2 {
			// Just making sure the loops are equal
			panic("Unequal loops!")
		}
		loops[i] = loop1
	}
	res := 1
	for _, n := range loops {
		res = lcm(res, n)
	}
	return res
}
