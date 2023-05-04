package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(string(data)))
}

func part1(password string) string {
	pw := []rune(password)
	valid := false
	for !valid {
		i := len(pw) - 1
		for i >= 0 {
			pw[i]++
			if pw[i] == 105 || pw[i] == 108 || pw[i] == 111 {
				pw[i]++
			}
			if pw[i] > 122 {
				pw[i] = 97
				i--
			} else {
				break
			}
		}
		for j := 0; j < len(pw)-2; j++ {
			if pw[j] == pw[j+1]-1 && pw[j+1] == pw[j+2]-1 {
				valid = true
				break
			}
		}
		if !valid {
			continue
		}
		valid = false
		for j := 0; j < len(pw)-1; j++ {
			if pw[j] == pw[j+1] {
				for k := j + 2; k < len(pw)-1; k++ {
					if pw[k] == pw[k+1] && pw[j] != pw[k] {
						valid = true
						break
					}
				}
				break
			}
		}
	}
	var nextPassword strings.Builder
	for _, r := range pw {
		nextPassword.WriteRune(r)
	}
	return nextPassword.String()
}
