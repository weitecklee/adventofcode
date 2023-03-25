package main

import (
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
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func dance(dancers string, step string) string {
	dancers2 := ""
	if step[0] == 's' {
		spin, _ := strconv.Atoi(step[1:])
		dancers2 = dancers[(len(dancers)-spin):] + dancers[:(len(dancers)-spin)]
	} else if step[0] == 'x' {
		swaps := strings.Split(step[1:], "/")
		swap1, _ := strconv.Atoi(swaps[0])
		swap2, _ := strconv.Atoi(swaps[1])
		for i := 0; i < len(dancers); i++ {
			if i == swap1 {
				dancers2 += string(dancers[swap2])
			} else if i == swap2 {
				dancers2 += string(dancers[swap1])
			} else {
				dancers2 += string(dancers[i])
			}
		}
	} else {
		swaps := strings.Split(step[1:], "/")
		for i := 0; i < len(dancers); i++ {
			if string(dancers[i]) == swaps[0] {
				dancers2 += swaps[1]
			} else if string(dancers[i]) == swaps[1] {
				dancers2 += swaps[0]
			} else {
				dancers2 += string(dancers[i])
			}
		}
	}
	return dancers2
}

func part1(input []string) string {
	n := 16
	dancers := ""
	for i := 0; i < n; i++ {
		dancers += string(i + 97)
	}
	for _, step := range input {
		dancers = dance(dancers, step)
	}
	return dancers
}

func part2(input []string) string {
	n := 16
	dancers := ""
	for i := 0; i < n; i++ {
		dancers += string(i + 97)
	}
	history := map[string]int{}
	i := 0
	j := 0
	for {
		history[strconv.Itoa(j)+dancers] = i
		dancers = dance(dancers, input[j])
		i++
		j++
		if j > len(input)-1 {
			j = 0
		}
		if _, ok := history[strconv.Itoa(j)+dancers]; ok {
			break
		}
	}
	period := i - history[strconv.Itoa(j)+dancers]
	leftover := 1000000000 % period
	for k := 0; k < leftover; k++ {
		dancers = dance(dancers, input[(j+k)%len(input)])
	}
	return dancers
}
