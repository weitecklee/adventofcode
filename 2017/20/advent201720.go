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
}

type Particle struct {
	p []int
	v []int
	a []int
	d int
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

func parseInput(input []string) []Particle {
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
	return particles
}

func part1(input []Particle) int {
	chosen := -1
	streak := 0
	for {
		minDist := math.MaxInt
		currChosen := -1
		for i, particle := range input {
			particle.move()
			if particle.d < minDist {
				minDist = particle.d
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
