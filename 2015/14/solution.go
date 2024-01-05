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
	reindeer := parseInput(input)
	fmt.Println(part1(reindeer))
	fmt.Println(part2(reindeer))
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

func part2(reindeer map[string]Reindeer) int {
	runningScore := map[string]int{}
	for i := 1; i <= 2503; i++ {
		res := 0
		winners := []string{}
		for name, deer := range reindeer {
			curr := deer.Distance(i)
			if curr > res {
				res = curr
				winners = []string{name}
			} else if curr == res {
				winners = append(winners, name)
			}
		}
		for _, name := range winners {
			runningScore[name]++
		}
	}
	res := 0
	for _, score := range runningScore {
		if score > res {
			res = score
		}
	}
	return res
}
