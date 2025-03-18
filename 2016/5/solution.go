package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
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
	puzzleInput := string(data)
	fmt.Println(part1(puzzleInput))
	fmt.Println(part2(puzzleInput))
}

func generateMD5Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func part1(puzzleInput string) string {
	i := 0
	fiveZeros := "00000"
	var sb strings.Builder
	for sb.Len() < 8 {
		stringToHash := puzzleInput + strconv.Itoa(i)
		hashRes := generateMD5Hash(stringToHash)
		if hashRes[:5] == fiveZeros {
			sb.WriteByte(hashRes[5])
		}
		i++
	}
	return sb.String()
}

func part2(puzzleInput string) string {
	i := 0
	fiveZeros := "00000"
	resMap := make(map[int]byte, 8)
	for len(resMap) < 8 {
		stringToHash := puzzleInput + strconv.Itoa(i)
		hashRes := generateMD5Hash(stringToHash)
		if hashRes[:5] == fiveZeros {
			pos := int(hashRes[5] - '0')
			if pos < 8 {
				if _, ok := resMap[pos]; !ok {
					resMap[pos] = hashRes[6]
				}
			}
		}
		i++
	}
	var sb strings.Builder
	for i := range 8 {
		sb.WriteByte(resMap[i])
	}
	return sb.String()
}
