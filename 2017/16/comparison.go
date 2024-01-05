package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), ",")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
	fmt.Println(part1v2(input))
	fmt.Println(part2v2(input))
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
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
	defer duration(track("part1"))
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
	defer duration(track("part2"))
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

func dancev2(dancers *[]string, step string) {
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

func part1v2(input []string) string {
	defer duration(track("part1v2"))
	n := 16
	dancers := []string{}
	for i := 0; i < n; i++ {
		dancers = append(dancers, string(i+97))
	}
	for _, step := range input {
		dancev2(&dancers, step)
	}
	return strings.Join(dancers, "")
}

func part2v2(input []string) string {
	defer duration(track("part2v2"))
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
		dancev2(&dancers, input[j])
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
		dancev2(&dancers, input[(j+k)%len(input)])
	}
	return strings.Join(dancers, "")
}
