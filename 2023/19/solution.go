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
	workflowMap, parts := parseInput(strings.Split(string(data), "\n\n"))
	fmt.Println(part1(workflowMap, parts))
}

var (
	workflowRegex = regexp.MustCompile(`^(\w+){(.*)}$`)
	ruleRegex     = regexp.MustCompile(`^([xmas])([<>])(\d+):(\w+)$`)
	partRegex     = regexp.MustCompile(`^{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}$`)
)

type Rule struct {
	category byte
	comp     byte
	n        int
	dst      string
}

func NewRule(str string) *Rule {
	match := ruleRegex.FindStringSubmatch(str)
	if match == nil {
		panic(fmt.Sprintf("Error matching rule regex with: %s", str))
	}
	n, err := strconv.Atoi(match[3])
	if err != nil {
		panic(err)
	}
	return &Rule{match[1][0], match[2][0], n, match[4]}
}

func (r *Rule) AssessPart(p Part) bool {
	if r.comp == '<' {
		return p[r.category] < r.n
	}
	return p[r.category] > r.n
}

type Workflow struct {
	name    string
	rules   []*Rule
	lastDst string
}

func NewWorkflow(str string) *Workflow {
	match := workflowRegex.FindStringSubmatch(str)
	if match == nil {
		panic(fmt.Sprintf("Error matching workflow regex with: %s", str))
	}
	name := match[1]
	ruleParts := strings.Split(match[2], ",")
	rules := make([]*Rule, len(ruleParts)-1)
	for j, rulePart := range ruleParts[:len(ruleParts)-1] {
		rules[j] = NewRule(rulePart)
	}
	return &Workflow{name, rules, ruleParts[len(ruleParts)-1]}
}

func (w *Workflow) AssessPart(p Part) string {
	for _, rule := range w.rules {
		if rule.AssessPart(p) {
			return rule.dst
		}
	}
	return w.lastDst
}

type Part map[byte]int

func NewPart(str string) Part {
	match := partRegex.FindStringSubmatch(str)
	if match == nil {
		panic(fmt.Sprintf("Error matching part regex with: %s", str))
	}
	x, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(match[2])
	if err != nil {
		panic(err)
	}
	a, err := strconv.Atoi(match[3])
	if err != nil {
		panic(err)
	}
	s, err := strconv.Atoi(match[4])
	if err != nil {
		panic(err)
	}
	return Part{'x': x, 'm': m, 'a': a, 's': s}
}

func (p Part) Rating() int {
	return p['x'] + p['m'] + p['a'] + p['s']
}

func (p Part) Assess(workflowMap map[string]*Workflow) bool {
	curr := "in"
	for curr != "A" && curr != "R" {
		curr = workflowMap[curr].AssessPart(p)
	}
	return curr == "A"
}

func parseInput(data []string) (map[string]*Workflow, []Part) {
	workflowLines := strings.Split(data[0], "\n")
	workflowMap := make(map[string]*Workflow, len(workflowLines))
	partLines := strings.Split(data[1], "\n")
	parts := make([]Part, len(partLines))

	for _, line := range workflowLines {
		workflow := NewWorkflow(line)
		workflowMap[workflow.name] = workflow
	}

	for i, line := range partLines {
		parts[i] = NewPart(line)
	}

	return workflowMap, parts
}

func part1(workflowMap map[string]*Workflow, parts []Part) int {
	res := 0
	for _, part := range parts {
		if part.Assess(workflowMap) {
			res += part.Rating()
		}
	}
	return res
}
