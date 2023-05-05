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
	relationships := parseInput(input)
	fmt.Println(part1(relationships))
}

func parseInput(input []string) map[string]map[string]int {
	relationships := map[string]map[string]int{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		if _, ok := relationships[parts[0]]; !ok {
			relationships[parts[0]] = map[string]int{}
		}
		n, _ := strconv.Atoi(parts[3])
		if parts[2] == "lose" {
			n *= -1
		}
		relationships[parts[0]][parts[10][:len(parts[10])-1]] = n
	}
	return relationships
}

func recurSeating(relationships *map[string]map[string]int, seated *map[string]bool, circle *[]string, maxHappiness *int, happiness int, n *int) {
	if len(*circle) == *n {
		first := (*circle)[0]
		last := (*circle)[len(*circle)-1]
		happiness += (*relationships)[first][last]
		happiness += (*relationships)[last][first]
		if happiness > *maxHappiness {
			*maxHappiness = happiness
		}
		return
	}
	curr := (*circle)[len(*circle)-1]
	for name := range (*relationships)[curr] {
		if !(*seated)[name] {
			(*seated)[name] = true
			*circle = append(*circle, name)
			happiness += (*relationships)[curr][name]
			happiness += (*relationships)[name][curr]
			recurSeating(relationships, seated, circle, maxHappiness, happiness, n)
			(*seated)[name] = false
			*circle = (*circle)[:len(*circle)-1]
			happiness -= (*relationships)[curr][name]
			happiness -= (*relationships)[name][curr]
		}
	}
}

func part1(relationships map[string]map[string]int) int {
	seated := map[string]bool{}
	circle := []string{}
	maxHappiness := 0
	happiness := 0
	n := len(relationships)
	for name := range relationships {
		seated[name] = true
		circle = append(circle, name)
		break
	}
	recurSeating(&relationships, &seated, &circle, &maxHappiness, happiness, &n)
	return maxHappiness
}
