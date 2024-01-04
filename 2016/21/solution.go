package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	instructions := parseInput(input)
	fmt.Println(part1(instructions))
	fmt.Println(part2(instructions))
}

type Instruction struct {
	category string
	dir      string
	letters  []string
	numbers  []int
}

func (instruc *Instruction) execute(password string) string {
	var password2 strings.Builder
	n := len(password)
	switch instruc.category {
	case "swap":
		if instruc.dir == "letter" {
			for _, c := range password {
				if string(c) == instruc.letters[0] {
					password2.WriteString(instruc.letters[1])
				} else if string(c) == instruc.letters[1] {
					password2.WriteString(instruc.letters[0])
				} else {
					password2.WriteRune(c)
				}
			}
		} else {
			for i, c := range password {
				if i == instruc.numbers[0] {
					password2.WriteByte(password[instruc.numbers[1]])
				} else if i == instruc.numbers[1] {
					password2.WriteByte(password[instruc.numbers[0]])
				} else {
					password2.WriteRune(c)
				}
			}
		}
	case "rotate":
		k := 0
		if instruc.dir == "based" {
			letter := instruc.letters[0]
			for i, c := range password {
				if string(c) == letter {
					k = i
					break
				}
			}
			if k >= 4 {
				k++
			}
			k++
			k %= n
			k = n - k
		} else {
			k = instruc.numbers[0]
			if instruc.dir == "right" {
				k = n - k
			}
		}
		for i := k; i < n; i++ {
			password2.WriteByte(password[i])
		}
		for i := 0; i < k; i++ {
			password2.WriteByte(password[i])
		}
	case "reverse":
		for i := 0; i < instruc.numbers[0]; i++ {
			password2.WriteByte(password[i])
		}
		for i := instruc.numbers[1]; i >= instruc.numbers[0]; i-- {
			password2.WriteByte(password[i])
		}
		for i := instruc.numbers[1] + 1; i < n; i++ {
			password2.WriteByte(password[i])
		}
	case "move":
		letter := password[instruc.numbers[0]]
		for i, c := range password {
			if i == instruc.numbers[0] {
				continue
			}
			if i == instruc.numbers[1] {
				if instruc.numbers[0] > instruc.numbers[1] {
					password2.WriteByte(letter)
					password2.WriteRune(c)
				} else {
					password2.WriteRune(c)
					password2.WriteByte(letter)
				}
			} else {
				password2.WriteRune(c)
			}
		}
	}
	return password2.String()
}

func parseInput(input []string) *[]Instruction {
	instructions := []Instruction{}
	reLetter := regexp.MustCompile(`\b\w\b`)
	reNumber := regexp.MustCompile(`\d`)
	for _, line := range input {
		tmp := Instruction{}
		parts := strings.Split(line, " ")
		tmp.category = parts[0]
		tmp.dir = parts[1]
		tmp.letters = reLetter.FindAllString(line, -1)
		nums := reNumber.FindAllString(line, -1)
		if nums != nil {
			numbers := []int{}
			for _, n := range nums {
				num, _ := strconv.Atoi(n)
				numbers = append(numbers, num)
			}
			tmp.numbers = numbers
		}
		instructions = append(instructions, tmp)
	}
	return &instructions
}

func scramble(password string, instructions *[]Instruction) string {
	for _, step := range *instructions {
		password = step.execute(password)
	}
	return password
}

func part1(instructions *[]Instruction) string {
	password := "abcdefgh"
	return scramble(password, instructions)
}

// Function to generate permutations of a string based on Heap's algorithm,
// provided by ChatGPT.

func generatePermutations(s []byte) [][]byte {
	permutations := [][]byte{}
	heapPermute(len(s), s, &permutations)
	return permutations
}

func heapPermute(size int, s []byte, permutations *[][]byte) {
	if size == 1 {
		perm := make([]byte, len(s))
		copy(perm, s)
		*permutations = append(*permutations, perm)
		return
	}
	for i := 0; i < size; i++ {
		heapPermute(size-1, s, permutations)
		if size%2 == 1 {
			s[0], s[size-1] = s[size-1], s[0]
		} else {
			s[i], s[size-1] = s[size-1], s[i]
		}
	}
}

func part2(instructions *[]Instruction) string {
	desired := "fbgdceah"
	s := []byte("abcdefgh")
	permutations := generatePermutations(s)
	for _, p := range permutations {
		if scramble(string(p), instructions) == desired {
			return string(p)
		}
	}
	return ""
}
