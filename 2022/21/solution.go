package main

import (
	"fmt"
	"math"
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
	fmt.Println(part1(parseInput(input)))
	fmt.Println(part2(parseInput(input)))
}

type Monkey struct {
	value     float64
	monkey1   *Monkey
	monkey2   *Monkey
	operation string
}

func (m *Monkey) Value() float64 {
	if m.value > 0 { // assume no monkey has value zero
		return m.value
	}
	switch m.operation {
	case "+":
		return m.monkey1.Value() + m.monkey2.Value()
	case "*":
		return m.monkey1.Value() * m.monkey2.Value()
	case "-":
		return m.monkey1.Value() - m.monkey2.Value()
	case "/":
		return m.monkey1.Value() / m.monkey2.Value()
	}
	return -1
}

func parseInput(input []string) *map[string]*Monkey {
	monkeys := map[string]*Monkey{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		curr := parts[0][:4]
		if _, ok := monkeys[curr]; !ok {
			monkey := Monkey{}
			monkeys[curr] = &monkey
		}
		if len(parts) > 2 {
			for _, name := range []string{parts[1], parts[3]} {
				if _, ok := monkeys[name]; !ok {
					monkey := Monkey{}
					monkeys[name] = &monkey
				}
			}
			monkeys[curr].monkey1 = monkeys[parts[1]]
			monkeys[curr].monkey2 = monkeys[parts[3]]
			monkeys[curr].operation = parts[2]
		} else {
			num, _ := strconv.Atoi(parts[1])
			monkeys[curr].value = float64(num)
		}
	}
	return &monkeys
}

func part1(monkeys *map[string]*Monkey) int {
	return int((*monkeys)["root"].Value())
}

func part2(monkeys *map[string]*Monkey) int {
	// find which side "humn" is on
	monkeyh := (*monkeys)["root"].monkey1
	monkey2 := (*monkeys)["root"].monkey2
	monkeyhvalue := monkeyh.Value()
	monkey2value := monkey2.Value()
	(*monkeys)["humn"].value = (*monkeys)["humn"].value * 2
	correlation := 1 // flag for positive/negative correlation
	monkeyhvalue2 := monkeyh.Value()
	if monkeyhvalue == monkeyhvalue2 {
		if monkey2.Value() < monkey2value {
			correlation = -1
		}
		monkeyh = monkey2
		monkey2 = (*monkeys)["root"].monkey1
	} else if monkeyhvalue2 < monkeyhvalue {
		correlation = -1
	}
	valueToMatch := monkey2.Value()
	lo := 1.0
	hi := math.MaxFloat64
	for lo < hi {
		mi := math.Floor(lo + (hi-lo)/2)
		(*monkeys)["humn"].value = mi
		monkeyhvalue = monkeyh.Value()
		if monkeyhvalue == valueToMatch {
			lo = mi
			break
		}
		if monkeyhvalue < valueToMatch {
			if correlation > 0 {
				lo = mi + 1
			} else {
				hi = mi - 1
			}
		} else {
			if correlation > 0 {
				hi = mi - 1
			} else {
				lo = mi + 1
			}
		}
	}
	return int(lo)
}
