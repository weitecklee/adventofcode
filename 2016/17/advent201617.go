package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	fmt.Println(part1(input))
}

type State struct {
	pos   [2]int
	route string
}

type Direction struct {
	diff [2]int
	char string
}

func doorsState(passcode string, route string) []bool {
	hash := md5.Sum([]byte(passcode + route))
	hash2 := hex.EncodeToString(hash[:])
	doors := []bool{}
	for _, r := range hash2[:4] {
		doors = append(doors, testDoor(r))
	}
	return doors
}

func testDoor(r rune) bool {
	return r == 'b' || r == 'c' || r == 'd' || r == 'e' || r == 'f'
}

func part1(passcode string) string {
	start := State{
		pos:   [2]int{0, 0},
		route: "",
	}
	queue := []State{start}
	directions := []Direction{
		{
			diff: [2]int{0, -1},
			char: "U",
		},
		{
			diff: [2]int{0, 1},
			char: "D",
		},
		{
			diff: [2]int{-1, 0},
			char: "L",
		},
		{
			diff: [2]int{1, 0},
			char: "R",
		},
	}
	exit := [2]int{3, 3}
	for i := 0; i < len(queue); i++ {
		state := queue[i]
		doors := doorsState(passcode, state.route)
		for j := 0; j < 4; j++ {
			if !doors[j] {
				continue
			}
			nextPos := [2]int{state.pos[0] + directions[j].diff[0], state.pos[1] + directions[j].diff[1]}
			if nextPos[0] < 0 || nextPos[1] < 0 || nextPos[0] > 3 || nextPos[1] > 3 {
				continue
			}
			nextState := State{
				pos:   nextPos,
				route: state.route + directions[j].char,
			}
			if nextPos == exit {
				return nextState.route
			}
			queue = append(queue, nextState)
		}
	}
	return ""
}
