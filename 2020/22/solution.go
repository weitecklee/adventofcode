package main

import (
	"fmt"
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
	parsedInput := parseInput(input)
	fmt.Println(part1(parsedInput))
	fmt.Println(part2(parsedInput))
}

type Player struct {
	deck []int
}

func (p *Player) draw() int {
	card := p.deck[0]
	p.deck = p.deck[1:]
	return card
}

func (p *Player) add(cards ...int) {
	p.deck = append(p.deck, cards...)
}

func (p Player) score() int {
	score := 0
	for i, card := range p.deck {
		score += card * (len(p.deck) - i)
	}
	return score
}

func parseInput(input []string) []Player {
	players := []Player{}
	player1 := Player{}
	player2 := Player{}
	row := 1
	for len(input[row]) > 0 {
		card, _ := strconv.Atoi(input[row])
		player1.add(card)
		row++
	}
	row += 2
	for row < len(input) && len(input[row]) > 0 {
		card, _ := strconv.Atoi(input[row])
		player2.add(card)
		row++
	}
	players = append(players, player1, player2)
	return players
}

func part1(players []Player) int {
	player1 := players[0]
	player2 := players[1]
	for len(player1.deck) > 0 && len(player2.deck) > 0 {
		card1 := player1.draw()
		card2 := player2.draw()
		if card1 > card2 {
			player1.add(card1, card2)
		} else {
			player2.add(card2, card1)
		}
	}
	if len(player1.deck) > 0 {
		return player1.score()
	}
	return player2.score()
}

func recordHistory(player1 Player, player2 Player) string {
	history1 := []string{}
	history2 := []string{}
	for _, card := range player1.deck {
		history1 = append(history1, strconv.Itoa(card))
	}
	for _, card := range player2.deck {
		history2 = append(history2, strconv.Itoa(card))
	}
	return strings.Join(history1, ",") + "|" + strings.Join(history2, ",")
}

func recurCombat(player1 Player, player2 Player) (int, int) {
	history := map[string]bool{}
	for len(player1.deck) > 0 && len(player2.deck) > 0 {
		currHistory := recordHistory(player1, player2)
		if history[currHistory] {
			return 1, player1.score()
		}
		history[currHistory] = true
		card1 := player1.draw()
		card2 := player2.draw()
		if len(player1.deck) >= card1 && len(player2.deck) >= card2 {
			subPlayer1 := Player{}
			subPlayer2 := Player{}
			subdeck1 := make([]int, card1)
			subdeck2 := make([]int, card2)
			copy(subdeck1, player1.deck)
			copy(subdeck2, player2.deck)
			subPlayer1.deck = subdeck1
			subPlayer2.deck = subdeck2
			winner, _ := recurCombat(subPlayer1, subPlayer2)
			if winner == 1 {
				player1.add(card1, card2)
			} else {
				player2.add(card2, card1)
			}
		} else if card1 > card2 {
			player1.add(card1, card2)
		} else {
			player2.add(card2, card1)
		}
	}
	if len(player1.deck) > 0 {
		return 1, player1.score()
	}
	return 2, player2.score()
}

func part2(players []Player) int {
	player1 := players[0]
	player2 := players[1]
	_, score := recurCombat(player1, player2)
	return score
}
