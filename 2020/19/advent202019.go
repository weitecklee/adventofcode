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
	rules, rules2, messages := parseInput(input)
	fmt.Println(part1(rules, rules2, messages))
	fmt.Println(part2(rules, rules2, messages))
}

func part1(rules map[string][][]string, rules2 map[string][]string, messages []string) int {
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
	for _, message := range messages {
		if rule0[message] {
			count++
		}
	}
	return count
}

func parseInput(input []string) (map[string][][]string, map[string][]string, []string) {
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
	return rules, rules2, input[row+1:]
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

func part2(rules map[string][][]string, rules2 map[string][]string, messages []string) int {
	for {
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
		_, ok1 := rules2["42"]
		_, ok2 := rules2["31"]
		if ok1 && ok2 {
			break
		}
	}
	rule42 := map[string]bool{}
	rule31 := map[string]bool{}
	for _, rule := range rules2["42"] {
		rule42[rule] = true
	}
	for _, rule := range rules2["31"] {
		rule31[rule] = true
	}
	count := 0
	segLen := len(rules2["42"][0])
	for _, message := range messages {
		pass42 := 0
		pass31 := 0
		stop42 := false
		fail31 := false
		for i := 0; i < len(message); i += segLen {
			seg := message[i : i+segLen]
			if !stop42 {
				if rule42[seg] {
					pass42++
				} else {
					stop42 = true
				}
			}
			if stop42 {
				if rule31[seg] {
					pass31++
				} else {
					fail31 = true
					break
				}
			}
		}
		if !fail31 && pass42 >= 2 && pass31 >= 1 && pass42 > pass31 {
			count++
		}
	}
	return count
}
