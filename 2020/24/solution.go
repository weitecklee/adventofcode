package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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
	tiles := part1(input)
	fmt.Println(len(tiles))
	fmt.Println(part2(tiles))
}

func stringCoord(c [2]int) string {
	return strconv.Itoa(c[0]) + "," + strconv.Itoa(c[1])
}

func deStringCoord(coord string) [2]int {
	s := strings.Split(coord, ",")
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])
	return [2]int{a, b}
}

func part1(input []string) map[string]bool {
	tiles := map[string]bool{} // black is true, white is false
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
		coordStr := stringCoord(coord)
		tiles[coordStr] = !tiles[coordStr]
		if !tiles[coordStr] {
			delete(tiles, coordStr)
		}
	}
	return tiles
}

func part2(tiles map[string]bool) int {
	checkAdjacent := [6][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {2, 0}, {-2, 0}}
	for i := 0; i < 100; i++ {
		adjacentTiles := map[string]int{}
		tilesWith2BlackAdj := map[string]bool{}
		for coordStr := range tiles {
			coord := deStringCoord(coordStr)
			for _, check := range checkAdjacent {
				coordAdj := [2]int{coord[0] + check[0], coord[1] + check[1]}
				coordS := stringCoord(coordAdj)
				adjacentTiles[coordS]++
				if adjacentTiles[coordS] == 2 {
					tilesWith2BlackAdj[coordS] = true
				} else if adjacentTiles[coordS] > 2 {
					delete(tilesWith2BlackAdj, coordS)
				}
			}
		}
		tiles2 := map[string]bool{}
		for coordStr := range tiles {
			if adjacentTiles[coordStr] == 1 || adjacentTiles[coordStr] == 2 {
				tiles2[coordStr] = true
			}
		}
		for coordStr := range tilesWith2BlackAdj {
			if !tiles[coordStr] {
				tiles2[coordStr] = true
			}
		}
		tiles = tiles2
	}
	return len(tiles)
}
