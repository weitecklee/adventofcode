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
	input := string(data)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func dragon(a string) string {
	var b strings.Builder
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '0' {
			b.WriteString("1")
		} else {
			b.WriteString("0")
		}
	}
	return a + "0" + b.String()
}

func checksum(s string, l int) string {
	for len(s) < l {
		s = dragon(s)
	}
	s = s[:l]
	for len(s)%2 == 0 {
		var check strings.Builder
		for i := 0; i < len(s); i += 2 {
			if s[i] == s[i+1] {
				check.WriteString("1")
			} else {
				check.WriteString("0")
			}
		}
		s = check.String()
	}
	return s
}

func part1(s string) string {
	return checksum(s, 272)
}

func part2(s string) string {
	return checksum(s, 35651584)
}
