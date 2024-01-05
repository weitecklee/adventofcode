package main

import (
	"container/ring"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	parsedInput := parseInput(string(data))
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
}

func parseInput(input string) []int {
	arr := []int{}
	for _, c := range input {
		n, _ := strconv.Atoi(string(c))
		arr = append(arr, n)
	}
	return arr
}

func game(cups *ring.Ring, turns int, cupMap map[int]*ring.Ring) *ring.Ring {
	max := cups.Len()
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
	cups := ring.New(len(input)) // package Ring implements circular lists, how convenient!
	cupMap := make(map[int]*ring.Ring, 9)
	for _, n := range input {
		cups.Value = n
		cupMap[n] = cups
		cups = cups.Next()
	}
	cups = game(cups, 100, cupMap)
	ret := ""
	for i := 0; i < 8; i++ {
		cups = cups.Next()
		ret += strconv.Itoa(cups.Value.(int))
	}
	return ret
}

func part2(input []int) int {
	cups := ring.New(1000000)
	cupMap := make(map[int]*ring.Ring, 1000000)
	for _, n := range input {
		cups.Value = n
		cupMap[n] = cups
		cups = cups.Next()
	}
	for i := len(input) + 1; i <= 1000000; i++ {
		cups.Value = i
		cupMap[i] = cups
		cups = cups.Next()
	}
	cups = game(cups, 10000000, cupMap)

	return cups.Next().Value.(int) * cups.Next().Next().Value.(int)
}
