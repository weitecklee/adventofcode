package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
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
	replacements, molecule := parseInput(input)
	fmt.Println(part1(replacements, molecule))
	fmt.Println(part2(molecule))
}

func parseInput(input []string) (*map[string][]string, string) {
	replacements := map[string][]string{}
	for i := 0; i < len(input)-2; i++ {
		parts := strings.Split(input[i], " => ")
		if _, ok := replacements[parts[0]]; !ok {
			replacements[parts[0]] = []string{}
		}
		replacements[parts[0]] = append(replacements[parts[0]], parts[1])
	}
	return &replacements, input[len(input)-1]
}

func replaceMolecules(replacements *map[string][]string, molecule string) *[]string {
	res := map[string]bool{}
	for orig, reps := range *replacements {
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
	molecules := make([]string, 0, len(res))
	for molec := range res {
		molecules = append(molecules, molec)
	}
	return &molecules
}

func part1(replacements *map[string][]string, molecule string) int {
	return len(*replaceMolecules(replacements, molecule))
}

/*
	First, tried BFS using refactored replaceMolecules function.
	Tree blows up way too big to be feasible.
	Then, tried the reverse by going from the target back to "e".
	Again, tree blows up way too big to be feasible.
	Reading others' solutions on reddit, they used greedy algo that always
	uses the longest replacement available. This is not a general solution.
	Indeed, it doesn't work for my input.

	Finally, used the solution provided by u/askalski, comment copied in solution.md for posterity.
	https://www.reddit.com/r/adventofcode/comments/3xflz8/day_19_solutions/cy4etju/
*/

func part2(molecule string) int {
	re := regexp.MustCompile(`[A-Z][a-z]?`)
	countMolecules := len(re.FindAllString(molecule, -1))
	countRn := strings.Count(molecule, "Rn")
	countAr := strings.Count(molecule, "Ar")
	countY := strings.Count(molecule, "Y")
	return countMolecules - countRn - countAr - 2*countY - 1
}
