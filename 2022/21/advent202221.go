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
	fmt.Println(part1(parseInput(input)))
}

type Monkey struct {
	value     int
	monkey1   *Monkey
	monkey2   *Monkey
	operation string
}

func (m *Monkey) Value() int {
	if m.value > 0 {
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
			monkeys[curr].value = num
		}
	}
	for _, monkey := range monkeys {
		// assume no monkey has value of zero
		if monkey.value == 0 {
			monkey.value = monkey.Value()
		}
	}
	return &monkeys
}

func part1(monkeys *map[string]*Monkey) int {
	return (*monkeys)["root"].value
}
