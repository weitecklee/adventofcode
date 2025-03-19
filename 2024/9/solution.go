package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	puzzleInput := parseInput(string(data))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func parseInput(data string) []int {
	puzzleInput := make([]int, len(data))
	for i, ch := range data {
		puzzleInput[i] = int(ch - '0')
	}
	return puzzleInput
}

func part1(puzzleInput []int) int {
	// use double pointers, put in `right` files into `left` empty blocks,
	// compute checksum as we go
	left := 0
	right := len(puzzleInput) - 1
	if right%2 == 1 {
		right--
	}
	var checksum, pos int
	waitingToBeSlotted := puzzleInput[right] // file blocks waiting to be slotted
	for left < right {
		if left%2 == 1 {
			// odd index -> empty block
			for range puzzleInput[left] {
				if waitingToBeSlotted == 0 {
					// still have empty space, find the next file block
					right -= 2
					waitingToBeSlotted = puzzleInput[right]
				}
				checksum += right / 2 * pos
				pos++
				waitingToBeSlotted--
			}
		} else {
			// even index -> file block
			for range puzzleInput[left] {
				checksum += left / 2 * pos
				pos++
			}
		}
		left++
	}
	for range waitingToBeSlotted {
		checksum += right / 2 * pos
		pos++
	}
	return checksum
}

func part2(puzzleInput []int) int {
	posMap := make(map[int]int, len(puzzleInput))
	// map of (file/empty) block to position, position points to start of block
	pos := 0
	for i, n := range puzzleInput {
		posMap[i] = pos
		pos += n
	}

	checksum := 0
	puzzleInput2 := slices.Clone(puzzleInput)
	right := len(puzzleInput) - 1
	if right%2 == 1 {
		right--
	}
	for right >= 0 {
		// find first empty block that fits
		pos := 1
		for pos < right && puzzleInput2[pos] < puzzleInput2[right] {
			pos += 2
		}

		if pos < right {
			// if found, slot file blocks in
			// have to account for if any files previously slotted in
			// so must compare old puzzleInput[pos] and new puzzleInput2[pos]
			for j := range puzzleInput2[right] {
				checksum += (puzzleInput[pos] - puzzleInput2[pos] + posMap[pos] + j) * (right / 2)
			}
			puzzleInput2[pos] -= puzzleInput2[right]
			puzzleInput2[right] = 0
		}

		right -= 2
	}

	// add up files that were not moved
	for i := 0; i < len(puzzleInput2); i += 2 {
		for j := range puzzleInput2[i] {
			checksum += (posMap[i] + j) * i / 2
		}
	}

	return checksum
}
