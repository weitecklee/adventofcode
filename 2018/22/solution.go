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

var (
	depth         int
	target        [2]int
	erosionLevel  = map[[2]int]int{}
	geologicIndex = map[[2]int]int{}
	risk          = map[[2]int]int{}
	directions    = [4][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	toolOptions = [3]map[string]struct{}{
		{"torch": {}, "climbing gear": {}},
		{"climbing gear": {}, "neither": {}},
		{"torch": {}, "neither": {}},
	}
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
	depth, target = parseInput(input)
	geologicIndex[[2]int{0, 0}] = 0
	geologicIndex[target] = 0
	fmt.Println(part1())
	fmt.Println(part2())
}

func parseInput(input []string) (int, [2]int) {
	re := regexp.MustCompile(`\d+`)
	parts0 := re.FindAllString(input[0], -1)
	depth, _ := strconv.Atoi(parts0[0])
	parts1 := re.FindAllString(input[1], -1)
	target0, _ := strconv.Atoi(parts1[0])
	target1, _ := strconv.Atoi(parts1[1])
	return depth, [2]int{target0, target1}
}

func calcErosionLevel(x, y int) int {
	if val, exists := erosionLevel[[2]int{x, y}]; exists {
		return val
	}
	geologicIndex := calcGeologicIndex(x, y)
	erosionLevel[[2]int{x, y}] = (geologicIndex + depth) % 20183
	return erosionLevel[[2]int{x, y}]
}

func calcGeologicIndex(x, y int) int {
	if val, exists := geologicIndex[[2]int{x, y}]; exists {
		return val
	}
	var gi int
	if x == 0 {
		gi = y * 48271
	} else if y == 0 {
		gi = x * 16807
	} else {
		gi = calcErosionLevel(x-1, y) * calcErosionLevel(x, y-1)
	}
	geologicIndex[[2]int{x, y}] = gi
	return gi
}

func calcRisk(x, y int) int {
	if val, exists := risk[[2]int{x, y}]; exists {
		return val
	}
	erosionLevel := calcErosionLevel(x, y)
	risk[[2]int{x, y}] = erosionLevel % 3
	return risk[[2]int{x, y}]
}

func calcDist(x, y int) int {
	return int(math.Abs(float64(x-target[0])) + math.Abs(float64(y-target[1])))
}

func part1() int {
	sum := 0
	for x := 0; x <= target[0]; x++ {
		for y := 0; y <= target[1]; y++ {
			sum += calcRisk(x, y)
		}
	}
	return sum
}

type Item struct {
	priority int
	time     int
	x, y     int
	tool     string
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Visited struct {
	x, y int
	tool string
}

func part2() int {
	tMin := math.MaxInt
	queue := &MinHeap{Item{
		priority: 0,
		time:     0,
		x:        0,
		y:        0,
		tool:     "torch",
	}}
	heap.Init(queue)
	visited := map[Visited]int{}

	for len(*queue) > 0 {
		item := heap.Pop(queue).(Item)
		if item.time >= tMin {
			continue
		}
		if time, exists := visited[Visited{item.x, item.y, item.tool}]; exists && time <= item.time {
			continue
		}
		visited[Visited{item.x, item.y, item.tool}] = item.time

		if item.x == target[0] && item.y == target[1] {
			if item.tool == "torch" {
				if item.time < tMin {
					tMin = item.time
				}
			} else {
				heap.Push(queue, Item{
					priority: item.time + 7,
					time:     item.time + 7,
					x:        item.x,
					y:        item.y,
					tool:     "torch",
				})
			}
			continue
		}

		risk := calcRisk(item.x, item.y)
		dist := calcDist(item.x, item.y)
		for _, dir := range directions {
			x2, y2 := item.x+dir[0], item.y+dir[1]
			if x2 < 0 || y2 < 0 {
				continue
			}
			risk2 := calcRisk(x2, y2)
			dist2 := calcDist(x2, y2)
			for tool := range toolOptions[risk] {
				if _, exists := toolOptions[risk2][tool]; !exists {
					continue
				}
				if tool == item.tool {
					heap.Push(queue, Item{
						priority: item.time + 1 + dist2,
						time:     item.time + 1,
						x:        x2,
						y:        y2,
						tool:     tool,
					})
				} else {
					heap.Push(queue, Item{
						priority: item.time + 7 + dist,
						time:     item.time + 7,
						x:        item.x,
						y:        item.y,
						tool:     tool,
					})
				}
			}
		}
	}

	return tMin
}
