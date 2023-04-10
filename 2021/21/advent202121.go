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
	fmt.Println(part2(positions))
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

func part2(positions [2]int) int {
	universes := map[[4]int]int{} // universe key is {player1Position, player1Score, player2Position, player2Score}, value is number of universes with that state
	initial := [4]int{positions[0], 0, positions[1], 0}
	outcomes := [7][2]int{{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1}} // possible outcomes of 3 rolls and number of occurrences for each outcome
	player1Wins := 0
	player2Wins := 0
	universes[initial] = 1
	for len(universes) > 0 {
		universes2 := map[[4]int]int{}
		for universe, n := range universes {
			for _, outcome := range outcomes {
				player1Position := (universe[0] + outcome[0]) % 10
				player1Score := universe[1] + player1Position + 1
				if player1Score >= 21 {
					player1Wins += n * outcome[1]
				} else {
					universes2[[4]int{player1Position, player1Score, universe[2], universe[3]}] += n * outcome[1]
				}
			}
		}
		universes = universes2
		universes3 := map[[4]int]int{}
		for universe, n := range universes {
			for _, outcome := range outcomes {
				player2Position := (universe[2] + outcome[0]) % 10
				player2Score := universe[3] + player2Position + 1
				if player2Score >= 21 {
					player2Wins += n * outcome[1]
				} else {
					universes3[[4]int{universe[0], universe[1], player2Position, player2Score}] += n * outcome[1]
				}
			}
		}
		universes = universes3
	}
	if player1Wins > player2Wins {
		return player1Wins
	}
	return player2Wins
}
