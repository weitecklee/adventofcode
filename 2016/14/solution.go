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

func hashChannel(salt string, hashChan chan string) {
	i := 0
	for {
		s := salt + strconv.Itoa(i)
		hash := md5.Sum([]byte(s))
		hashChan <- hex.EncodeToString(hash[:])
		i++
	}
}

type Triplet struct {
	index int
	char  string
}

func findTriplet(s string) string {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return string(s[i])
		}
	}
	return ""
}

func findQuintuplets(s string) []string {
	res := []string{}
	for i := 0; i < len(s)-4; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] && s[i] == s[i+3] && s[i] == s[i+4] {
			res = append(res, string(s[i]))
		}
	}
	return res
}

func part1(salt string) int {
	hashChan := make(chan string, 100000)
	triplets := []Triplet{}
	quintuplets := map[string]map[int]bool{}
	go hashChannel(salt, hashChan)
	for i := 0; i < 100000; i++ {
		hash := <-hashChan
		triplet := findTriplet(hash)
		if len(triplet) > 0 {
			triplets = append(triplets, Triplet{
				index: i,
				char:  triplet,
			})
		}
		quints := findQuintuplets(hash)
		for _, quint := range quints {
			if _, ok := quintuplets[quint]; !ok {
				quintuplets[quint] = map[int]bool{}
			}
			quintuplets[quint][i] = true
		}
	}
	keys := 0
	i := 0
	res := 0
	for keys < 64 {
		check := triplets[i]
		for j := range quintuplets[check.char] {
			if j > check.index && j <= check.index+1000 {
				keys++
				res = check.index
				break
			}
		}
		i++
	}
	return res
}

func hash2016Channel(salt string, hashChan chan string) {
	i := 0
	for {
		s := salt + strconv.Itoa(i)
		hash := md5.Sum([]byte(s))
		for j := 0; j < 2016; j++ {
			hash = md5.Sum([]byte(hex.EncodeToString(hash[:])))
		}
		hashChan <- hex.EncodeToString(hash[:])
		i++
	}
}

func part2(salt string) int {
	hashChan := make(chan string, 100000)
	triplets := []Triplet{}
	quintuplets := map[string]map[int]bool{}
	go hash2016Channel(salt, hashChan)
	for i := 0; i < 100000; i++ {
		hash := <-hashChan
		triplet := findTriplet(hash)
		if len(triplet) > 0 {
			triplets = append(triplets, Triplet{
				index: i,
				char:  triplet,
			})
		}
		quints := findQuintuplets(hash)
		for _, quint := range quints {
			if _, ok := quintuplets[quint]; !ok {
				quintuplets[quint] = map[int]bool{}
			}
			quintuplets[quint][i] = true
		}
	}
	keys := 0
	i := 0
	res := 0
	for keys < 64 {
		check := triplets[i]
		for j := range quintuplets[check.char] {
			if j > check.index && j <= check.index+1000 {
				keys++
				res = check.index
				break
			}
		}
		i++
	}
	return res
}
