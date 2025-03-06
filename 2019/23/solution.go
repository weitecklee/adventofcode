package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/weitecklee/adventofcode/2019/intcode"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	puzzleInput := parseInput(strings.Split(string(data), ","))
	fmt.Println(solve(puzzleInput))
}

func parseInput(data []string) []int {
	numbers := make([]int, 0, len(data))
	for _, s := range data {
		if n, err := strconv.Atoi(s); err != nil {
			panic(err)
		} else {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

type Packet struct {
	dst int
	x   int
	y   int
}

func collectPackets(packets *[]Packet, ch chan int) {
	var ret, x, y int
	for {
		ret = <-ch
		if ret == intcode.REQUESTSIGNAL {
			break
		}
		x = <-ch
		y = <-ch
		*packets = append(*packets, Packet{ret, x, y})
	}
}

func solve(puzzleInput []int) (int, int) {
	network := make(map[int]chan int, 50)
	var packet, natPacket Packet
	var packets []Packet
	var ch chan int
	natHistory := make(map[int]struct{})
	part1 := math.MinInt

	for i := range 50 {
		ch := make(chan int)
		ic := intcode.NewIntcodeProgram(puzzleInput, ch)
		go ic.Run()
		<-ch
		ch <- i
		network[i] = ch
		<-ch
		ch <- -1
		collectPackets(&packets, ch)
	}

	for {
		for len(packets) > 0 {
			packet = packets[0]
			packets = packets[1:]
			if packet.dst == 255 {
				if part1 == math.MinInt {
					part1 = packet.y
				}
				natPacket = packet
				continue
			}
			ch = network[packet.dst]
			ch <- packet.x
			<-ch
			ch <- packet.y
			collectPackets(&packets, ch)
		}

		ch = network[0]
		if _, ok := natHistory[natPacket.y]; ok {
			return part1, natPacket.y
		}
		natHistory[natPacket.y] = struct{}{}
		ch <- natPacket.x
		<-ch
		ch <- natPacket.y
		collectPackets(&packets, ch)
	}
}
