package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"

	"github.com/weitecklee/adventofcode/utils"
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
	fmt.Println(solve(puzzleInput))
}

type juncBox [3]int

type Pair struct {
	a, b, dist int
}

type UnionFind struct {
	parent []int
	size   []int
}

func (uf *UnionFind) FindParent(n int) int {
	if uf.parent[n] != n {
		uf.parent[n] = uf.FindParent(uf.parent[n])
	}
	return uf.parent[n]
}

func (uf *UnionFind) Union(a, b int) {
	rootA := uf.FindParent(a)
	rootB := uf.FindParent(b)

	if rootA == rootB {
		return
	}

	if uf.size[rootA] < uf.size[rootB] {
		rootA, rootB = rootB, rootA
	}

	uf.parent[rootB] = rootA
	uf.size[rootA] += uf.size[rootB]

}

func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{
		make([]int, n),
		make([]int, n),
	}
	for i := range uf.parent {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return &uf
}

func parseInput(data []string) []juncBox {
	points := make([]juncBox, len(data))
	for i, line := range data {
		points[i] = juncBox(utils.ExtractInts(line))
	}
	return points
}

func calcDistance(point1, point2 juncBox) int {
	res := 0
	for i := range point1 {
		res += utils.PowInt(point1[i]-point2[i], 2)
	}
	return res
}

func solve(puzzleInput []juncBox) (int, int) {
	pairs := make([]Pair, 0, len(puzzleInput)*(len(puzzleInput)-1)/2)
	for i, box1 := range puzzleInput {
		for j, box2 := range puzzleInput[i+1:] {
			dist := calcDistance(box1, box2)
			pair := Pair{i, i + j + 1, dist}
			pairs = append(pairs, pair)
		}
	}

	uf := NewUnionFind(1000)
	slices.SortFunc(pairs, func(a, b Pair) int {
		return a.dist - b.dist
	})

	for i := range 1000 {
		pair := pairs[i]
		uf.Union(pair.a, pair.b)
	}

	sizeMap := make(map[int]int)
	for _, p := range uf.parent {
		sizeMap[p] = uf.size[p]
	}

	sizes := make([]int, 0, len(sizeMap))
	for _, s := range sizeMap {
		sizes = append(sizes, s)
	}
	slices.Sort(sizes)

	part1 := 1
	for i := 1; i <= 3; i++ {
		part1 *= sizes[len(sizes)-i]
	}

	var part2 int
	i := 1000
	for {
		pair := pairs[i]
		uf.Union(pair.a, pair.b)
		if uf.size[uf.FindParent(pair.a)] == 1000 {
			part2 = puzzleInput[pair.a][0] * puzzleInput[pair.b][0]
			break
		}
		i += 1
	}

	return part1, part2
}
