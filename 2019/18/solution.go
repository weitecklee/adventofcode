package main

import (
	"container/heap"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"

	"github.com/weitecklee/adventofcode/utils"
)

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	nodeMap, allKeys := parseInput(string(data))
	fmt.Println(part1(nodeMap, allKeys))
	nodeMap2, allKeys2, regionalKeys := parseInput2(string(data))
	fmt.Println(part2(nodeMap2, allKeys2, regionalKeys))
}

type Node struct {
	sym       rune
	pos       [2]int
	neighbors map[*Node]int
}

func (n *Node) IsStart() bool {
	return !unicode.IsLetter(n.sym)
}

func (n *Node) IsKey() bool {
	return unicode.IsLetter(n.sym) && n.sym == unicode.ToLower(n.sym)
}

func (n *Node) IsGate() bool {
	return unicode.IsLetter(n.sym) && n.sym != unicode.ToLower(n.sym)
}

type QueueEntry struct {
	pos   [2]int
	steps int
}

type Value struct {
	steps    int
	currNode *Node
	keys     uint32
}

type Value2 struct {
	steps     int
	currNodes []*Node
	keys      uint32
}

func (v *Value2) State() [5]rune {
	var state [5]rune
	for i, node := range v.currNodes {
		state[i] = node.sym
	}
	state[4] = rune(v.keys)
	return state
}

func parseInput(s string) (map[rune]*Node, uint32) {
	maze := strings.Split(s, "\n")
	var allKeys uint32
	nodeMap := make(map[rune]*Node)
	for r, row := range maze {
		for c, ch := range row {
			if ch != '#' && ch != '.' {
				nodeMap[ch] = &Node{ch, [2]int{r, c}, make(map[*Node]int)}
				if unicode.IsLetter(ch) && ch == unicode.ToLower(ch) {
					allKeys = addKey(allKeys, ch)
				}
			}
		}
	}
	for _, node := range nodeMap {
		findNeighbors(maze, node, nodeMap)
	}
	return nodeMap, allKeys
}

func findNeighbors(maze []string, node *Node, nodeMap map[rune]*Node) {
	queue := []QueueEntry{{node.pos, 0}}
	visited := make(map[[2]int]struct{})
	visited[node.pos] = struct{}{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		steps := curr.steps + 1
		for _, d := range directions {
			pos := [2]int{curr.pos[0] + d[0], curr.pos[1] + d[1]}
			if pos[0] < 0 || pos[1] < 0 || pos[0] >= len(maze) || pos[1] >= len(maze[0]) {
				continue
			}
			if _, ok := visited[pos]; ok {
				continue
			}
			visited[pos] = struct{}{}
			ch := rune(maze[pos[0]][pos[1]])
			if ch == '#' {
				continue
			}
			if ch == '.' {
				queue = append(queue, QueueEntry{pos, steps})
			} else {
				neighbor := nodeMap[ch]
				node.neighbors[neighbor] = steps
			}
		}
	}
}

func addKey(mask uint32, key rune) uint32 {
	return mask | (1 << (key - 'a'))
}

func hasKey(mask uint32, key rune) bool {
	return mask&(1<<(key-'a')) != 0
}

func missingKeys(regionalKeys, keys uint32) bool {
	return regionalKeys&^keys != 0
}

func part1(nodeMap map[rune]*Node, allKeys uint32) int {
	queue := utils.NewMinHeap[Value]()
	start := nodeMap['@']
	heap.Push(queue, &utils.Item[Value]{
		Priority: 0,
		Value: Value{
			steps:    0,
			currNode: start,
			keys:     0,
		}})
	visited := make(map[[2]rune]int)
	visited[[2]rune{'@', 0}] = 0

	for len(queue.PriorityQueue) > 0 {
		item := heap.Pop(queue).(*utils.Item[Value])
		if item.Value.keys == allKeys {
			return item.Value.steps
		}
		for neighbor, d := range item.Value.currNode.neighbors {
			steps := item.Value.steps + d
			keys := item.Value.keys
			state := [2]rune{neighbor.sym, rune(keys)}
			ch := neighbor.sym

			// 3 possibilities:
			// 1. picking up a new key => update keys and state
			// 2. at a gate without the key => discontinue
			// 3. all others => continue
			if neighbor.IsKey() && !hasKey(item.Value.keys, ch) {
				// picking up  new key
				keys = addKey(item.Value.keys, ch)
				state[1] = rune(keys)
			} else if neighbor.IsGate() && !hasKey(item.Value.keys, unicode.ToLower(ch)) {
				// at gate without key
				continue
			}

			// all others (passing through gate, already picked up key, back at starting location)
			if n, ok := visited[state]; ok && n <= steps {
				continue
			}
			visited[state] = steps
			heap.Push(queue, &utils.Item[Value]{
				Priority: steps,
				Value: Value{
					steps:    steps,
					currNode: neighbor,
					keys:     keys,
				},
			})
		}
	}
	return -1
}

func parseInput2(s string) (map[rune]*Node, uint32, [4]uint32) {
	maze := strings.Split(s, "\n")
	var startR, startC int
	for r, row := range maze {
		for c, ch := range row {
			if ch == '@' {
				startR = r
				startC = c
				break
			}
		}
		if startR != 0 {
			break
		}
	}
	tmp := []byte(maze[startR-1])
	tmp[startC-1] = '1' // first new entrance
	tmp[startC] = '#'
	tmp[startC+1] = '2' // second new entrance
	maze[startR-1] = string(tmp)
	tmp = []byte(maze[startR])
	tmp[startC-1] = '#'
	tmp[startC] = '#'
	tmp[startC+1] = '#'
	maze[startR] = string(tmp)
	tmp = []byte(maze[startR+1])
	tmp[startC-1] = '3' // third new entrance
	tmp[startC] = '#'
	tmp[startC+1] = '4' // fourth new entrance
	maze[startR+1] = string(tmp)

	var allKeys uint32
	nodeMap := make(map[rune]*Node)
	for r, row := range maze {
		for c, ch := range row {
			if ch != '#' && ch != '.' {
				nodeMap[ch] = &Node{ch, [2]int{r, c}, make(map[*Node]int)}
				if unicode.IsLetter(ch) && ch == unicode.ToLower(ch) {
					allKeys = addKey(allKeys, ch)
				}
			}
		}
	}
	for _, node := range nodeMap {
		findNeighbors(maze, node, nodeMap)
	}
	regionalKeys := findRegionalKeys(nodeMap)
	return nodeMap, allKeys, regionalKeys
}

func findRegionalKeys(nodeMap map[rune]*Node) [4]uint32 {
	var regionalKeys [4]uint32
	for i := range 4 {
		start := nodeMap[rune('1'+i)]
		visited := make(map[rune]struct{})
		queue := []*Node{start}
		visited[start.sym] = struct{}{}
		var keys uint32
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			for neighbor := range curr.neighbors {
				if _, ok := visited[neighbor.sym]; ok {
					continue
				}
				visited[neighbor.sym] = struct{}{}
				if neighbor.IsKey() {
					keys = addKey(keys, neighbor.sym)
				}
				queue = append(queue, neighbor)
			}
		}
		regionalKeys[i] = keys
	}
	return regionalKeys
}

func part2(nodeMap map[rune]*Node, allKeys uint32, regionalKeys [4]uint32) int {
	queue := utils.NewMinHeap[Value2]()
	initial := Value2{
		steps:     0,
		currNodes: []*Node{nodeMap['1'], nodeMap['2'], nodeMap['3'], nodeMap['4']},
		keys:      0,
	}
	heap.Push(queue, &utils.Item[Value2]{
		Priority: 0,
		Value:    initial,
	})
	visited := make(map[[5]rune]int)
	visited[initial.State()] = 0

	for len(queue.PriorityQueue) > 0 {
		item := heap.Pop(queue).(*utils.Item[Value2])
		if item.Value.keys == allKeys {
			return item.Value.steps
		}
		for i, node := range item.Value.currNodes {
			// check if region has keys to pick up, otherwise skip
			if !missingKeys(regionalKeys[i], item.Value.keys) {
				continue
			}

			for neighbor, d := range node.neighbors {
				tmp := make([]*Node, len(item.Value.currNodes))
				copy(tmp, item.Value.currNodes)
				tmp[i] = neighbor
				steps := item.Value.steps + d
				keys := item.Value.keys
				ch := neighbor.sym

				if neighbor.IsKey() && !hasKey(item.Value.keys, ch) {
					keys = addKey(item.Value.keys, ch)
				} else if neighbor.IsGate() && !hasKey(item.Value.keys, unicode.ToLower(ch)) {
					continue
				}
				curr := Value2{
					steps:     steps,
					currNodes: tmp,
					keys:      keys,
				}
				state := curr.State()

				if n, ok := visited[state]; ok && n <= steps {
					continue
				}
				visited[state] = steps
				heap.Push(queue, &utils.Item[Value2]{
					Priority: steps,
					Value:    curr,
				})
			}
		}
	}
	return -1
}
