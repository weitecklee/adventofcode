package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	components := parseInput(input)
	fmt.Println(part1(components))
	fmt.Println(part2(components))
}

type Component struct {
	ports    [2]int
	strength int
}

func parseInput(input []string) *map[int][]*Component {
	components := map[int][]*Component{}
	for _, line := range input {
		nums := strings.Split(line, "/")
		port1, _ := strconv.Atoi(nums[0])
		port2, _ := strconv.Atoi(nums[1])
		comp := Component{
			ports:    [2]int{port1, port2},
			strength: port1 + port2,
		}
		components[port1] = append(components[port1], &comp)
		components[port2] = append(components[port2], &comp)
	}
	return &components
}

func bridge(currentPort int, components *map[int][]*Component, usedComponents map[*Component]bool, maxStrength *int, strength int) {
	for _, component := range (*components)[currentPort] {
		if !usedComponents[component] {
			usedComponents[component] = true
			port := component.ports[0]
			if port == currentPort {
				port = component.ports[1]
			}
			bridge(port, components, usedComponents, maxStrength, strength+component.strength)
			usedComponents[component] = false
		}
	}
	if strength > *maxStrength {
		*maxStrength = strength
	}
}

func part1(components *map[int][]*Component) int {
	maxStrength := 0
	usedComponents := map[*Component]bool{}
	bridge(0, components, usedComponents, &maxStrength, 0)
	return maxStrength
}

func longBridge(currentPort int, components *map[int][]*Component, usedComponents map[*Component]bool, maxStrength *int, strength int, maxLength *int, length int) {
	for _, component := range (*components)[currentPort] {
		if !usedComponents[component] {
			usedComponents[component] = true
			port := component.ports[0]
			if port == currentPort {
				port = component.ports[1]
			}
			longBridge(port, components, usedComponents, maxStrength, strength+component.strength, maxLength, length+1)
			usedComponents[component] = false
		}
	}
	if length > *maxLength {
		*maxLength = length
		*maxStrength = strength
	} else if length == *maxLength && strength > *maxStrength {
		*maxStrength = strength
	}
}

func part2(components *map[int][]*Component) int {
	maxStrength := 0
	maxLength := 0
	usedComponents := map[*Component]bool{}
	longBridge(0, components, usedComponents, &maxStrength, 0, &maxLength, 0)
	return maxStrength
}
