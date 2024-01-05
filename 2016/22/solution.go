package main

import (
	"fmt"
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
	fmt.Println(part1(input))
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

func part1(input []string) int {
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
