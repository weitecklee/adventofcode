package main

import (
	"container/heap"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"runtime"

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
	puzzleInput := utils.ExtractInts(string(data))
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

type Spell struct {
	name   string
	costMP int
}

var (
	spells = []*Spell{
		{"Magic Missile", 53},
		{"Drain", 73},
		{"Shield", 113},
		{"Poison", 173},
		{"Recharge", 229},
	}
	playerHP = 50
	playerMP = 500
)

type Value struct {
	playerHP   int
	bossHP     int
	playerMP   int
	spentMP    int
	effects    map[string]int
	playerTurn bool
	armor      int
}

func solve(puzzleInput []int, isPart2 bool) int {
	bossHP := puzzleInput[0]
	bossDamage := puzzleInput[1]
	queue := utils.NewMinHeap[Value]()
	heap.Push(queue, &utils.Item[Value]{
		Priority: 0,
		Value:    Value{playerHP, bossHP, playerMP, 0, make(map[string]int), true, 0},
	})

	for len(queue.PriorityQueue) > 0 {
		item := heap.Pop(queue).(*utils.Item[Value])
		playerHP := item.Value.playerHP
		bossHP := item.Value.bossHP
		playerMP := item.Value.playerMP
		spentMP := item.Value.spentMP
		effects := item.Value.effects
		playerTurn := item.Value.playerTurn
		armor := item.Value.armor

		if isPart2 && playerTurn {
			playerHP--
			if playerHP <= 0 {
				continue
			}
		}

		for name, timer := range effects {
			timer--
			switch name {
			case "Shield":
				if timer == 0 {
					armor -= 7
				}
			case "Poison":
				bossHP -= 3
			case "Recharge":
				playerMP += 101
			}
			if timer == 0 {
				delete(effects, name)
			} else {
				effects[name] = timer
			}
		}

		if bossHP <= 0 {
			return spentMP
		}

		if playerTurn {
			for _, spell := range spells {
				if _, exists := effects[spell.name]; exists {
					continue
				}
				if spell.costMP > playerMP {
					continue
				}

				effects2 := maps.Clone(effects)
				playerHP2, bossHP2, armor2 := playerHP, bossHP, armor
				switch spell.name {
				case "Magic Missile":
					bossHP2 -= 4
				case "Drain":
					bossHP2 -= 2
					playerHP2 += 2
				case "Shield":
					effects2["Shield"] = 6
					armor2 += 7
				case "Poison":
					effects2["Poison"] = 6
				case "Recharge":
					effects2["Recharge"] = 5
				}

				heap.Push(queue, &utils.Item[Value]{
					Priority: bossHP2 + spentMP + spell.costMP,
					Value: Value{
						playerHP2, bossHP2, playerMP - spell.costMP, spentMP + spell.costMP, effects2, !playerTurn, armor2,
					}})
			}
		} else {
			damage := max(bossDamage-armor, 1)
			playerHP2 := playerHP - damage
			if playerHP2 > 0 {
				heap.Push(queue, &utils.Item[Value]{
					Priority: bossHP + spentMP,
					Value: Value{
						playerHP2, bossHP, playerMP, spentMP, maps.Clone(effects), !playerTurn, armor,
					}})
			}
		}
	}
	return -1
}

func part1(puzzleInput []int) int {
	return solve(puzzleInput, false)
}

func part2(puzzleInput []int) int {
	return solve(puzzleInput, true)
}
