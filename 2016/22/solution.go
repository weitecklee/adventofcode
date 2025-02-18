package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
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
	nodes := parseInput(input)
	fmt.Println(part1(nodes))
	fmt.Println(part2(nodes))
}

type Node struct {
	pos   [2]int
	size  int
	used  int
	avail int
}

func (n *Node) Viability(n2 *Node) bool {
	return n.used > 0 && n.used <= n2.avail
}

func parseInput(input []string) []Node {
	nodes := []Node{}
	re := regexp.MustCompile(`\d+`)
	for i := 2; i < len(input); i++ {
		numsStr := re.FindAllString(input[i], -1)
		nums := []int{}
		for _, s := range numsStr {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
		nodes = append(nodes, Node{
			pos:   [2]int{nums[0], nums[1]},
			size:  nums[2],
			used:  nums[3],
			avail: nums[4],
		})
	}
	return nodes
}

func part1(nodes []Node) int {
	pairs := 0
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[i].Viability(&nodes[j]) {
				pairs++
			}
			if nodes[j].Viability(&nodes[i]) {
				pairs++
			}
		}
	}
	return pairs
}

func dist(a, b [2]int) int {
	return int(math.Abs(float64(a[0]-b[0]))) + int(math.Abs(float64(a[1]-b[1])))
}

func part2(nodes []Node) int {

	nodeMap := map[[2]int]*Node{}
	var xMax, yMax int
	var emptyNode *Node
	for i := 0; i < len(nodes); i++ {
		nodeMap[nodes[i].pos] = &nodes[i]
		if nodes[i].pos[0] > xMax {
			xMax = nodes[i].pos[0]
		}
		if nodes[i].pos[1] > yMax {
			yMax = nodes[i].pos[1]
		}
		if nodes[i].used == 0 {
			emptyNode = &nodes[i]
		}
	}

	target := [2]int{xMax - 1, 0}
	steps := 0

	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	queue := [][4]int{{0, 0, emptyNode.pos[0], emptyNode.pos[1]}}
	visited := map[[2]int]int{}

	for len(queue) > 0 {
		curr, x, y := queue[0][1], queue[0][2], queue[0][3]
		queue = queue[1:]
		if x == target[0] && y == target[1] {
			steps = curr
			break
		}
		curr++

		for _, d := range directions {
			dx, dy := d[0], d[1]
			x2, y2 := x+dx, y+dy
			if x2 < 0 || x2 > xMax || y2 < 0 || y2 > yMax {
				continue
			}
			if nd, ok := nodeMap[[2]int{x2, y2}]; ok && nd.used > 100 {
				continue
			}
			if v, ok := visited[[2]int{x2, y2}]; ok && v <= curr {
				continue
			}
			visited[[2]int{x2, y2}] = curr
			queue = append(queue, [4]int{curr + dist(target, [2]int{x2, y2}), curr, x2, y2})
		}

		sort.Slice(queue, func(i, j int) bool {
			return queue[i][0] < queue[j][0]
		})
	}

	return steps + 1 + (xMax-1)*5
}
