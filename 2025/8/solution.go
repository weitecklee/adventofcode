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
type circuit map[*juncBox]struct{}

func parseInput(data []string) []*juncBox {
	points := make([]*juncBox, len(data))
	for i, line := range data {
		box := juncBox(utils.ExtractInts(line))
		points[i] = &box
	}
	return points
}

func calcDistance(point1, point2 *juncBox) int {
	res := 0
	for i := range point1 {
		res += utils.PowInt(point1[i]-point2[i], 2)
	}
	return res
}

func connectBoxes(box1, box2 *juncBox, circuits *[]*circuit) {
	var circ1, circ2 *circuit
	circs := *circuits

	for _, circ := range circs {
		if _, ok := (*circ)[box1]; ok {
			circ1 = circ
		}
		if _, ok := (*circ)[box2]; ok {
			circ2 = circ
		}
		if circ1 != nil && circ2 != nil {
			break
		}
	}

	if circ1 == nil && circ2 == nil {
		circ := make(circuit)
		circ[box1] = struct{}{}
		circ[box2] = struct{}{}
		circs = append(circs, &circ)
	} else if circ1 != nil && circ2 == nil {
		(*circ1)[box2] = struct{}{}
	} else if circ1 == nil && circ2 != nil {
		(*circ2)[box1] = struct{}{}
	} else if circ1 == circ2 {
	} else {
		for box := range *circ2 {
			(*circ1)[box] = struct{}{}
		}
		for i, circ := range circs {
			if circ == circ2 {
				circs = append(circs[:i], circs[i+1:]...)
				break
			}
		}
	}

	*circuits = circs
}

func solve(puzzleInput []*juncBox) (int, int) {
	distances := make(map[*[2]*juncBox]int)
	pairs := make([]*[2]*juncBox, 0, len(puzzleInput)*(len(puzzleInput)-1)/2)
	for i, box1 := range puzzleInput {
		for _, box2 := range puzzleInput[i+1:] {
			dist := calcDistance(box1, box2)
			pair := [2]*juncBox{box1, box2}
			distances[&pair] = dist
			pairs = append(pairs, &pair)
		}
	}

	slices.SortFunc(pairs, func(a, b *[2]*juncBox) int {
		return distances[a] - distances[b]
	})

	circuits := make([]*circuit, 0, 1000)
	for range 1000 {
		pair := pairs[0]
		pairs = pairs[1:]
		connectBoxes(pair[0], pair[1], &circuits)
	}

	sizes := make([]int, len(circuits))
	for i, circ := range circuits {
		sizes[i] = len(*circ)
	}
	slices.Sort(sizes)
	part1 := 1
	for i := 1; i <= 3; i++ {
		part1 *= sizes[len(sizes)-i]
	}

	var part2 int
	for {
		pair := pairs[0]
		pairs = pairs[1:]
		connectBoxes(pair[0], pair[1], &circuits)
		if len(*circuits[0]) == 1000 {
			part2 = pair[0][0] * pair[1][0]
			break
		}
	}

	return part1, part2
}
