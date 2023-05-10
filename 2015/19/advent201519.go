package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	replacements, molecule := parseInput(input)
	fmt.Println(part1(replacements, molecule))
}

func parseInput(input []string) (map[string][]string, string) {
	replacements := map[string][]string{}
	for i := 0; i < len(input)-2; i++ {
		parts := strings.Split(input[i], " => ")
		if _, ok := replacements[parts[0]]; !ok {
			replacements[parts[0]] = []string{}
		}
		replacements[parts[0]] = append(replacements[parts[0]], parts[1])
	}
	return replacements, input[len(input)-1]
}

func part1(replacements map[string][]string, molecule string) int {
	res := map[string]bool{}
	for orig, reps := range replacements {
		re := regexp.MustCompile(orig)
		indices := re.FindAllStringIndex(molecule, -1)
		if indices != nil {
			for _, rep := range reps {
				for _, index := range indices {
					replacedMolecule := molecule[:index[0]] + rep + molecule[index[1]:]
					res[replacedMolecule] = true
				}
			}
		}
	}
	return len(res)
}
