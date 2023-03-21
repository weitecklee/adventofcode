package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
}

type Prog struct {
	weight   int
	children []*Prog
	parent   *Prog
}

func parseInput(input []string) map[string]*Prog {
	re := regexp.MustCompile(`\w+`)
	progs := map[string]*Prog{}
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		name := matches[0]
		weight, _ := strconv.Atoi(matches[1])
		if _, ok := progs[name]; !ok {
			parent := Prog{}
			progs[name] = &parent
		}
		parent := progs[name]
		parent.weight = weight
		for _, match := range matches[2:] {
			if _, ok := progs[match]; !ok {
				child := Prog{}
				progs[match] = &child
			}
			child := progs[match]
			child.parent = parent
			parent.children = append(parent.children, child)
		}
	}
	return progs
}

func part1(progs map[string]*Prog) string {
	for name, prog := range progs {
		if prog.parent == nil {
			return name
		}
	}
	return ""
}
