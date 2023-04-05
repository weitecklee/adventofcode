package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	area := parseInput(input)
	fmt.Println(part1(area))
	fmt.Println(part2(area))
}

func parseInput(input []string) [][]string {
	parsed := [][]string{}
	for _, line := range input {
		parsed = append(parsed, strings.Split(line, ""))
	}
	return parsed
}

func simulate(area [][]string) [][]string {
	area2 := [][]string{}
	for a := 0; a < len(area); a++ {
		row := []string{}
		for b := 0; b < len(area[0]); b++ {
			count := map[string]int{}
			for m := -1; m <= 1; m++ {
				for n := -1; n <= 1; n++ {
					if (m == 0 && n == 0) || a+m < 0 || b+n < 0 || a+m >= len(area) || b+n >= len(area[0]) {
						continue
					}
					count[area[a+m][b+n]]++
				}
			}
			spot := area[a][b]
			switch area[a][b] {
			case ".":
				if count["|"] >= 3 {
					spot = "|"
				}
			case "|":
				if count["#"] >= 3 {
					spot = "#"
				}
			case "#":
				if count["#"] < 1 || count["|"] < 1 {
					spot = "."
				}
			}
			row = append(row, spot)
		}
		area2 = append(area2, row)
	}
	return area2
}

func calculateResourceValue(area [][]string) int {
	count := map[string]int{}
	for _, row := range area {
		for _, c := range row {
			count[c]++
		}
	}
	return count["#"] * count["|"]
}

func part1(area [][]string) int {
	for i := 0; i < 10; i++ {
		area = simulate(area)
	}
	return calculateResourceValue(area)
}

func stringify(area [][]string) string {
	areaString := ""
	for _, line := range area {
		areaString += strings.Join(line, "")
	}
	return areaString
}

func part2(area [][]string) int {
	history := map[string]int{}
	i := 1
	areaString := stringify(area)
	history[areaString] = 1
	for {
		i++
		area = simulate(area)
		areaString = stringify(area)
		if history[areaString] > 0 {
			break
		}
		history[areaString] = i
	}
	period := i - history[areaString]
	minutesBeforePeriod := history[areaString] - 1
	minutesLeftToSimulate := (1000000000 - minutesBeforePeriod) % period
	for j := 0; j < minutesLeftToSimulate; j++ {
		area = simulate(area)
	}
	return calculateResourceValue(area)
}
