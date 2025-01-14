package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
	cuboids := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(cuboids))
	fmt.Println(part2(cuboids))
}

type Cuboid struct {
	state  bool
	xRange [2]int
	yRange [2]int
	zRange [2]int
}

func parseInput(input []string) []Cuboid {
	cuboids := []Cuboid{}
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range input {
		state := true
		if line[1] == 'f' {
			state = false
		}
		matches := re.FindAllString(line, -1)
		nums := []int{}
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			nums = append(nums, num)
		}
		cuboids = append(cuboids, Cuboid{
			state:  state,
			xRange: [2]int{nums[0], nums[1]},
			yRange: [2]int{nums[2], nums[3]},
			zRange: [2]int{nums[4], nums[5]},
		})
	}
	return cuboids
}

func part1(cuboids []Cuboid) int {
	cubes := map[[3]int]bool{}
	for _, cuboid := range cuboids {
		xMin := -50
		if cuboid.xRange[0] > xMin {
			xMin = cuboid.xRange[0]
		}
		xMax := 50
		if cuboid.xRange[1] < xMax {
			xMax = cuboid.xRange[1]
		}
		yMin := -50
		if cuboid.yRange[0] > yMin {
			yMin = cuboid.yRange[0]
		}
		yMax := 50
		if cuboid.yRange[1] < yMax {
			yMax = cuboid.yRange[1]
		}
		zMin := -50
		if cuboid.zRange[0] > zMin {
			zMin = cuboid.zRange[0]
		}
		zMax := 50
		if cuboid.zRange[1] < zMax {
			zMax = cuboid.zRange[1]
		}
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				for k := zMin; k <= zMax; k++ {
					cubes[[3]int{i, j, k}] = cuboid.state
				}
			}
		}
	}
	count := 0
	for _, on := range cubes {
		if on {
			count++
		}
	}
	return count
}

func part2(cuboids []Cuboid) int {
	var overlaps []Cuboid
	for _, cuboid1 := range cuboids {
		var tmp []Cuboid
		if cuboid1.state {
			tmp = append(tmp, cuboid1)
		}
		for _, cuboid2 := range overlaps {
			xMin := max(cuboid1.xRange[0], cuboid2.xRange[0])
			xMax := min(cuboid1.xRange[1], cuboid2.xRange[1])
			if xMin > xMax {
				continue
			}
			yMin := max(cuboid1.yRange[0], cuboid2.yRange[0])
			yMax := min(cuboid1.yRange[1], cuboid2.yRange[1])
			if yMin > yMax {
				continue
			}
			zMin := max(cuboid1.zRange[0], cuboid2.zRange[0])
			zMax := min(cuboid1.zRange[1], cuboid2.zRange[1])
			if zMin > zMax {
				continue
			}
			tmp = append(tmp, Cuboid{
				state:  !cuboid2.state,
				xRange: [2]int{xMin, xMax},
				yRange: [2]int{yMin, yMax},
				zRange: [2]int{zMin, zMax},
			})
		}
		overlaps = append(overlaps, tmp...)
	}
	res := 0
	for _, cuboid := range overlaps {
		size := (cuboid.xRange[1] - cuboid.xRange[0] + 1) * (cuboid.yRange[1] - cuboid.yRange[0] + 1) * (cuboid.zRange[1] - cuboid.zRange[0] + 1)
		if cuboid.state {
			res += size
		} else {
			res -= size
		}
	}
	return res
}
