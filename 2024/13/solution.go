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
	clawMachines := parseInput(strings.Split(string(data), "\n\n"))
	fmt.Println(part1(clawMachines))
	fmt.Println(part2(clawMachines))
}

type ClawMachine struct {
	buttonA [2]int
	buttonB [2]int
	prize   [2]int
}

func NewClawMachine(config string) *ClawMachine {
	lines := strings.Split(config, "\n")
	return &ClawMachine{parseLine(lines[0]), parseLine(lines[1]), parseLine(lines[2])}
}

func (cm *ClawMachine) PressesToWin(deviation int) (int, int) {
	a0, b0, c0 := cm.buttonA[0], cm.buttonB[0], cm.prize[0]+deviation
	a1, b1, c1 := cm.buttonA[1], cm.buttonB[1], cm.prize[1]+deviation
	det := a0*b1 - b0*a1
	if det == 0 {
		return 0, 0
	}
	xNumer := c0*b1 - c1*b0
	yNumer := a0*c1 - a1*c0
	if xNumer%det != 0 || yNumer%det != 0 {
		return 0, 0
	}
	return xNumer / det, yNumer / det
}

func (cm *ClawMachine) TokensToWin(deviation int) int {
	a, b := cm.PressesToWin(deviation)
	return 3*a + b
}

var numRegex = regexp.MustCompile(`\d+`)

func parseLine(line string) [2]int {
	nums := numRegex.FindAllString(line, 2)
	n0, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}
	n1, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}
	return [2]int{n0, n1}
}

func parseInput(data []string) []*ClawMachine {
	clawMachines := make([]*ClawMachine, len(data))
	for i, s := range data {
		clawMachines[i] = NewClawMachine(s)
	}
	return clawMachines
}

func part1(clawMachines []*ClawMachine) int {
	res := 0
	for _, cm := range clawMachines {
		res += cm.TokensToWin(0)
	}
	return res
}

func part2(clawMachines []*ClawMachine) int {
	res := 0
	for _, cm := range clawMachines {
		res += cm.TokensToWin(10000000000000)
	}
	return res
}
