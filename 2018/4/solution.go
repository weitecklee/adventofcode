package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
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
	puzzleInput := strings.Split(string(data), "\n")
	guardMap := parseInput(puzzleInput)
	fmt.Println(part1(guardMap))
	fmt.Println(part2(guardMap))
}

var timestampRegex = regexp.MustCompile(`:(\d\d)\] (\w+)`)

type Guard struct {
	id                 int
	timestamps         []string
	sleepsAtMinute     [60]int
	totalSleeps        int
	mostSleepsAtMinute int
}

func NewGuard(id int) *Guard {
	return &Guard{id, []string{}, [60]int{}, 0, 0}
}

func (g *Guard) Assess() {
	var sleepStart, sleepEnd int
	for _, timestamp := range g.timestamps {
		match := timestampRegex.FindStringSubmatch(timestamp)
		if match == nil {
			panic(fmt.Sprintf("Error matching regex for timestamp: %s", timestamp))
		}
		if match[2] == "falls" {
			minute, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			sleepStart = minute
		} else if match[2] == "wakes" {
			minute, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			sleepEnd = minute
			for min := sleepStart; min < sleepEnd; min++ {
				g.sleepsAtMinute[min]++
			}
		}
	}
	var totalSleeps, mostSleeps, mostSleepsAtMinute int
	for minute, sleeps := range g.sleepsAtMinute {
		totalSleeps += sleeps
		if sleeps > mostSleeps {
			mostSleeps = sleeps
			mostSleepsAtMinute = minute
		}
	}
	g.totalSleeps = totalSleeps
	g.mostSleepsAtMinute = mostSleepsAtMinute
}

func parseInput(puzzleInput []string) map[int]*Guard {
	sort.Strings(puzzleInput)
	guardRegex := regexp.MustCompile(`Guard #(\d+)`)
	guardMap := make(map[int]*Guard)
	var currGuard *Guard
	for _, line := range puzzleInput {
		if match := guardRegex.FindStringSubmatch(line); match != nil {
			id, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			if _, ok := guardMap[id]; !ok {
				guardMap[id] = NewGuard(id)
			}
			currGuard = guardMap[id]
		}
		currGuard.timestamps = append(currGuard.timestamps, line)
	}
	for _, guard := range guardMap {
		guard.Assess()
	}
	return guardMap
}

func part1(guardMap map[int]*Guard) int {
	var guardId, minuteAsleep, mostMinutesAsleep int
	for _, guard := range guardMap {
		if guard.totalSleeps > mostMinutesAsleep {
			mostMinutesAsleep = guard.totalSleeps
			minuteAsleep = guard.mostSleepsAtMinute
			guardId = guard.id
		}
	}
	return guardId * minuteAsleep
}

func part2(guardMap map[int]*Guard) int {
	var guardId, minuteAsleep, mostSleepsAtMinute int
	for _, guard := range guardMap {
		if guard.sleepsAtMinute[guard.mostSleepsAtMinute] > mostSleepsAtMinute {
			mostSleepsAtMinute = guard.sleepsAtMinute[guard.mostSleepsAtMinute]
			minuteAsleep = guard.mostSleepsAtMinute
			guardId = guard.id
		}
	}
	return guardId * minuteAsleep
}
