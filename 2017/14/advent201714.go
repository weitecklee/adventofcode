package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(data))
}

func part1(input []byte) int {
	lengths := []int{}
	for _, b := range input {
		lengths = append(lengths, int(b))
	}
	lengths = append(lengths, int('-'))
	listLen := 256
	count := 0
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
	}
	return count
}
