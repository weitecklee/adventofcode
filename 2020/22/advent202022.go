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
	parsedInput := parseInput(input)
	fmt.Println(part1(parsedInput))
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
