package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.ReadFile("input202019.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
}

func part1(input []string) int {
	rules := map[string][][]string{}
	rules2 := map[string][]string{}
	re := regexp.MustCompile(`\d+`)
	re2 := regexp.MustCompile(`[a-z]`)
	row := 0
	for re.MatchString(input[row]) {
		line := strings.Replace(input[row], ":", "", 1)
		nums := strings.Split(line, " ")
		letter := re2.FindString(line)
		if letter == "" {
			tmpRule := []string{}
			for _, rule := range nums[1:] {
				if rule == "|" {
					rules[nums[0]] = append(rules[nums[0]], tmpRule)
					tmpRule = []string{}
				} else {
					tmpRule = append(tmpRule, rule)
				}
			}
			rules[nums[0]] = append(rules[nums[0]], tmpRule)

		} else {
			rules2[nums[0]] = []string{letter}
		}
		row++
	}
	for len(rules) > 0 {
		for n, rule := range rules {
			skip := false
			for _, r := range rule {
				for _, r2 := range r {
					_, ok := rules2[r2]
					if !ok {
						skip = true
						break
					}
				}
				if skip {
					break
				}
			}
			if skip {
				continue
			}
			for _, r := range rule {
				if len(r) == 1 {
					rules2[n] = append(rules2[n], rules2[r[0]]...)
				} else {
					newRules := rules2[r[0]]
					for _, r2 := range r[1:] {
						newRules = combine(newRules, rules2[r2])
					}
					rules2[n] = append(rules2[n], newRules...)
				}
			}
			delete(rules, n)
		}
	}
	rule0 := map[string]bool{}
	for _, rule := range rules2["0"] {
		rule0[rule] = true
	}
	count := 0
	row++
	for row < len(input) {
		if rule0[input[row]] {
			count++
		}
		row++
	}
	return count
}

func combine(rule1 []string, rule2 []string) []string {
	newRule := []string{}
	for _, r1 := range rule1 {
		for _, r2 := range rule2 {
			newRule = append(newRule, r1+r2)
		}
	}
	return newRule
}
