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
	fmt.Println(part1(parseInput(input)))
}

func parseInput(input []string) [][]string {
	parsed := [][]string{}
	for _, line := range input {
		parsed = append(parsed, strings.Split(line, ""))
	}
	return parsed
}

func part1(area [][]string) int {
	for i := 0; i < 10; i++ {
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
		area = area2
	}
	count := map[string]int{}
	for _, row := range area {
		for _, c := range row {
			count[c]++
		}
	}
	return count["#"] * count["|"]
}
