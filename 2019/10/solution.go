package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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
	angles := calcAngles(h, w)
	ans1, stationPos := part1(asteroidMap, h, w, angles)
	fmt.Println(ans1)
	fmt.Println(part2(asteroidMap, h, w, stationPos, angles))
}

var variations = [][2]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func parseInput(data []string) (map[[2]int]struct{}, int, int) {
	asteroidMap := make(map[[2]int]struct{})
	for r := range data {
		for c := range data[r] {
			if data[r][c] == '#' {
				asteroidMap[[2]int{c, r}] = struct{}{}
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

func angleToCoord(x, y int) float64 {
	theta := math.Atan2(-float64(y), float64(x))
	angle := math.Pi/2 - theta
	if angle >= 0 {
		return angle
	}
	return angle + 2*math.Pi
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
	sort.SliceStable(angles, func(i, j int) bool {
		return angleToCoord(angles[i][0], angles[i][1]) < angleToCoord(angles[j][0], angles[j][1])
	})
	return angles
}

func part1(asteroidMap map[[2]int]struct{}, h, w int, angles [][2]int) (int, [2]int) {
	var res, count int
	var pos, stationPos [2]int
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
			stationPos = asteroid
		}
	}
	return res, stationPos
}

func part2(asteroidMap map[[2]int]struct{}, h, w int, stationPos [2]int, angles [][2]int) int {
	var asteroidsAtAngles []*[][2]int
	for _, angle := range angles {
		var asteroidsAtAngle [][2]int
		pos := [2]int{stationPos[0] + angle[0], stationPos[1] + angle[1]}
		for pos[0] >= 0 && pos[1] >= 0 && pos[0] < h && pos[1] < w {
			if _, ok := asteroidMap[pos]; ok {
				asteroidsAtAngle = append(asteroidsAtAngle, pos)
			}
			pos[0] += angle[0]
			pos[1] += angle[1]
		}
		if len(asteroidsAtAngle) > 0 {
			asteroidsAtAngles = append(asteroidsAtAngles, &asteroidsAtAngle)
		}
	}

	i := 0
	for len(asteroidsAtAngles) > 0 {
		j := 0
		for _, asteroids := range asteroidsAtAngles {
			if len(*asteroids) > 0 {
				i++
				if i == 200 {
					asteroid200 := (*asteroids)[0]
					return asteroid200[0]*100 + asteroid200[1]
				}
				*asteroids = (*asteroids)[1:]
			} else {
				asteroidsAtAngles[j] = asteroids
				j++
			}
		}
		asteroidsAtAngles = asteroidsAtAngles[:j]
	}

	return -1
}
