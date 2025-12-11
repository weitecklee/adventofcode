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
	fmt.Println(part2(nodeMap))

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
				[]*Node{},
			}
			nodeMap[name] = &node
		}
		node := nodeMap[name]
		node.outputs = make([]*Node, 0, len(matches)-1)
		for _, match := range matches[1:] {
			if _, ok := nodeMap[match]; !ok {
				output := Node{
					match,
					[]*Node{},
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

func makeMemo(nodeMap map[string]*Node) map[[2]*Node]int {
	memo := make(map[[2]*Node]int)
	for _, node := range nodeMap {
		for _, output := range node.outputs {
			memo[[2]*Node{node, output}] = 1
		}
	}
	return memo
}

func dfs(fromNode *Node, toNode *Node, pathMemo map[[2]*Node]int) int {
	pair := [2]*Node{fromNode, toNode}
	if v, ok := pathMemo[pair]; ok {
		return v
	}

	res := 0
	for _, node := range fromNode.outputs {
		res += dfs(node, toNode, pathMemo)
	}

	pathMemo[pair] = res
	return res
}

func part2(nodeMap map[string]*Node) int {
	pathMemo := makeMemo(nodeMap)

	svrNode := nodeMap["svr"]
	dacNode := nodeMap["dac"]
	fftNode := nodeMap["fft"]
	outNode := nodeMap["out"]

	svr2dac := dfs(svrNode, dacNode, pathMemo)
	dac2fft := dfs(dacNode, fftNode, pathMemo)
	fft2out := dfs(fftNode, outNode, pathMemo)
	svr2fft := dfs(svrNode, fftNode, pathMemo)
	fft2dac := dfs(fftNode, dacNode, pathMemo)
	dac2out := dfs(dacNode, outNode, pathMemo)

	return svr2dac*dac2fft*fft2out + svr2fft*fft2dac*dac2out
}
