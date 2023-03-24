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
	input := strings.Split(string(data), ",")
	fmt.Println(part1(parseInput(input)))
	fmt.Println(part2(data))
}

func parseInput(input []string) []int {
	arr := []int{}
	for _, c := range input {
		n, _ := strconv.Atoi(c)
		arr = append(arr, n)
	}
	return arr
}

func part1(input []int) int {
	listLen := 256
	circList := ring.New(listLen)
	circListMap := make(map[int]*ring.Ring, listLen)
	for i := 0; i < listLen; i++ {
		circList.Value = i
		circListMap[i] = circList
		circList = circList.Next()
	}
	pos := 0
	skip := 0
	for _, n := range input {
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
	return circListMap[0].Value.(int) * circListMap[1].Value.(int)
}

func part2(input []byte) string {
	lengths := []int{}
	for _, b := range input {
		lengths = append(lengths, int(b))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)
	listLen := 256
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
		for _, n := range lengths {
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
	return knotHash
}
