package main

import (
	"container/heap"
	"fmt"
	"math"
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

type Item struct {
	priority int
	steps    int
	pos      [2]int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
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

	var queue PriorityQueue
	heap.Init(&queue)
	heap.Push(&queue, &Item{dist(emptyNode.pos, target), 0, emptyNode.pos, 0})
	visited := map[[2]int]int{}

	for len(queue) > 0 {
		item := heap.Pop(&queue).(*Item)
		if item.pos == target {
			steps = item.steps
			break
		}

		curr, x, y := item.steps, item.pos[0], item.pos[1]
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
			heap.Push(&queue, &Item{dist([2]int{x2, y2}, target) + curr, curr, [2]int{x2, y2}, 0})
		}
	}

	return steps + 1 + (xMax-1)*5
}
