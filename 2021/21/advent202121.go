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
	positions := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(positions))
}

func parseInput(input []string) [2]int {
	positions := [2]int{}
	re := regexp.MustCompile(`\d+`)
	for i, line := range input {
		nums := re.FindAllString(line, -1)
		num, _ := strconv.Atoi(nums[1])
		positions[i] = num - 1
	}
	return positions
}

func rollDie(start int) (int, int) {
	switch start {
	case 98:
		return 1, 297
	case 99:
		return 2, 200
	case 100:
		return 3, 103
	default:
		return start + 3, 3*start + 3
	}
}

func part1(positions [2]int) int {
	player1Score := 0
	player2Score := 0
	rolls := 0
	die := 1
	move := 0
	for {
		die, move = rollDie(die)
		rolls += 3
		positions[0] = (positions[0] + move) % 10
		player1Score += positions[0] + 1
		if player1Score >= 1000 {
			return player2Score * rolls
		}
		die, move = rollDie(die)
		rolls += 3
		positions[1] = (positions[1] + move) % 10
		player2Score += positions[1] + 1
		if player2Score >= 1000 {
			return player1Score * rolls
		}
	}
}
