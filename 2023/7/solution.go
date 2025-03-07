package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	hands := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(hands))
	fmt.Println(part2(hands))
}

var cardStrengthMap = makeCardStrengthMap("23456789TJQKA")
var cardStrengthMap2 = makeCardStrengthMap("J23456789TQKA")

func makeCardStrengthMap(cardsInOrder string) map[byte]int {
	res := make(map[byte]int, len(cardsInOrder))
	for i := range cardsInOrder {
		res[cardsInOrder[i]] = i
	}
	return res
}

type Hand struct {
	cards     string
	bid       int
	strength  int
	strength2 int
}

func NewHand(line string) *Hand {
	parts := strings.Split(line, " ")
	bid, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return &Hand{parts[0], bid, calcHandStrength(parts[0], false), calcHandStrength(parts[0], true)}
}

func (h1 *Hand) play(h2 *Hand) bool {
	if h1.strength != h2.strength {
		return h1.strength < h2.strength
	}
	for i := range h1.cards {
		if h1.cards[i] != h2.cards[i] {
			return cardStrengthMap[h1.cards[i]] < cardStrengthMap[h2.cards[i]]
		}
	}
	return false
}

func (h1 *Hand) play2(h2 *Hand) bool {
	if h1.strength2 != h2.strength2 {
		return h1.strength2 < h2.strength2
	}
	for i := range h1.cards {
		if h1.cards[i] != h2.cards[i] {
			return cardStrengthMap2[h1.cards[i]] < cardStrengthMap2[h2.cards[i]]
		}
	}
	return false
}

func calcHandStrength(cards string, withJokers bool) int {
	cardMap := make(map[rune]int, len(cards))
	for _, card := range cards {
		cardMap[card]++
	}
	if withJokers {
		nJokers := cardMap['J']
		if nJokers == 5 {
			return 6
		}
		delete(cardMap, 'J')
		nMax := 0
		var cardToPretend rune
		for card, n := range cardMap {
			if n > nMax {
				nMax = n
				cardToPretend = card
			}
		}
		cardMap[cardToPretend] += nJokers
	}
	if len(cardMap) == 1 {
		// five of a kind
		return 6
	}
	if len(cardMap) == 5 {
		// high card
		return 0
	}
	if len(cardMap) == 4 {
		// one pair
		return 1
	}
	if len(cardMap) == 2 {
		for _, n := range cardMap {
			if n == 4 || n == 1 {
				// four a kind
				return 5
			}
			// full house
			return 4
		}
	}
	for _, n := range cardMap {
		if n == 3 {
			// three of a kind
			return 3
		} else if n == 2 {
			// return two pair
			return 2
		}
	}
	return -1
}

func parseInput(data []string) []*Hand {
	hands := make([]*Hand, len(data))
	for i, line := range data {
		hands[i] = NewHand(line)
	}
	return hands
}

func part1(hands []*Hand) int {
	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].play(hands[j])
	})
	res := 0
	for i, hand := range hands {
		res += (i + 1) * hand.bid
	}
	return res
}

func part2(hands []*Hand) int {
	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].play2(hands[j])
	})
	res := 0
	for i, hand := range hands {
		res += (i + 1) * hand.bid
	}
	return res
}
