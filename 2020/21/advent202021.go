package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	count, allergens := part1(input)
	fmt.Println(count)
	fmt.Println(part2(allergens))
}

func part1(input []string) (int, map[string]map[string]bool) {
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
	return count, allergens
}

func part2(allergens map[string]map[string]bool) string {
	determined := map[string]string{}
	determined2 := map[string]string{}
	n := len(allergens)
	for len(determined) < n {
		for allergen, ingredients := range allergens {
			for ingredient := range ingredients {
				if _, ok := determined[ingredient]; ok {
					delete(ingredients, ingredient)
				}
			}
			if len(ingredients) == 1 {
				for ingredient := range ingredients {
					determined[ingredient] = allergen
					determined2[allergen] = ingredient
				}
				delete(allergens, allergen)
			}
		}
	}
	toSort := []string{}
	for _, allergen := range determined {
		toSort = append(toSort, allergen)
	}
	sort.Strings(toSort)
	ans := []string{}
	for _, allergen := range toSort {
		ans = append(ans, determined2[allergen])
	}
	return strings.Join(ans, ",")
}
