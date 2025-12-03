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
	fmt.Println(part1_2(puzzleInput))
	fmt.Println(part2_2(puzzleInput))
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

func findLargestInWindow(window []int) (int, int) {
	max := -1
	idx := -1
	for i, n := range window {
		if n > max {
			max = n
			idx = i
			if max == 9 {
				break
			}
		}
	}
	return max, idx
}

func findLargestJoltage2(digits []int, windowLen int) int {
	res := 0
	idx := 0
	for i := range windowLen {
		n, idx2 := findLargestInWindow(digits[idx : len(digits)-windowLen+i+1])
		idx += idx2 + 1
		res = res*10 + n
	}
	return res
}

func part1_2(puzzleInput [][]int) int {
	res := 0
	for _, row := range puzzleInput {
		res += findLargestJoltage2(row, 2)
	}
	return res
}

func part2_2(puzzleInput [][]int) int {
	res := 0
	for _, row := range puzzleInput {
		res += findLargestJoltage2(row, 12)
	}
	return res
}

/*
	Implementation #1: Linked list struct with number of nodes equal to batteries.
	The method starts from the end of the row. When a node receives a value, it
	first checks if the next node	(if there is one) is empty, and passes the value
	on if true. Otherwise, it	checks if the value is equal to higher than the one
	it has. If so, it takes that value and passes its old value down(if possible).
	Implementation #2: Sliding window.
	Window is constructed starting from next possible index of row, ending at end
	of row but leaving enough elements for remaining batteries.

	Implementation #1: avg 192µs
	Implementation #2: avg 84µs
*/
