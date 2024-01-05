package main

import (
	"container/ring"
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
	input, _ := strconv.Atoi(string(data))
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input int) int {
	spinlock := ring.New(1)
	spinlock.Value = 0
	for i := 1; i <= 2017; i++ {
		tmp := ring.New(1)
		tmp.Value = i
		spins := input % i
		for j := 0; j < spins; j++ {
			spinlock = spinlock.Next()
		}
		spinlock.Link(tmp)
		spinlock = tmp
	}
	return spinlock.Next().Value.(int)
}

func part2(input int) int {
	// only need to keep track of insert position, record whenever it is right after zero (pos = 1)
	pos := 0
	res := 0
	for i := 1; i <= 50000000; i++ {
		pos = (pos+input)%i + 1
		if pos == 1 {
			res = i
		}
	}
	return res
}
