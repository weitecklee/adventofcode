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
	fmt.Println(part1(parseInput(input)))
}

func parseInput(input []string) [][3][4]int {
	i := 0
	samples := [][3][4]int{}
	re := regexp.MustCompile(`\d+`)
	for {
		if len(input[i]) == 0 || input[i][0] != 'B' {
			break
		}
		sample := [3][4]int{}
		for j := 0; j < 3; j++ {
			nums := re.FindAllString(input[i+j], -1)
			row := [4]int{}
			for k, num := range nums {
				n, _ := strconv.Atoi(num)
				row[k] = n
			}
			sample[j] = row
		}
		samples = append(samples, sample)
		i += 4
	}
	return samples
}

func execute(opcode int, instruc [4]int, before [4]int) [4]int {
	switch opcode {
	case 0:
		before[instruc[3]] = before[instruc[1]] + before[instruc[2]]
	case 1:
		before[instruc[3]] = before[instruc[1]] + instruc[2]
	case 2:
		before[instruc[3]] = before[instruc[1]] * before[instruc[2]]
	case 3:
		before[instruc[3]] = before[instruc[1]] * instruc[2]
	case 4:
		before[instruc[3]] = before[instruc[1]] & before[instruc[2]]
	case 5:
		before[instruc[3]] = before[instruc[1]] & instruc[2]
	case 6:
		before[instruc[3]] = before[instruc[1]] | before[instruc[2]]
	case 7:
		before[instruc[3]] = before[instruc[1]] | instruc[2]
	case 8:
		before[instruc[3]] = before[instruc[1]]
	case 9:
		before[instruc[3]] = instruc[1]
	case 10:
		if instruc[1] > before[instruc[2]] {
			before[instruc[3]] = 1
		} else {
			before[instruc[3]] = 0
		}
	case 11:
		if before[instruc[1]] > instruc[2] {
			before[instruc[3]] = 1
		} else {
			before[instruc[3]] = 0
		}
	case 12:
		if before[instruc[1]] > before[instruc[2]] {
			before[instruc[3]] = 1
		} else {
			before[instruc[3]] = 0
		}
	case 13:
		if instruc[1] == before[instruc[2]] {
			before[instruc[3]] = 1
		} else {
			before[instruc[3]] = 0
		}
	case 14:
		if before[instruc[1]] == instruc[2] {
			before[instruc[3]] = 1
		} else {
			before[instruc[3]] = 0
		}
	case 15:
		if before[instruc[1]] == before[instruc[2]] {
			before[instruc[3]] = 1
		} else {
			before[instruc[3]] = 0
		}
	}
	return before
}

func part1(samples [][3][4]int) int {
	count := 0
	for _, sample := range samples {
		before := sample[0]
		instruc := sample[1]
		after := sample[2]
		matches := 0
		for i := 0; i < 16; i++ {
			result := execute(i, instruc, before)
			same := true
			for j := 0; j < 4; j++ {
				if result[j] != after[j] {
					same = false
					break
				}
			}
			if same {
				matches++
			}
		}
		if matches >= 3 {
			count++
		}
	}
	return count
}
