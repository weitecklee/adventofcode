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
	asteroidMap, h, w := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(asteroidMap, h, w))
}

var variations = [][2]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func parseInput(data []string) (map[[2]int]struct{}, int, int) {
	asteroidMap := make(map[[2]int]struct{})
	for r := range data {
		for c := range data[r] {
			if data[r][c] == '#' {
				asteroidMap[[2]int{r, c}] = struct{}{}
			}
		}
	}
	return asteroidMap, len(data), len(data[0])
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func calcAngles(h, w int) [][2]int {
	var angles [][2]int
	angles = append(angles, [2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0})
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if gcd(i, j) == 1 {
				for _, variation := range variations {
					angles = append(angles, [2]int{variation[0] * i, variation[1] * j})
				}
			}
		}
	}
	return angles
}

func part1(asteroidMap map[[2]int]struct{}, h, w int) int {
	angles := calcAngles(h, w)
	var res, count int
	var pos [2]int
	for asteroid := range asteroidMap {
		count = 0
		for _, angle := range angles {
			pos = [2]int{asteroid[0] + angle[0], asteroid[1] + angle[1]}
			for pos[0] >= 0 && pos[1] >= 0 && pos[0] < h && pos[1] < w {
				if _, ok := asteroidMap[pos]; ok {
					count++
					break
				}
				pos[0] += angle[0]
				pos[1] += angle[1]
			}
		}
		if count > res {
			res = count
		}
	}
	return res
}
