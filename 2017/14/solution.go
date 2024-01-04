package main

import (
	"container/ring"
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
	count, hash := part1(data)
	fmt.Println(count)
	fmt.Println(part2(hash))
}

func part1(input []byte) (int, *[][]string) {
	lengths := []int{}
	for _, b := range input {
		lengths = append(lengths, int(b))
	}
	lengths = append(lengths, int('-'))
	listLen := 256
	count := 0
	hash := [][]string{}
	for i := 0; i < 128; i++ {
		str := strconv.Itoa(i)
		lengths2 := lengths[:]
		for _, c := range str {
			lengths2 = append(lengths2, int(c))
		}
		lengths2 = append(lengths2, 17, 31, 73, 47, 23)
		circList := ring.New(listLen)
		circListMap := make(map[int]*ring.Ring, listLen)
		for i := 0; i < listLen; i++ {
			circList.Value = i
			circListMap[i] = circList
			circList = circList.Next()
		}
		pos := 0
		skip := 0
		for round := 0; round < 64; round++ {
			for _, n := range lengths2 {
				arr := []int{}
				for i := 0; i < n; i++ {
					arr = append(arr, circList.Value.(int))
					circList = circList.Next()
				}
				circList = circList.Prev()
				for i := 0; i < n; i++ {
					circList.Value = arr[i]
					circList = circList.Prev()
				}
				pos += n + skip
				skip++
				pos %= listLen
				circList = circListMap[pos]
			}
		}
		knotHash := ""
		for block := 0; block < 16; block++ {
			curr := circListMap[block*16].Value.(int)
			for i := 1; i < 16; i++ {
				curr ^= circListMap[block*16+i].Value.(int)
			}
			knotHash += fmt.Sprintf("%02x", curr)
		}
		knotHash2 := ""
		for _, c := range knotHash {
			n, _ := strconv.ParseInt(string(c), 16, 16)
			knotHash2 += fmt.Sprintf("%04b", n)
		}
		for _, c := range knotHash2 {
			if c == '1' {
				count++
			}
		}
		hash = append(hash, strings.Split(knotHash2, ""))
	}
	return count, &hash
}

func marker(hash *[][]string, i int, j int) {
	if i > 0 && (*hash)[i-1][j] == "1" {
		(*hash)[i-1][j] = "X"
		marker(hash, i-1, j)
	}
	if j > 0 && (*hash)[i][j-1] == "1" {
		(*hash)[i][j-1] = "X"
		marker(hash, i, j-1)
	}
	if i < len(*hash)-1 && (*hash)[i+1][j] == "1" {
		(*hash)[i+1][j] = "X"
		marker(hash, i+1, j)
	}
	if j < len((*hash)[0])-1 && (*hash)[i][j+1] == "1" {
		(*hash)[i][j+1] = "X"
		marker(hash, i, j+1)
	}
}

func part2(hash *[][]string) int {
	count := 0
	for i := range *hash {
		for j := range (*hash)[i] {
			if (*hash)[i][j] == "1" {
				count++
				(*hash)[i][j] = "X"
				marker(hash, i, j)
			}
		}
	}
	return count
}
