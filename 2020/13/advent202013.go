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
	data, err := os.ReadFile("input202013.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	timestamp, _ := strconv.Atoi(input[0])
	r := regexp.MustCompile(`\d+`)
	departs := r.FindAllString(input[1], -1)
	earliest := math.MaxInt
	busID := 0
	for _, departStr := range departs {
		depart, _ := strconv.Atoi(departStr)
		early := depart*int(math.Ceil(float64(timestamp)/float64(depart))) - timestamp
		if early < earliest {
			earliest = early
			busID = depart
		}
	}
	return earliest * busID
}

func part2(input []string) int {
	departs := strings.Split(input[1], ",")
	schedule := [][]int{}
	for i, departStr := range departs {
		if departStr != "x" {
			depart, _ := strconv.Atoi(departStr)
			schedule = append(schedule, []int{depart, i})
		}
	}
	timestamp := schedule[0][0]
	period := schedule[0][0]
	for _, time := range schedule[1:] {
		depart, id := time[0], time[1]
		for (timestamp+id)%depart > 0 {
			timestamp += period
		}
		period *= depart
	}
	return timestamp
}
