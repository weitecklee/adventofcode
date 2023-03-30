package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	components := parseInput(input)
	fmt.Println(part1(components))
}

type Component struct {
	ports    [2]int
	strength int
}

func parseInput(input []string) map[int][]*Component {
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
	return components
}

func bridge(currentPort int, components map[int][]*Component, usedComponents map[*Component]bool, maxStrength *int, strength int) {
	for _, component := range components[currentPort] {
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

func part1(components map[int][]*Component) int {
	maxStrength := 0
	usedComponents := map[*Component]bool{}
	bridge(0, components, usedComponents, &maxStrength, 0)
	return maxStrength
}
