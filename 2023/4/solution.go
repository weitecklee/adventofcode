package main

import (
	"fmt"
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
	cards := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(cards))
	fmt.Println(part2(cards))
}

type Card struct {
	number            int
	winningNumbersMap map[int]struct{}
	numbersYouHave    []int
	copies            int
}

func NewCard(n int, winningNumbers, numbersYouHave []int) *Card {
	winningNumbersMap := make(map[int]struct{}, len(winningNumbers))
	for _, num := range winningNumbers {
		winningNumbersMap[num] = struct{}{}
	}
	return &Card{n, winningNumbersMap, numbersYouHave, 1}
}

func (c *Card) CountWinningNumbers() int {
	count := 0
	for _, n := range c.numbersYouHave {
		if _, exists := c.winningNumbersMap[n]; exists {
			count++
		}
	}
	return count

}

func (c *Card) Points() int {
	count := c.CountWinningNumbers()
	if count == 0 {
		return 0
	}
	return 1 << (count - 1)
}

var numRegex = regexp.MustCompile(`\d+`)

func convertToInts(s string) []int {
	match := numRegex.FindAllString(s, -1)
	nums := make([]int, len(match))
	for i, str := range match {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

func parseInput(data []string) []*Card {
	cards := make([]*Card, len(data))
	for i, line := range data {
		parts := strings.Split(line, "|")
		part0parts := strings.Split(parts[0], ":")
		cards[i] = NewCard(convertToInts(part0parts[0])[0], convertToInts(part0parts[1]), convertToInts(parts[1]))
	}
	return cards
}

func part1(cards []*Card) int {
	res := 0
	for _, card := range cards {
		res += card.Points()
	}
	return res
}

func part2(cards []*Card) int {
	res := 0
	for i, card := range cards {
		winCount := card.CountWinningNumbers()
		for j := range winCount {
			cards[i+j+1].copies += card.copies
		}
		res += card.copies
	}
	return res
}
