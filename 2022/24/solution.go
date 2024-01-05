package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(solve(parseInput(input)))
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

func parseInput(input []string) (*map[int]*[]*Blizzard, *map[int]*[]*Blizzard, int, int) {
	horizontals := map[int]*[]*Blizzard{}
	verticals := map[int]*[]*Blizzard{}
	for j, row := range input {
		for i, c := range row {
			if c == '.' || c == '#' {
				continue
			}
			tmp := Blizzard{
				pos0: [2]int{i, j},
			}
			if _, ok := horizontals[j]; !ok {
				tmpArr := []*Blizzard{}
				horizontals[j] = &tmpArr
			}
			if _, ok := verticals[i]; !ok {
				tmpArr := []*Blizzard{}
				verticals[i] = &tmpArr
			}
			switch c {
			case '<':
				tmp.dir = [2]int{-1, 0}
				*horizontals[j] = append(*horizontals[j], &tmp)
			case '>':
				tmp.dir = [2]int{1, 0}
				*horizontals[j] = append(*horizontals[j], &tmp)
			case '^':
				tmp.dir = [2]int{0, -1}
				*verticals[i] = append(*verticals[i], &tmp)
			case 'v':
				tmp.dir = [2]int{0, 1}
				*verticals[i] = append(*verticals[i], &tmp)
			}
		}
	}
	return &horizontals, &verticals, len(input) - 1, len(input[0]) - 1
}

func journey(startPos [2]int, endPos [2]int, startTime int, horizontals *map[int]*[]*Blizzard, verticals *map[int]*[]*Blizzard, ht int, wd int) int {
	queue := [][3]int{{startPos[0], startPos[1], startTime}}
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
			if checkPos[0] < 1 || checkPos[0] >= wd || checkPos[1] < 0 || checkPos[1] > ht {
				continue
			}
			if checkPos[0] == endPos[0] && checkPos[1] == endPos[1] {
				return time
			}
			if checkPos[1] == 0 {
				if checkPos[0] == 1 {
					queue = append(queue, [3]int{checkPos[0], checkPos[1], time})
				}
				continue
			}
			if checkPos[1] == ht {
				if checkPos[0] == wd-1 {
					queue = append(queue, [3]int{checkPos[0], checkPos[1], time})
				}
				continue
			}
			for _, blizzard := range *(*horizontals)[checkPos[1]] {
				blizzardPos := blizzard.locationAtTime(time, ht, wd)
				if blizzardPos[0] == checkPos[0] && blizzardPos[1] == checkPos[1] {
					continue loop
				}
			}
			for _, blizzard := range *(*verticals)[checkPos[0]] {
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

func solve(horizontals *map[int]*[]*Blizzard, verticals *map[int]*[]*Blizzard, ht int, wd int) (int, int) {
	part1 := journey([2]int{1, 0}, [2]int{wd - 1, ht}, 0, horizontals, verticals, ht, wd)
	backToStart := journey([2]int{wd - 1, ht}, [2]int{1, 0}, part1, horizontals, verticals, ht, wd)
	part2 := journey([2]int{1, 0}, [2]int{wd - 1, ht}, backToStart, horizontals, verticals, ht, wd)
	return part1, part2
}
