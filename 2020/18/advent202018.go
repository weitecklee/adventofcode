package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func evaluate(expression string) string {
	re1 := regexp.MustCompile(`^\([^()]*\)$`)
	if re1.MatchString(expression) {
		expression = expression[1 : len(expression)-1]
	}
	re2 := regexp.MustCompile(`\([^(]*?\)`)
	for re2.MatchString(expression) {
		expression = re2.ReplaceAllStringFunc(expression, evaluate)
	}
	parts := strings.Split(expression, " ")
	curr, _ := strconv.Atoi(parts[0])
	for i := 1; i < len(parts); i += 2 {
		tmp, _ := strconv.Atoi(parts[i+1])
		switch parts[i] {
		case "+":
			curr += tmp
		case "*":
			curr *= tmp
		}
	}
	return strconv.Itoa(curr)
}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		ev, _ := strconv.Atoi(evaluate(line))
		sum += ev
	}
	return sum
}

func add(expression string) string {
	r := regexp.MustCompile(`\d+`)
	nums := r.FindAllString(expression, -1)
	res := 0
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		res += n
	}
	return strconv.Itoa(res)
}

func evaluate2(expression string) string {
	re1 := regexp.MustCompile(`^\([^()]*\)$`)
	if re1.MatchString(expression) {
		expression = expression[1 : len(expression)-1]
	}
	re2 := regexp.MustCompile(`\([^(]*?\)`)
	for re2.MatchString(expression) {
		expression = re2.ReplaceAllStringFunc(expression, evaluate2)
	}
	re3 := regexp.MustCompile(`\d+ \+ \d+`)
	for re3.MatchString(expression) {
		expression = re3.ReplaceAllStringFunc(expression, add)
	}
	parts := strings.Split(expression, " ")
	curr := 1
	for i := 0; i < len(parts); i += 2 {
		tmp, _ := strconv.Atoi(parts[i])
		curr *= tmp
	}
	return strconv.Itoa(curr)
}

func part2(input []string) int {
	sum := 0
	for _, line := range input {
		ev, _ := strconv.Atoi(evaluate2(line))
		sum += ev
	}
	return sum
}
