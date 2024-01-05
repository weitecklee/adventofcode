package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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

func part1(input string) int {
	i := 0
	for {
		s := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(s))
		hashStr := hex.EncodeToString(hash[:])
		if hashStr[:5] == "00000" {
			return i
		}
		i++
	}
	return -1
}

func part2(input string) int {
	i := 0
	for {
		s := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(s))
		hashStr := hex.EncodeToString(hash[:])
		if hashStr[:6] == "000000" {
			return i
		}
		i++
	}
	return -1
}
