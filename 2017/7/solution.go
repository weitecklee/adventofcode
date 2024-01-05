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
	progMap := parseInput(input)
	allParent := part1(progMap)
	fmt.Println(allParent)
	fmt.Println(part2(allParent, progMap))
}

type Prog struct {
	name        string
	weight      int
	children    []*Prog
	parent      *Prog
	totalWeight int
}

func parseInput(input []string) map[string]*Prog {
	re := regexp.MustCompile(`\w+`)
	progMap := map[string]*Prog{}
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		name := matches[0]
		weight, _ := strconv.Atoi(matches[1])
		if _, ok := progMap[name]; !ok {
			parent := Prog{}
			parent.name = name
			progMap[name] = &parent
		}
		parent := progMap[name]
		parent.weight = weight
		for _, match := range matches[2:] {
			if _, ok := progMap[match]; !ok {
				child := Prog{}
				child.name = match
				progMap[match] = &child
			}
			child := progMap[match]
			child.parent = parent
			parent.children = append(parent.children, child)
		}
	}
	return progMap
}

func part1(progMap map[string]*Prog) string {
	for name, prog := range progMap {
		if prog.parent == nil {
			return name
		}
	}
	return ""
}

func weigh(prog *Prog) {
	totalWeight := prog.weight
	for _, child := range prog.children {
		if child.totalWeight == 0 {
			weigh(child)
		}
		totalWeight += child.totalWeight
	}
	prog.totalWeight = totalWeight
}

func part2(allParent string, progMap map[string]*Prog) int {
	weigh(progMap[allParent])
	unbalancedProg := progMap[allParent]
	/*
		Assume allParent is unbalanced. This should be necessarily true
		but it works for example and input. Also assume that unbalanced stacks
		are three or more so it's easy to find odd one out.
		Recursively find the odd one out until it's balanced, keep record of the
		discrepancy that needs to be compensated for. If the odd one out is balanced,
		that is the weight that needs to be changed.
	*/
	evenWeight := 0
	oddWeight := 0
	for {
		unbalanced := false
		weight := unbalancedProg.children[0].totalWeight
		for _, child := range unbalancedProg.children[1:] {
			if child.totalWeight != weight {
				unbalanced = true
			}
		}
		if unbalanced {
			if unbalancedProg.children[0].totalWeight == unbalancedProg.children[1].totalWeight {
				evenWeight = unbalancedProg.children[0].totalWeight
				for _, child := range unbalancedProg.children[2:] {
					if child.totalWeight != evenWeight {
						oddWeight = child.totalWeight
						unbalancedProg = child
						break
					}
				}
			} else {
				if unbalancedProg.children[0].totalWeight == unbalancedProg.children[2].totalWeight {
					evenWeight = unbalancedProg.children[0].totalWeight
					oddWeight = unbalancedProg.children[1].totalWeight
					unbalancedProg = unbalancedProg.children[1]
				} else {
					evenWeight = unbalancedProg.children[1].totalWeight
					oddWeight = unbalancedProg.children[0].totalWeight
					unbalancedProg = unbalancedProg.children[0]
				}
			}
		} else {
			return unbalancedProg.weight + evenWeight - oddWeight
		}
	}
}
