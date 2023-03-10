package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
}

func part1(input []string) int {
	tiles := map[string]bool{} // black is true, white is false
	count := 0
	for _, line := range input {
		coord := [2]int{0, 0}
		i := 0
		for i < len(line) {
			dir := string(line[i])
			if dir == "s" || dir == "n" {
				i++
				dir += string(line[i])
			}
			switch dir {
			case "ne":
				coord[0]++
				coord[1]++
			case "se":
				coord[0]++
				coord[1]--
			case "e":
				coord[0] += 2
			case "nw":
				coord[0]--
				coord[1]++
			case "sw":
				coord[0]--
				coord[1]--
			case "w":
				coord[0] -= 2
			default:
				panic("Unexpected direction")
			}
			i++
		}
		coordStr := strconv.Itoa(coord[0]) + "," + strconv.Itoa(coord[1])
		tiles[coordStr] = !tiles[coordStr]
		if tiles[coordStr] {
			count++
		} else {
			count--
		}
	}
	return count
}
