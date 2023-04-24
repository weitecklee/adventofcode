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

func part1(steps []string) string {
	password := "abcdefgh"
	n := len(password)
	for _, step := range steps {
		parts := strings.Split(step, " ")
		var password2 strings.Builder
		switch parts[0] {
		case "swap":
			if parts[1] == "letter" {
				for _, c := range password {
					if string(c) == parts[2] {
						password2.WriteString(parts[5])
					} else if string(c) == parts[5] {
						password2.WriteString(parts[2])
					} else {
						password2.WriteRune(c)
					}
				}
			} else {
				pos1, _ := strconv.Atoi(parts[2])
				pos2, _ := strconv.Atoi(parts[5])
				for i, c := range password {
					if i == pos1 {
						password2.WriteByte(password[pos2])
					} else if i == pos2 {
						password2.WriteByte(password[pos1])
					} else {
						password2.WriteRune(c)
					}
				}
			}
		case "rotate":
			k := 0
			if parts[1] == "based" {
				letter := parts[6]
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
				k, _ = strconv.Atoi(parts[2])
				if parts[1] == "right" {
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
			pos1, _ := strconv.Atoi(parts[2])
			pos2, _ := strconv.Atoi(parts[4])
			for i := 0; i < pos1; i++ {
				password2.WriteByte(password[i])
			}
			for i := pos2; i >= pos1; i-- {
				password2.WriteByte(password[i])
			}
			for i := pos2 + 1; i < n; i++ {
				password2.WriteByte(password[i])
			}
		case "move":
			pos1, _ := strconv.Atoi(parts[2])
			pos2, _ := strconv.Atoi(parts[5])
			letter := password[pos1]
			for i, c := range password {
				if i == pos1 {
					continue
				}
				if i == pos2 {
					if pos1 > pos2 {
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
		password = password2.String()
	}
	return password
}
