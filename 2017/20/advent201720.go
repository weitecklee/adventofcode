package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(parseInput(input)))
	fmt.Println(part2(parseInput(input)))
}

type Particle struct {
	p   []int
	v   []int
	a   []int
	d   int
	pos string
}

func (p *Particle) move() {
	d := 0
	for i := 0; i < 3; i++ {
		p.v[i] += p.a[i]
		p.p[i] += p.v[i]
		d += int(math.Abs(float64(p.p[i])))
	}
	p.d = d
}

func parseInput(input []string) *[]Particle {
	particles := []Particle{}
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		nums := []int{}
		for _, m := range matches {
			n, _ := strconv.Atoi(m)
			nums = append(nums, n)
		}
		particle := Particle{}
		particle.p = nums[0:3]
		particle.v = nums[3:6]
		particle.a = nums[6:9]
		d := 0
		for i := 0; i < 3; i++ {
			d += int(math.Abs(float64(particle.p[i])))
		}
		particle.d = d
		particles = append(particles, particle)
	}
	return &particles
}

func part1(particles *[]Particle) int {
	chosen := -1
	streak := 0
	for {
		minDist := math.MaxInt
		currChosen := -1
		for i := range *particles {
			(*particles)[i].move()
			if (*particles)[i].d < minDist {
				minDist = (*particles)[i].d
				currChosen = i
			}
		}
		if currChosen == chosen {
			streak++
			if streak >= 10000 {
				break
			}
		} else {
			streak = 0
		}
		chosen = currChosen
	}
	return chosen
}

func part2(particles *[]Particle) int {
	streak := 0
	for {
		remainingParticles := []Particle{}
		posMap := map[string]int{}
		for i := range *particles {
			(*particles)[i].move()
			(*particles)[i].pos = fmt.Sprintf("%d,%d,%d", (*particles)[i].p[0], (*particles)[i].p[1], (*particles)[i].p[2])
			posMap[(*particles)[i].pos]++
		}
		for _, particle := range *particles {
			if posMap[particle.pos] == 1 {
				remainingParticles = append(remainingParticles, particle)
			}
		}
		if len(remainingParticles) == len(*particles) {
			streak++
			if streak >= 100 {
				break
			}
		} else {
			streak = 0
		}
		particles = &remainingParticles
	}
	return len(*particles)
}
