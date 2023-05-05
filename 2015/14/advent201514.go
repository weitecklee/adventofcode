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
	reindeer := parseInput(input)
	fmt.Println(part1(reindeer))
}

type Reindeer struct {
	speed    int
	flyTime  int
	restTime int
}

func (r *Reindeer) Distance(time int) int {
	n := time / (r.flyTime + r.restTime)
	leftOver := time - n*(r.flyTime+r.restTime)
	if leftOver > r.flyTime {
		leftOver = r.flyTime
	}
	return (n*r.flyTime + leftOver) * r.speed
}

func parseInput(input []string) map[string]Reindeer {
	reindeer := map[string]Reindeer{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		speed, _ := strconv.Atoi(parts[3])
		flyTime, _ := strconv.Atoi(parts[6])
		restTime, _ := strconv.Atoi(parts[13])
		reindeer[parts[0]] = Reindeer{
			speed:    speed,
			flyTime:  flyTime,
			restTime: restTime,
		}
	}
	return reindeer
}

func part1(reindeer map[string]Reindeer) int {
	res := 0
	for _, deer := range reindeer {
		curr := deer.Distance(2503)
		if curr > res {
			res = curr
		}
	}
	return res
}
