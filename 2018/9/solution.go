package main

import (
	"container/ring"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/weitecklee/adventofcode/utils"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	puzzleInput := utils.ExtractInts(string(data))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func newMarble(n int) *ring.Ring {
	marble := ring.New(1)
	marble.Value = n
	return marble
}

func play(nPlayers, lastPoints int) int {
	current := ring.New(1)
	current.Value = 0
	currPlayer := 0
	scores := make(map[int]int, nPlayers)
	for n := range lastPoints {
		if n%23 == 0 {
			scores[currPlayer] += n
			for range 8 {
				current = current.Prev()
			}
			removed := current.Unlink(1)
			scores[currPlayer] += removed.Value.(int)
			current = current.Next()
		} else {
			current = current.Next()
			marble := newMarble(n)
			current.Link(marble)
			current = marble
		}
		currPlayer = (currPlayer + 1) % nPlayers
	}
	res := 0
	for _, score := range scores {
		if score > res {
			res = score
		}
	}
	return res
}

func part1(nums []int) int {
	return play(nums[0], nums[1])
}

func part2(nums []int) int {
	return play(nums[0], nums[1]*100)
}
