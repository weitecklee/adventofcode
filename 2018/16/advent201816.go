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
	fmt.Println(solve(parseInput(input)))
}

func parseInput(input []string) ([][3][4]int, [][4]int) {
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
	for len(input[i]) == 0 {
		i++
	}
	testProgram := [][4]int{}
	for i < len(input) {
		nums := re.FindAllString(input[i], -1)
		row := [4]int{}
		for k, num := range nums {
			n, _ := strconv.Atoi(num)
			row[k] = n
		}
		testProgram = append(testProgram, row)
		i++
	}
	return samples, testProgram
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

func solve(samples [][3][4]int, testProgram [][4]int) (int, int) {
	count := 0
	opcode := map[int]map[int]bool{}
	for i := 0; i < 16; i++ {
		opcode[i] = map[int]bool{}
	}
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
				opcode[instruc[0]][i] = true
			}
		}
		if matches >= 3 {
			count++
		}
	}
	opcodemap := map[int]int{}
	for len(opcodemap) < 16 {
		determined := []int{}
		for i, poss := range opcode {
			if len(poss) == 1 {
				for j := range poss {
					opcodemap[i] = j
					determined = append(determined, j)
				}
				delete(opcode, i)
			}
		}
		for _, n := range determined {
			for i := range opcode {
				delete(opcode[i], n)
			}
		}
	}
	before := [4]int{0, 0, 0, 0}
	for _, instruc := range testProgram {
		before = execute(opcodemap[instruc[0]], instruc, before)
	}
	return count, before[0]
}
