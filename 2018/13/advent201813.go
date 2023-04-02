package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	tracks, carts := parseInput(input)
	fmt.Println(part1(tracks, carts))
}

type Cart struct {
	loc  [2]int
	dir  [2]int
	turn int
}

func (c *Cart) move(tracks *map[[2]int]string) {
	c.loc[0] += c.dir[0]
	c.loc[1] += c.dir[1]
	switch (*tracks)[c.loc] {
	case "/":
		c.dir[0], c.dir[1] = -c.dir[1], -c.dir[0]
	case "\\":
		c.dir[0], c.dir[1] = c.dir[1], c.dir[0]
	case "+":
		switch c.turn {
		case 0:
			if c.dir[0] == 0 {
				c.dir[0], c.dir[1] = c.dir[1], c.dir[0]
			} else {
				c.dir[0], c.dir[1] = -c.dir[1], -c.dir[0]
			}
			c.turn++
		case 1:
			c.turn++
		case 2:
			if c.dir[0] == 0 {
				c.dir[0], c.dir[1] = -c.dir[1], -c.dir[0]
			} else {
				c.dir[0], c.dir[1] = c.dir[1], c.dir[0]
			}
			c.turn = 0
		}
	case "|":
	case "-":
	default:
		panic("What happened?")
	}
}

func parseInput(input []string) (*map[[2]int]string, *[]*Cart) {
	tracks := map[[2]int]string{}
	carts := []*Cart{}
	for j, row := range input {
		for i, c := range row {
			loc := [2]int{i, j}
			switch c {
			case 'v':
				cart := Cart{
					loc:  loc,
					dir:  [2]int{0, 1},
					turn: 0,
				}
				carts = append(carts, &cart)
				tracks[loc] = "|"
			case '^':
				cart := Cart{
					loc:  loc,
					dir:  [2]int{0, -1},
					turn: 0,
				}
				carts = append(carts, &cart)
				tracks[loc] = "|"
			case '<':
				cart := Cart{
					loc:  loc,
					dir:  [2]int{-1, 0},
					turn: 0,
				}
				carts = append(carts, &cart)
				tracks[loc] = "-"
			case '>':
				cart := Cart{
					loc:  loc,
					dir:  [2]int{1, 0},
					turn: 0,
				}
				carts = append(carts, &cart)
				tracks[loc] = "-"
			case '|':
				fallthrough
			case '-':
				fallthrough
			case '/':
				fallthrough
			case '\\':
				fallthrough
			case '+':
				tracks[loc] = string(c)
			}
		}
	}
	return &tracks, &carts
}

func part1(tracks *map[[2]int]string, carts *[]*Cart) string {
	for {
		sort.Slice(*carts, func(i, j int) bool {
			return (*carts)[i].loc[1] < (*carts)[j].loc[1]
		})
		locs := map[[2]int]bool{}
		for _, cart := range *carts {
			if locs[cart.loc] {
				return fmt.Sprintf("%d,%d", cart.loc[0], cart.loc[1])
			}
			cart.move(tracks)
			locs[cart.loc] = true
		}
	}
	return ""
}
