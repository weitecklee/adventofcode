package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(string(data)))
	fmt.Println(part1v2(string(data)))
}

/*
	Comparison between using adding to string and using strings.Builder.
	First method takes 7 seconds on average.
	Second method takes 7 milliseconds on average.
*/

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

func lookAndSay(input string) string {
	var curr byte
	i := 0
	res := ""
	for i < len(input) {
		curr = input[i]
		n := 0
		for i < len(input) && input[i] == curr {
			n++
			i++
		}
		res += strconv.Itoa(n) + string(curr)
	}
	return res
}

func lookAndSay2(input string) string {
	var curr byte
	i := 0
	var res strings.Builder
	for i < len(input) {
		curr = input[i]
		n := 0
		for i < len(input) && input[i] == curr {
			n++
			i++
		}
		res.WriteString(strconv.Itoa(n))
		res.WriteByte(curr)
	}
	return res.String()
}

func part1(input string) int {
	defer duration(track("part1"))
	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}
	return len(input)
}

func part1v2(input string) int {
	defer duration(track("part1v2"))
	for i := 0; i < 40; i++ {
		input = lookAndSay2(input)
	}
	return len(input)
}
