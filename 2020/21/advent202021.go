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
	fmt.Println(part1(input))
}

func part1(input []string) int {
	re := regexp.MustCompile(`\w+`)
	allergens := map[string]map[string]bool{}
	masterIngredients := map[string]int{}
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		ingred := map[string]bool{}
		currentAllergens := []string{}
		foundContains := false
		for _, match := range matches {
			if !foundContains {
				if match == "contains" {
					foundContains = true
				} else {
					ingred[match] = true
					masterIngredients[match]++
				}
			} else {
				currentAllergens = append(currentAllergens, match)
			}
		}
		for _, allergen := range currentAllergens {
			if _, ok := allergens[allergen]; ok {
				newIngred := map[string]bool{}
				for ing := range allergens[allergen] {
					if ingred[ing] {
						newIngred[ing] = true
					}
				}
				allergens[allergen] = newIngred
			} else {
				allergens[allergen] = ingred
			}
		}
	}
	possibleIngredients := map[string]bool{}
	for _, ingredients := range allergens {
		for ingredient := range ingredients {
			possibleIngredients[ingredient] = true
		}
	}
	count := 0
	for ingredient, n := range masterIngredients {
		if !possibleIngredients[ingredient] {
			count += n
		}
	}
	return count
}
