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
	fmt.Println(part1(input))
}

func parseInput(input []string) (int, int) {
	cardPublicKey, _ := strconv.Atoi(input[0])
	doorPublicKey, _ := strconv.Atoi(input[1])
	return cardPublicKey, doorPublicKey
}

func part1(input []string) int {
	cardPublicKey, _ := strconv.Atoi(input[0])
	doorPublicKey, _ := strconv.Atoi(input[1])
	subjectNumber := 7
	loops := 0
	value := 1
	for value != cardPublicKey {
		loops++
		value *= subjectNumber
		value %= 20201227
	}
	cardLoops := loops
	// loops = 0
	// value = 1
	// for value != doorPublicKey {
	// 	loops++
	// 	value *= subjectNumber
	// 	value %= 20201227
	// }
	// doorLoops := loops
	value = 1
	for i := 0; i < cardLoops; i++ {
		value *= doorPublicKey
		value %= 20201227
	}
	// value = 1
	// for i := 0; i < doorLoops; i++ {
	// 	value *= cardPublicKey
	// 	value %= 20201227
	// }
	return value
}
