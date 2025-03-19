package main

import (
	"fmt"
	"math"
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
	puzzleInput := parseInput(strings.Split(string(data), "\n"))
	regions := createRegions(puzzleInput)
	fmt.Println(part1(regions))
	fmt.Println(part2(regions))
}

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type Region struct {
	points map[[2]int]struct{}
}

func NewRegion() *Region {
	return &Region{make(map[[2]int]struct{})}
}

func (r *Region) Area() int {
	return len(r.points)
}

func (r *Region) Perimeter() int {
	res := 0
	for point := range r.points {
		for _, d := range directions {
			tmp := [2]int{point[0] + d[0], point[1] + d[1]}
			if _, ok := r.points[tmp]; !ok {
				res++
			}
		}
	}
	return res
}

func (r *Region) Price() int {
	return r.Area() * r.Perimeter()
}

func (r *Region) Sides() int {
	if len(r.points) <= 2 {
		return 4
	}
	rMin := math.MaxInt
	rMax := math.MinInt
	cMin, cMax := rMin, rMax
	var _r, _c int
	for point := range r.points {
		_r, _c = point[0], point[1]
		if _r < rMin {
			rMin = _r
		}
		if _r > rMax {
			rMax = _r
		}
		if _c < cMin {
			cMin = _c
		}
		if _c > cMax {
			cMax = _c
		}
	}
	res := 0
	var e, f, g, h bool
	for _r := rMin; _r <= rMax+1; _r++ {
		for _c := cMin; _c <= cMax+1; _c++ {
			// e | f
			//-------
			// g | h
			_, e = r.points[[2]int{_r - 1, _c - 1}]
			_, f = r.points[[2]int{_r - 1, _c}]
			_, g = r.points[[2]int{_r, _c - 1}]
			_, h = r.points[[2]int{_r, _c}]
			/*
				Corners counted (twice to account for pairs):
				■■  ■■  ◻◻  ◻◻
				■◻  ◻■  ◻■  ■◻
			*/
			if h != g && ((h == f && f == e) || h != f) {
				res += 2
			}
		}
	}
	return res
}

func (r *Region) DiscountedPrice() int {
	return r.Area() * r.Sides()
}

func parseInput(data []string) [][]byte {
	puzzleInput := make([][]byte, len(data))
	for i := range data {
		puzzleInput[i] = []byte(data[i])
	}
	return puzzleInput
}

func createRegions(puzzleInput [][]byte) []*Region {
	var regions []*Region
	for r := range puzzleInput {
		for c := range puzzleInput[r] {
			if puzzleInput[r][c] != ' ' {
				regions = append(regions, mapRegion(puzzleInput, r, c))
			}
		}
	}
	return regions

}

func mapRegion(puzzleInput [][]byte, r, c int) *Region {
	region := NewRegion()
	queue := [][2]int{{r, c}}
	regionChr := puzzleInput[r][c]
	puzzleInput[r][c] = ' '
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		region.points[curr] = struct{}{}
		for _, d := range directions {
			r, c := curr[0]+d[0], curr[1]+d[1]
			if r < 0 || c < 0 || r >= len(puzzleInput) || c >= len(puzzleInput[0]) {
				continue
			}
			if puzzleInput[r][c] != regionChr {
				continue
			}
			puzzleInput[r][c] = ' '
			point := [2]int{r, c}
			queue = append(queue, point)
		}
	}
	return region
}

func part1(regions []*Region) int {
	res := 0
	for _, region := range regions {
		res += region.Price()
	}
	return res
}

func part2(regions []*Region) int {
	res := 0
	for _, region := range regions {
		res += region.DiscountedPrice()
	}
	return res
}
