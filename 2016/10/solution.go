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
	parseInput(puzzleInput)
	fmt.Println(part1())
	fmt.Println(part2())
}

var (
	botRegex   = regexp.MustCompile(`^bot (\d+) gives low to (\w+) (\d+) and high to (\w+) (\d+)$`)
	valueRegex = regexp.MustCompile(`^value (\d+) goes to bot (\d+)$`)
	botMap     = make(map[int]*Bot)
	outputMap  = make(map[int]int)
)

type Destination struct {
	kind string
	id   int
}

type Bot struct {
	id         int
	microchips []int
	lowDst     *Destination
	highDst    *Destination
	history    map[int]struct{}
}

func (b *Bot) AddMicrochip(microchip int) {
	b.microchips = append(b.microchips, microchip)
	b.history[microchip] = struct{}{}
	if len(b.microchips) == 2 {
		b.AssessMicrochips()
	}
}

func (b *Bot) AssessMicrochips() {
	low, high := b.microchips[0], b.microchips[1]
	if low > high {
		low, high = high, low
	}
	b.microchips = b.microchips[:0]
	if b.lowDst.kind == "bot" {
		botMap[b.lowDst.id].AddMicrochip(low)
	} else {
		outputMap[b.lowDst.id] = low
	}
	if b.highDst.kind == "bot" {
		botMap[b.highDst.id].AddMicrochip(high)
	} else {
		outputMap[b.highDst.id] = high
	}
}

func NewBot(line string) *Bot {
	match := botRegex.FindStringSubmatch(line)
	if match == nil {
		panic(fmt.Sprintf("Error matching bot regex with : %s", line))
	}
	botId, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}
	lowId, err := strconv.Atoi(match[3])
	if err != nil {
		panic(err)
	}
	lowDst := Destination{match[2], lowId}
	highId, err := strconv.Atoi(match[5])
	if err != nil {
		panic(err)
	}
	highDst := Destination{match[4], highId}
	return &Bot{botId, []int{}, &lowDst, &highDst, map[int]struct{}{}}
}

func parseInput(puzzleInput []string) {
	sort.Strings(puzzleInput)
	for _, line := range puzzleInput {
		if line[0] == 'b' {
			bot := NewBot(line)
			botMap[bot.id] = bot
		} else {
			match := valueRegex.FindStringSubmatch(line)
			if match == nil {
				panic(fmt.Sprintf("Error matching value regex with : %s", line))
			}
			value, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			botId, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			if _, ok := botMap[botId]; !ok {
				panic(fmt.Sprintf("Cannot find bot %d", botId))
			}
			botMap[botId].AddMicrochip(value)
		}
	}
}

func part1() int {
	for _, bot := range botMap {
		if _, ok1 := bot.history[61]; ok1 {
			if _, ok2 := bot.history[17]; ok2 {
				return bot.id
			}
		}
	}
	return -1
}

func part2() int {
	return outputMap[0] * outputMap[1] * outputMap[2]
}
