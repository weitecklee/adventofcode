package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
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
	fmt.Println(part1(parseInput(strings.Split(string(data), "\n"))))
	fmt.Println(part2(parseInput(strings.Split(string(data), "\n"))))
}

type Step struct {
	name     byte
	nPrereqs int
	prereqs  map[*Step]struct{}
	postreqs map[*Step]struct{}
	timeDone int
}

func NewStep(name byte) *Step {
	return &Step{name, 0, make(map[*Step]struct{}), make(map[*Step]struct{}), math.MaxInt}
}

func (s *Step) addPrereq(s2 *Step) {
	s.nPrereqs++
	s.prereqs[s2] = struct{}{}
}

func (s *Step) addPostreq(s2 *Step) {
	s.postreqs[s2] = struct{}{}
}

func parseInput(data []string) map[byte]*Step {
	stepMap := make(map[byte]*Step, 26)
	stepRegex := regexp.MustCompile(`Step (\w).*?step (\w)`)
	for _, line := range data {
		match := stepRegex.FindStringSubmatch(line)
		step1Name := match[1][0]
		step2Name := match[2][0]
		if _, ok := stepMap[step1Name]; !ok {
			stepMap[step1Name] = NewStep(step1Name)
		}
		if _, ok := stepMap[step2Name]; !ok {
			stepMap[step2Name] = NewStep(step2Name)
		}
		step1 := stepMap[step1Name]
		step2 := stepMap[step2Name]
		step1.addPostreq(step2)
		step2.addPrereq(step1)
	}
	return stepMap
}

func part1(stepMap map[byte]*Step) string {
	var sb strings.Builder
	var stepDone *Step
	var nameDone byte
	steps := make(map[*Step]struct{}, len(stepMap))
	for _, step := range stepMap {
		steps[step] = struct{}{}
	}
	for len(steps) > 0 {
		nameDone = 255
		for step := range steps {
			if step.nPrereqs == 0 && step.name < nameDone {
				nameDone = step.name
				stepDone = step
			}
		}
		delete(steps, stepDone)
		sb.WriteByte(stepDone.name)
		for postStep := range stepDone.postreqs {
			postStep.nPrereqs--
		}
	}
	return sb.String()
}

func part2(stepMap map[byte]*Step) int {
	steps := make(map[*Step]struct{}, len(stepMap))
	for _, step := range stepMap {
		steps[step] = struct{}{}
	}
	t := 0
	workers := make(map[*Step]struct{}, 5)
	for len(steps) > 0 || len(workers) > 0 {
		tDone := math.MaxInt
		var stepsDone []*Step
		// iterate through workers, find the first ones done
		for worker := range workers {
			if worker.timeDone < tDone {
				tDone = worker.timeDone
				stepsDone = []*Step{worker}
			} else if worker.timeDone == t {
				stepsDone = append(stepsDone, worker)
			}
		}
		for _, step := range stepsDone {
			// set as done by deleting them, update their postreqs, update t
			delete(workers, step)
			for postStep := range step.postreqs {
				postStep.nPrereqs--
			}
			t = step.timeDone
		}
		// collect any steps that can be done
		var canBeDone []*Step
		for step := range steps {
			if step.nPrereqs == 0 {
				canBeDone = append(canBeDone, step)
			}
		}
		// sort them by name if there are multiple
		sort.SliceStable(canBeDone, func(i, j int) bool {
			return canBeDone[i].name < canBeDone[j].name
		})
		for len(workers) < 5 && len(canBeDone) > 0 {
			// assign them to workers in order
			step := canBeDone[0]
			canBeDone = canBeDone[1:]
			delete(steps, step)
			workers[step] = struct{}{}
			step.timeDone = t + 61 + int(step.name-'A')
		}
	}
	return t

}
