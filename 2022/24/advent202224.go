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

type Blizzard struct {
	pos0 [2]int
	dir  [2]int
}

func (b *Blizzard) locationAtTime(t int, ht int, wd int) [2]int {
	pos := [2]int{}
	if b.dir[0] >= 0 {
		pos[0] = (b.pos0[0]+b.dir[0]*t-1)%(wd-1) + 1
	} else {
		pos[0] = b.pos0[0] + (b.dir[0]*t)%(wd-1)
		if pos[0] < 1 {
			pos[0] = wd + pos[0] - 1
		}
	}
	if b.dir[1] >= 0 {
		pos[1] = (b.pos0[1]+b.dir[1]*t-1)%(ht-1) + 1
	} else {
		pos[1] = b.pos0[1] + (b.dir[1]*t)%(ht-1)
		if pos[1] < 1 {
			pos[1] = ht + pos[1] - 1
		}
	}
	return pos
}

func parseInput(input []string) (map[int][]Blizzard, map[int][]Blizzard, int, int) {
	horizontals := map[int][]Blizzard{}
	verticals := map[int][]Blizzard{}
	for j, row := range input {
		for i, c := range row {
			if c == '.' || c == '#' {
				continue
			}
			tmp := Blizzard{
				pos0: [2]int{i, j},
			}
			switch c {
			case '<':
				tmp.dir = [2]int{-1, 0}
				horizontals[j] = append(horizontals[j], tmp)
			case '>':
				tmp.dir = [2]int{1, 0}
				horizontals[j] = append(horizontals[j], tmp)
			case '^':
				tmp.dir = [2]int{0, -1}
				verticals[i] = append(verticals[i], tmp)
			case 'v':
				tmp.dir = [2]int{0, 1}
				verticals[i] = append(verticals[i], tmp)
			}
		}
	}
	return horizontals, verticals, len(input) - 1, len(input[0]) - 1
}

func part1(horizontals map[int][]Blizzard, verticals map[int][]Blizzard, ht int, wd int) int {
	queue := [][3]int{{1, 0, 0}}
	checks := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{0, 0},
	}
	checked := map[[3]int]bool{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		pos := [2]int{curr[0], curr[1]}
		time := curr[2] + 1
	loop:
		for _, check := range checks {
			checkPos := [2]int{pos[0] + check[0], pos[1] + check[1]}
			toCheck := [3]int{checkPos[0], checkPos[1], time}
			if checked[toCheck] {
				continue
			}
			checked[toCheck] = true
			if checkPos[0] < 1 || checkPos[0] >= wd || checkPos[1] < 0 {
				continue
			}
			if checkPos[1] == 0 && checkPos[0] != 1 {
				continue
			}
			if checkPos[1] == ht {
				if checkPos[0] == wd-1 {
					return time
				}
				continue
			}
			for _, blizzard := range horizontals[checkPos[1]] {
				blizzardPos := blizzard.locationAtTime(time, ht, wd)
				if blizzardPos[0] == checkPos[0] && blizzardPos[1] == checkPos[1] {
					continue loop
				}
			}
			for _, blizzard := range verticals[checkPos[0]] {
				blizzardPos := blizzard.locationAtTime(time, ht, wd)
				if blizzardPos[0] == checkPos[0] && blizzardPos[1] == checkPos[1] {
					continue loop
				}
			}
			queue = append(queue, [3]int{checkPos[0], checkPos[1], time})
		}
	}
	return -1
}
