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
	input, _ := strconv.Atoi(string(data))
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func printRing(ring *ring.Ring) {
	v := ring.Value
	var res strings.Builder
	res.WriteString(strconv.Itoa(v.(int)))
	ring = ring.Next()
	for ring.Value != v {
		res.WriteString(strconv.Itoa(ring.Value.(int)))
		ring = ring.Next()
	}
	fmt.Println(res.String())
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
	spinlock := ring.New(1)
	spinlock.Value = 0
	spinzero := spinlock
	for i := 1; i <= 50000000; i++ {
		tmp := ring.New(1)
		tmp.Value = i
		spins := input % i
		for j := 0; j < spins; j++ {
			spinlock = spinlock.Next()
		}
		spinlock.Link(tmp)
		spinlock = tmp
	}
	return spinzero.Next().Value.(int)
}
