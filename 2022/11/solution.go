package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	fmt.Println(part1(parseInput(strings.Split(string(data), "\n\n"))))
	fmt.Println(part2(parseInput(strings.Split(string(data), "\n\n"))))
}

var PRIME = 1

type Monkey struct {
	id             int
	items          []*Item
	oper           func(i *Item)
	test           func(i *Item) bool
	ifTrue         int
	ifFalse        int
	ifTrueMonkey   *Monkey
	ifFalseMonKey  *Monkey
	itemsInspected int
}

func NewMonkey(id int, items []int, operStr string, testN, trueN, falseN int) *Monkey {
	monkeyItems := make([]*Item, len(items))
	for i, n := range items {
		monkeyItems[i] = &Item{n}
	}

	operParts := strings.Split(operStr, " ")
	op := operParts[len(operParts)-2]
	var operFunc func(i *Item)
	if operParts[len(operParts)-1] == "old" {
		if op == "+" {
			operFunc = func(i *Item) {
				i.worryLevel += i.worryLevel
			}
		} else if op == "*" {
			operFunc = func(i *Item) {
				i.worryLevel *= i.worryLevel
			}
		}
	} else {
		val, err := strconv.Atoi(operParts[len(operParts)-1])
		if err != nil {
			panic(err)
		}
		if op == "+" {
			operFunc = func(i *Item) {
				i.worryLevel += val
			}
		} else if op == "*" {
			operFunc = func(i *Item) {
				i.worryLevel *= val
			}
		}
	}

	return &Monkey{
		id,
		monkeyItems,
		operFunc,
		func(i *Item) bool {
			return i.worryLevel%testN == 0
		},
		trueN,
		falseN,
		nil,
		nil,
		0,
	}
}

func (m *Monkey) AddItem(item *Item) {
	m.items = append(m.items, item)
}

func (m *Monkey) InspectItem(item *Item) {
	m.oper(item)
}

func (m *Monkey) TakeTurn(withRelief bool) {
	for len(m.items) > 0 {
		item := m.items[0]
		m.items = m.items[1:]
		m.InspectItem(item)
		if withRelief {
			Relief(item)
		} else {
			Normalize(item)
		}
		if m.test(item) {
			m.ifTrueMonkey.AddItem(item)
		} else {
			m.ifFalseMonKey.AddItem(item)
		}
		m.itemsInspected++
	}
}

type Item struct {
	worryLevel int
}

func Relief(item *Item) {
	item.worryLevel /= 3
}

func Normalize(item *Item) {
	item.worryLevel %= PRIME
}

func parseInput(data []string) map[int]*Monkey {
	PRIME = 1
	monkeyMap := make(map[int]*Monkey, len(data))
	for _, block := range data {
		lines := strings.Split(block, "\n")
		id := utils.ExtractInts(lines[0])[0]
		items := utils.ExtractInts(lines[1])
		testN := utils.ExtractInts(lines[3])[0]
		trueN := utils.ExtractInts(lines[4])[0]
		falseN := utils.ExtractInts(lines[5])[0]
		monkeyMap[id] = NewMonkey(id, items, lines[2], testN, trueN, falseN)
		PRIME *= testN
	}
	for _, monkey := range monkeyMap {
		monkey.ifTrueMonkey = monkeyMap[monkey.ifTrue]
		monkey.ifFalseMonKey = monkeyMap[monkey.ifFalse]
	}
	return monkeyMap
}

func calcMonkeyBusinessLevel(monkeyMap map[int]*Monkey) int {
	topTwo := [2]int{}
	for _, monkey := range monkeyMap {
		for i, n := range topTwo {
			if monkey.itemsInspected > n {
				copy(topTwo[i+1:], topTwo[i:])
				topTwo[i] = monkey.itemsInspected
				break
			}
		}
	}
	return topTwo[0] * topTwo[1]
}

func part1(monkeyMap map[int]*Monkey) int {
	for range 20 {
		for i := 0; i < len(monkeyMap); i++ {
			monkeyMap[i].TakeTurn(true)
		}
	}
	return calcMonkeyBusinessLevel(monkeyMap)
}

func part2(monkeyMap map[int]*Monkey) int {
	for range 10000 {
		for i := 0; i < len(monkeyMap); i++ {
			monkeyMap[i].TakeTurn(false)
		}
	}
	return calcMonkeyBusinessLevel(monkeyMap)
}
