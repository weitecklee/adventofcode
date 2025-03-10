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
	puzzleInput := parseInput(strings.Split(string(data), " "))
	root := createTree(puzzleInput)
	fmt.Println(part1(root))
	fmt.Println(part2(root))
}

type Node struct {
	nChildren int
	nMetadata int
	children  []*Node
	metadata  []int
	parent    *Node
}

func (n *Node) sumOfMetadata() int {
	res := 0
	for _, n := range n.metadata {
		res += n
	}
	return res
}

func (n *Node) sumOfMetadataRecursive() int {
	res := n.sumOfMetadata()
	for _, child := range n.children {
		res += child.sumOfMetadataRecursive()
	}
	return res
}

func (n *Node) value() int {
	if n.nChildren == 0 {
		return n.sumOfMetadata()
	}
	res := 0
	for _, i := range n.metadata {
		if i > 0 && i <= n.nChildren {
			res += n.children[i-1].value()
		}
	}
	return res
}

func NewNode(nChildren, nMetadata int, parent *Node) *Node {
	return &Node{
		nChildren,
		nMetadata,
		make([]*Node, 0, nChildren),
		make([]int, 0, nMetadata),
		parent,
	}
}

func parseInput(data []string) []int {
	puzzleInput := make([]int, len(data))
	for i, s := range data {
		if num, err := strconv.Atoi(s); err != nil {
			panic(err)
		} else {
			puzzleInput[i] = num
		}
	}
	return puzzleInput
}

func createTree(puzzleInput []int) *Node {
	parent := NewNode(puzzleInput[0], puzzleInput[1], nil)
	i := 2
	for i < len(puzzleInput) {
		if len(parent.children) < parent.nChildren {
			child := NewNode(puzzleInput[i], puzzleInput[i+1], parent)
			parent.children = append(parent.children, child)
			parent = child
			i += 2
		} else if len(parent.metadata) != parent.nMetadata {
			parent.metadata = append(parent.metadata, puzzleInput[i:i+parent.nMetadata]...)
			i += parent.nMetadata
		} else {
			parent = parent.parent
		}
	}
	return parent
}

func part1(root *Node) int {
	return root.sumOfMetadataRecursive()
}

func part2(root *Node) int {
	return root.value()
}
