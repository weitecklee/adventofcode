package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

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
	fmt.Println(part1(parseInput(strings.Split(string(data), "\n"))))
	fmt.Println(part2(parseInput(strings.Split(string(data), "\n"))))

}

type Moon struct {
	pos [3]int
	vel [3]int
}

func (m *Moon) potentialEnergy() int {
	return utils.AbsInt(m.pos[0]) + utils.AbsInt(m.pos[1]) + utils.AbsInt(m.pos[2])
}

func (m *Moon) kineticEnergy() int {
	return utils.AbsInt(m.vel[0]) + utils.AbsInt(m.vel[1]) + utils.AbsInt(m.vel[2])
}

func (m *Moon) totalEnergy() int {
	return m.potentialEnergy() * m.kineticEnergy()
}

func parseInput(data []string) []*Moon {
	moons := make([]*Moon, 0, len(data))
	numRegex := regexp.MustCompile(`-?\d+`)
	var n1, n2, n3 int
	var err error
	for _, line := range data {
		nums := numRegex.FindAllString(line, -1)
		if n1, err = strconv.Atoi(nums[0]); err != nil {
			panic(err)
		}
		if n2, err = strconv.Atoi(nums[1]); err != nil {
			panic(err)
		}
		if n3, err = strconv.Atoi(nums[2]); err != nil {
			panic(err)
		}
		moons = append(moons, &Moon{[3]int{n1, n2, n3}, [3]int{0, 0, 0}})
	}
	return moons
}

func gravitateMoons(moons []*Moon) {
	for i, moon1 := range moons {
		for _, moon2 := range moons[i+1:] {
			for k := range 3 {
				if moon1.pos[k] > moon2.pos[k] {
					moon1.vel[k]--
					moon2.vel[k]++
				} else if moon1.pos[k] < moon2.pos[k] {
					moon1.vel[k]++
					moon2.vel[k]--
				}
			}
		}
	}
	for _, moon := range moons {
		for k := range 3 {
			moon.pos[k] += moon.vel[k]
		}
	}
}

func part1(moons []*Moon) int {
	for range 1000 {
		gravitateMoons(moons)
	}
	res := 0
	for _, moon := range moons {
		res += moon.totalEnergy()
	}
	return res
}

func moonSignatures(moons []*Moon) [][8]int {
	signatures := make([][8]int, 3)
	for k := range 3 {
		signature := [8]int{moons[0].pos[k], moons[1].pos[k], moons[2].pos[k], moons[3].pos[k], moons[0].vel[k], moons[1].vel[k], moons[2].vel[k], moons[3].vel[k]}
		signatures[k] = signature
	}
	return signatures
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmSlice(s []int) int {
	res := s[0]
	for _, num := range s[1:] {
		res = lcm(res, num)
	}
	return res
}

func part2(moons []*Moon) int {
	if len(moons) != 4 {
		panic("Expected number of moons to be 4")
	}
	signatureMaps := make([]map[[8]int]int, 3)
	for k := range 3 {
		signatureMaps[k] = make(map[[8]int]int)
	}
	for k, signature := range moonSignatures(moons) {
		signatureMaps[k][signature] = 0
	}
	stepsToRepeat := make([]int, 3)
	var stepsFound bool
	i := 0
	for !stepsFound {
		i++
		gravitateMoons(moons)
		for k, signature := range moonSignatures(moons) {
			if stepsToRepeat[k] != 0 {
				continue
			}
			if j, ok := signatureMaps[k][signature]; ok {
				stepsToRepeat[k] = i - j
			} else {
				signatureMaps[k][signature] = i
			}
		}
		stepsFound = true
		for k := range 3 {
			if stepsToRepeat[k] == 0 {
				stepsFound = false
			}
		}
	}

	return lcmSlice(stepsToRepeat)
}
