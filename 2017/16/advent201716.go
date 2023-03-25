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

func dance(dancers *[]string, step string) {
	if step[0] == 's' {
		spin, _ := strconv.Atoi(step[1:])
		(*dancers) = append((*dancers)[len(*dancers)-spin:], (*dancers)[:len(*dancers)-spin]...)
	} else if step[0] == 'x' {
		swaps := strings.Split(step[1:], "/")
		swap1, _ := strconv.Atoi(swaps[0])
		swap2, _ := strconv.Atoi(swaps[1])
		(*dancers)[swap1], (*dancers)[swap2] = (*dancers)[swap2], (*dancers)[swap1]
	} else {
		swaps := strings.Split(step[1:], "/")
		swap1 := -1
		swap2 := -1
		i := 0
		for swap1 < 0 || swap2 < 0 {
			if (*dancers)[i] == swaps[0] {
				swap1 = i
			} else if (*dancers)[i] == swaps[1] {
				swap2 = i
			}
			i++
		}
		(*dancers)[swap1], (*dancers)[swap2] = (*dancers)[swap2], (*dancers)[swap1]
	}
}

func part1(input []string) string {
	n := 16
	dancers := []string{}
	for i := 0; i < n; i++ {
		dancers = append(dancers, string(i+97))
	}
	for _, step := range input {
		dance(&dancers, step)
	}
	return strings.Join(dancers, "")
}

func part2(input []string) string {
	n := 16
	dancers := []string{}
	for i := 0; i < n; i++ {
		dancers = append(dancers, string(i+97))
	}
	history := map[string]int{}
	i := 0
	j := 0
	danceStr := strings.Join(dancers, "")
	for {
		history[strconv.Itoa(j)+danceStr] = i
		dance(&dancers, input[j])
		i++
		j++
		if j > len(input)-1 {
			j = 0
		}
		danceStr = strings.Join(dancers, "")
		if _, ok := history[strconv.Itoa(j)+danceStr]; ok {
			break
		}
	}
	period := i - history[strconv.Itoa(j)+danceStr]
	leftover := 1000000000 % period
	for k := 0; k < leftover; k++ {
		dance(&dancers, input[(j+k)%len(input)])
	}
	return strings.Join(dancers, "")
}
