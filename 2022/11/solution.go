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
	parseInput(strings.Split(string(data), "\n\n"))
	fmt.Println(part1())
}

var monkeyMap = make(map[int]*Monkey)

type Monkey struct {
	id             int
	items          []*Item
	oper           func(i *Item)
	test           func(i *Item) bool
	ifTrue         int
	ifFalse        int
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
		0,
	}
}

func (m *Monkey) AddItem(item *Item) {
	m.items = append(m.items, item)
}

func (m *Monkey) InspectItem(item *Item) {
	m.oper(item)
}

func (m *Monkey) TakeTurn() {
	for len(m.items) > 0 {
		item := m.items[0]
		m.items = m.items[1:]
		m.InspectItem(item)
		Relief(item)
		if m.test(item) {
			monkeyMap[m.ifTrue].AddItem(item)
		} else {
			monkeyMap[m.ifFalse].AddItem(item)
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

func parseInput(data []string) {
	for _, block := range data {
		lines := strings.Split(block, "\n")
		id := utils.ExtractInts(lines[0])[0]
		items := utils.ExtractInts(lines[1])
		testN := utils.ExtractInts(lines[3])[0]
		trueN := utils.ExtractInts(lines[4])[0]
		falseN := utils.ExtractInts(lines[5])[0]
		monkeyMap[id] = NewMonkey(id, items, lines[2], testN, trueN, falseN)
	}
}

func part1() int {
	for range 20 {
		for i := 0; i < len(monkeyMap); i++ {
			monkeyMap[i].TakeTurn()
		}
	}
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
