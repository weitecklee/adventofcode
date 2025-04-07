package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
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
	puzzleInput := strings.Split(string(data), "\n")
	boss := parseInput(puzzleInput)
	shops := parseItemShop(ITEMSHOP)
	fmt.Println(solve(boss, shops))
}

const ITEMSHOP = `Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3`

type Item struct {
	name   string
	hp     int
	cost   int
	damage int
	armor  int
}

var playerBase = &Item{"Player", 100, 0, 0, 0}

type Character struct {
	Item
	equipment []*Item
}

func (ch *Character) Equip(item *Item) {
	ch.hp += item.hp
	ch.cost += item.cost
	ch.damage += item.damage
	ch.armor += item.armor
	ch.equipment = append(ch.equipment, item)
}

func (ch *Character) Attack(target *Character) {
	damage := ch.damage - target.armor
	if damage < 1 {
		damage = 1
	}
	target.hp -= damage
}

func NewCharacter(base *Item, equipment []*Item) *Character {
	ch := &Character{*base, make([]*Item, 0, len(equipment))}
	for _, item := range equipment {
		ch.Equip(item)
	}
	return ch
}

func parseItemShop(itemShop string) [][]*Item {
	shopStrings := strings.Split(itemShop, "\n\n")
	shops := make([][]*Item, len(shopStrings))
	itemRegex := regexp.MustCompile(`^(.*?)\s{2,}(\d+)\s+(\d+)\s+(\d+)$`)
	for i, part := range shopStrings {
		lines := strings.Split(part, "\n")
		items := make([]*Item, 0, len(lines)+1)
		for _, line := range lines[1:] {
			parts := itemRegex.FindStringSubmatch(line)
			if parts == nil {
				panic(fmt.Sprintf("Could not match itemRegex with: %s", line))
			}
			name := parts[1]
			hp := 0
			cost, err := strconv.Atoi(parts[2])
			if err != nil {
				panic(err)
			}
			damage, err := strconv.Atoi(parts[3])
			if err != nil {
				panic(err)
			}
			armor, err := strconv.Atoi(parts[4])
			if err != nil {
				panic(err)
			}
			items = append(items, &Item{name, hp, cost, damage, armor})
		}
		shops[i] = items
	}
	shops[1] = append(shops[1], &Item{"No Armor", 0, 0, 0, 0})
	shops[2] = append(shops[2], &Item{"No Ring", 0, 0, 0, 0})
	return shops
}

func parseInput(data []string) *Item {
	numRegex := regexp.MustCompile(`\d+`)
	hpString := numRegex.FindString(data[0])
	hp, err := strconv.Atoi(hpString)
	if err != nil {
		panic(err)
	}
	damageString := numRegex.FindString(data[1])
	damage, err := strconv.Atoi(damageString)
	if err != nil {
		panic(err)
	}
	armorString := numRegex.FindString(data[2])
	armor, err := strconv.Atoi(armorString)
	if err != nil {
		panic(err)
	}
	return &Item{"Boss", hp, 0, damage, armor}
}

func simulate(player, boss *Character) bool {
	for {
		player.Attack(boss)
		if boss.hp <= 0 {
			return true
		}
		boss.Attack(player)
		if player.hp <= 0 {
			return false
		}
	}
}

func solve(bossBase *Item, shops [][]*Item) (int, int) {
	var scenarios [][2]*Character
	weaponShop := shops[0]
	armorShop := shops[1]
	ringShop := shops[2]
	for _, weapon := range weaponShop {
		for _, armor := range armorShop {
			for _, ring1 := range ringShop {
				for _, ring2 := range ringShop {
					player := NewCharacter(playerBase, []*Item{weapon, armor, ring1, ring2})
					boss := NewCharacter(bossBase, []*Item{})
					scenarios = append(scenarios, [2]*Character{player, boss})
				}
			}
		}
	}

	part1 := math.MaxInt
	part2 := 0
	for _, scenario := range scenarios {
		player, boss := scenario[0], scenario[1]
		playerWon := simulate(player, boss)
		if playerWon {
			if player.cost < part1 {
				part1 = player.cost
			}
		} else {
			if player.cost > part2 {
				part2 = player.cost
			}
		}
	}

	return part1, part2
}
