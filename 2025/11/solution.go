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
	fmt.Println(solve(nodeMap))
}

type Graph struct {
	nodeMap  map[string]*Node
	pathMemo map[[2]*Node]int
}

type Node struct {
	outputs []*Node
}

func NewGraph(nodeMap map[string]*Node) *Graph {
	nNodes := len(nodeMap)
	g := Graph{
		nodeMap,
		make(map[[2]*Node]int, nNodes*(nNodes-1)),
	}
	g.initialize()
	return &g
}

func (g *Graph) initialize() {
	for _, node := range g.nodeMap {
		for _, output := range node.outputs {
			g.pathMemo[[2]*Node{node, output}] = 1
		}
	}
}

func (g *Graph) NumPaths(fromStr, toStr string) int {
	fromNode := g.nodeMap[fromStr]
	toNode := g.nodeMap[toStr]
	return g.dfs(fromNode, toNode)
}

func (g *Graph) dfs(fromNode *Node, toNode *Node) int {
	pair := [2]*Node{fromNode, toNode}
	if v, ok := g.pathMemo[pair]; ok {
		return v
	}

	res := 0
	for _, node := range fromNode.outputs {
		res += g.dfs(node, toNode)
	}

	g.pathMemo[pair] = res
	return res
}

func parseInput(data []string) map[string]*Node {
	nodeMap := make(map[string]*Node, len(data))
	nodeRegex := regexp.MustCompile(`\w+`)
	for _, line := range data {
		matches := nodeRegex.FindAllString(line, -1)
		name := matches[0]
		if _, ok := nodeMap[name]; !ok {
			node := Node{
				[]*Node{},
			}
			nodeMap[name] = &node
		}
		node := nodeMap[name]
		node.outputs = make([]*Node, 0, len(matches)-1)
		for _, match := range matches[1:] {
			if _, ok := nodeMap[match]; !ok {
				output := Node{
					[]*Node{},
				}
				nodeMap[match] = &output
			}
			node.outputs = append(node.outputs, nodeMap[match])
		}
	}
	return nodeMap
}

func solve(nodeMap map[string]*Node) (int, int) {
	graph := NewGraph(nodeMap)

	svr2dac := graph.NumPaths("svr", "dac")
	dac2fft := graph.NumPaths("dac", "fft")
	fft2out := graph.NumPaths("fft", "out")
	svr2fft := graph.NumPaths("svr", "fft")
	fft2dac := graph.NumPaths("fft", "dac")
	dac2out := graph.NumPaths("dac", "out")

	part1 := graph.NumPaths("you", "out")
	part2 := svr2dac*dac2fft*fft2out + svr2fft*fft2dac*dac2out

	return part1, part2
}
