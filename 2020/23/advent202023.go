package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parsedInput := parseInput(string(data))
	fmt.Println(part1(parsedInput))
}

func parseInput(input string) []int {
	arr := []int{}
	for _, c := range input {
		n, _ := strconv.Atoi(string(c))
		arr = append(arr, n)
	}
	return arr
}

func contains(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}
	return false
}

func move(cups [10][2]int, curr int) ([10][2]int, int) {
	next3 := []int{}
	tmp := curr
	for i := 0; i < 3; i++ {
		next3 = append(next3, cups[tmp][1])
		tmp = cups[tmp][1]
	}
	dest := curr - 1
	if dest == 0 {
		dest = 9
	}
	for contains(next3, dest) {
		dest--
		if dest == 0 {
			dest = 9
		}
	}
	cups[curr][1] = cups[next3[2]][1]
	oldNext := cups[dest][1]
	cups[dest][1] = next3[0]
	cups[next3[2]][1] = oldNext
	return cups, cups[curr][1]
}

func part1(input []int) string {
	cups := [10][2]int{} // doubly linked list approach, cups[i] = [previous, next] for cup i
	cups[input[0]][0] = input[8]
	cups[input[0]][1] = input[1]
	cups[input[8]][0] = input[7]
	cups[input[8]][1] = input[0]
	for i := 1; i < len(input)-1; i++ {
		cups[input[i]][0] = input[i-1]
		cups[input[i]][1] = input[i+1]
	}
	curr := input[0]
	for i := 0; i < 100; i++ {
		cups, curr = move(cups, curr)
	}
	ret := ""
	tmp := 1
	for cups[tmp][1] != 1 {
		ret += strconv.Itoa(cups[tmp][1])
		tmp = cups[tmp][1]
	}
	return ret
}
