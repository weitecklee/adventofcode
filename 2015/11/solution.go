package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
	part1 := nextPassword(string(data))
	part2 := nextPassword(part1)
	fmt.Println(part1, part2)
}

func nextPassword(password string) string {
	pw := []rune(password)
	for {
		incrementPassword(&pw)
		if validatePassword(&pw) {
			break
		}
	}
	var nextPassword strings.Builder
	for _, r := range pw {
		nextPassword.WriteRune(r)
	}
	return nextPassword.String()
}

func incrementPassword(pw *[]rune) {
	i := len(*pw) - 1
	for i >= 0 {
		(*pw)[i]++
		if (*pw)[i] == 105 || (*pw)[i] == 108 || (*pw)[i] == 111 {
			(*pw)[i]++
		}
		if (*pw)[i] > 122 {
			(*pw)[i] = 97
			i--
		} else {
			break
		}
	}
}

func validatePassword(pw *[]rune) bool {
	valid := false
	for j := 0; j < len(*pw)-2; j++ {
		if (*pw)[j] == (*pw)[j+1]-1 && (*pw)[j+1] == (*pw)[j+2]-1 {
			valid = true
			break
		}
	}
	if !valid {
		return false
	}
	valid = false
	for j := 0; j < len(*pw)-1; j++ {
		if (*pw)[j] == (*pw)[j+1] {
			for k := j + 2; k < len((*pw))-1; k++ {
				if (*pw)[k] == (*pw)[k+1] && (*pw)[j] != (*pw)[k] {
					valid = true
					break
				}
			}
			break
		}
	}
	return valid
}
