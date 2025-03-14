package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
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
	fmt.Println(part2(workflowMap))
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
		return &Rule{'x', '>', 0, str}
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

func (r *Rule) Reverse() *Rule {
	var comp byte = '<'
	if comp == r.comp {
		comp = '>'
	}
	n := r.n
	if r.comp == '>' {
		n++
	} else {
		n--
	}
	return &Rule{r.category, comp, n, r.dst}
}

type Workflow struct {
	name  string
	rules []*Rule
}

func NewWorkflow(str string) *Workflow {
	match := workflowRegex.FindStringSubmatch(str)
	if match == nil {
		panic(fmt.Sprintf("Error matching workflow regex with: %s", str))
	}
	name := match[1]
	ruleParts := strings.Split(match[2], ",")
	rules := make([]*Rule, len(ruleParts))
	for j, rulePart := range ruleParts {
		rules[j] = NewRule(rulePart)
	}
	return &Workflow{name, rules}
}

func (w *Workflow) AssessPart(p Part) string {
	for _, rule := range w.rules {
		if rule.AssessPart(p) {
			return rule.dst
		}
	}
	return ""
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

type QueueEntry struct {
	rules    []*Rule
	workflow *Workflow
}

func part2(workflowMap map[string]*Workflow) int {
	queue := []QueueEntry{{[]*Rule{}, workflowMap["in"]}}
	var combinations [][]*Rule
	for len(queue) > 0 {
		rules := queue[0].rules
		workflow := queue[0].workflow
		queue = queue[1:]
		for i := range workflow.rules {
			currentRules := slices.Clone(rules)
			for j := range i {
				currentRules = append(currentRules, workflow.rules[j].Reverse())
			}
			currentRules = append(currentRules, workflow.rules[i])

			if workflow.rules[i].dst == "A" {
				combinations = append(combinations, currentRules)
			} else if workflow.rules[i].dst != "R" {
				queue = append(queue, QueueEntry{currentRules, workflowMap[workflow.rules[i].dst]})
			}
		}
	}

	res := 0
	for _, combo := range combinations {
		approved := map[byte][]int{'x': {1, 4000}, 'm': {1, 4000}, 'a': {1, 4000}, 's': {1, 4000}}
		for _, rule := range combo {
			if rule.comp == '>' {
				if approved[rule.category][0] < rule.n+1 {
					approved[rule.category][0] = rule.n + 1
				}
			} else {
				if approved[rule.category][1] > rule.n-1 {
					approved[rule.category][1] = rule.n - 1
				}
			}
		}
		count := 1
		for _, bounds := range approved {
			count *= bounds[1] - bounds[0] + 1
		}
		res += count
	}
	return res
}
