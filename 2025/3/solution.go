package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func parseInput(data []string) [][]int {
	res := make([][]int, len(data))
	for i, line := range data {
		row := make([]int, len(line))
		for j, c := range line {
			row[j] = int(c - '0')
		}
		res[i] = row
	}
	return res
}

type LinkedList struct {
	node *Node
}

func newLinkedList(len int) *LinkedList {
	ll := LinkedList{
		node: &Node{0, nil},
	}
	node := ll.node
	for range len - 1 {
		node.next = &Node{0, nil}
		node = node.next
	}
	return &ll
}

func (ll *LinkedList) addValue(v int) {
	ll.node.addValue(v)
}

func (ll *LinkedList) getValue() int {
	res := 0
	node := ll.node
	for node != nil {
		res = res*10 + node.value
		node = node.next
	}
	return res
}

type Node struct {
	value int
	next  *Node
}

func (nd *Node) addValue(v int) {
	if nd.next != nil && nd.next.value == 0 {
		nd.next.addValue(v)
	} else if nd.value <= v {
		if nd.next != nil {
			nd.next.addValue(nd.value)
		}
		nd.value = v
	}
}

func part1(puzzleInput [][]int) int {
	res := 0
	for _, row := range puzzleInput {
		j := findLargestJoltage(row, 2)
		res += j
	}
	return res
}

func part2(puzzleInput [][]int) int {
	res := 0
	for _, row := range puzzleInput {
		j := findLargestJoltage(row, 12)
		res += j
	}
	return res
}

func findLargestJoltage(digits []int, nDigits int) int {
	ll := newLinkedList(nDigits)
	for i := len(digits) - 1; i >= 0; i-- {
		ll.addValue(digits[i])
	}
	return ll.getValue()
}
