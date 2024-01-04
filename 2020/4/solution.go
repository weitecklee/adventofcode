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

func part1(input []string) int {
	fields := []string{
		"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid",
	}
	r, _ := regexp.Compile(`(\w{3}):(\S+)`)
	row := 0
	valids := 0
	for row < len(input) {
		currFields := make(map[string]bool)
		present := 0
		for row < len(input) && len(input[row]) > 0 {
			matches := r.FindAllStringSubmatch(input[row], -1)
			for _, match := range matches {
				currFields[match[1]] = true
			}
			row++
		}
		for _, field := range fields[:len(fields)-1] {
			if currFields[field] {
				present++
			}
		}
		if present == len(fields)-1 {
			valids++
		}
		row++
	}
	return valids
}

func validate(field string, data string) bool {
	switch field {
	case "byr":
		year, _ := strconv.Atoi(data)
		return year >= 1920 && year <= 2002
	case "iyr":
		year, _ := strconv.Atoi(data)
		return year >= 2010 && year <= 2020
	case "eyr":
		year, _ := strconv.Atoi(data)
		return year >= 2020 && year <= 2030
	case "hgt":
		r, _ := regexp.Compile(`^(\d+)(cm|in)$`)
		match := r.FindStringSubmatch(data)
		if match == nil {
			return false
		}
		hgt, _ := strconv.Atoi(match[1])
		if match[2] == "cm" {
			return hgt >= 150 && hgt <= 193
		}
		if match[2] == "in" {
			return hgt >= 59 && hgt <= 76
		}
	case "hcl":
		matched, _ := regexp.MatchString(`^#[a-f0-9]{6}$`, data)
		return matched
	case "ecl":
		ecl := map[string]bool{
			"amb": true,
			"blu": true,
			"brn": true,
			"gry": true,
			"grn": true,
			"hzl": true,
			"oth": true,
		}
		return ecl[data]
	case "pid":
		matched, _ := regexp.MatchString(`^\d{9}$`, data)
		return matched
	}
	return false
}

func part2(input []string) int {
	r, _ := regexp.Compile(`(\w{3}):(\S+)`)
	row := 0
	valids := 0
	for row < len(input) {
		present := 0
		invalid := false
		for row < len(input) && len(input[row]) > 0 {
			if !invalid {
				matches := r.FindAllStringSubmatch(input[row], -1)
				for _, match := range matches {
					if match[1] == "cid" {
						continue
					}
					if validate(match[1], match[2]) {
						present++
					} else {
						invalid = true
						break
					}
				}
			}
			row++
		}
		if !invalid && present == 7 {
			valids++
		}
		row++
	}
	return valids
}
