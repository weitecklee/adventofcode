package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parsedInput := parseInput(string(data))
	fmt.Println(part1(parsedInput))
}

func parseInput(input string) []int {
	arr := []int{}
	for _, c := range input {
		n, _ := strconv.Atoi(string(c))
		arr = append(arr, n)
	}
	return arr
}

func game(cups *ring.Ring, turns int) *ring.Ring {
	cupMap := map[int]*ring.Ring{} // package Ring implements circular lists, how convenient!
	max := cups.Len()
	for i := 0; i < cups.Len(); i++ {
		cupMap[cups.Value.(int)] = cups
		cups = cups.Next()
	}
	for i := 0; i < turns; i++ {
		removed := cups.Unlink(3)
		dest := cups.Value.(int)
		for {
			dest--
			if dest == 0 {
				dest = max
			}
			inRemoved := false
			removed.Do(func(val any) {
				if val.(int) == dest {
					inRemoved = true
				}
			})
			if !inRemoved {
				break
			}
		}
		cupMap[dest].Link(removed)
		cups = cups.Next()
	}
	return cupMap[1]
}

func part1(input []int) string {
	cups := ring.New(len(input))
	for _, n := range input {
		cups.Value = n
		cups = cups.Next()
	}
	cups = game(cups, 100)
	ret := ""
	for i := 0; i < 8; i++ {
		cups = cups.Next()
		ret += strconv.Itoa(cups.Value.(int))
	}
	return ret
}
