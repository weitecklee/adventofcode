package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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

func presents(n int) int {
	sum := 1 + n
	k := int(math.Sqrt(float64(n)))
	for i := 2; i < k; i++ {
		if n%i == 0 {
			sum += i
			sum += n / i
		}
	}
	if k*k == n {
		sum += k
	}
	return sum * 10
}

func part1(input int) int {
	i := 1
	n := presents(i)
	for n < input {
		i++
		n = presents(i)
	}
	return i
}

func presents2(n int) int {
	sum := 0
	k := int(math.Sqrt(float64(n)))
	for i := 1; i < k; i++ {
		if n%i == 0 {
			if 50*i >= n {
				sum += i
			}
			if 50*n/i >= n {
				sum += n / i
			}
		}
	}
	if k*k == n && 50*k >= n {
		sum += k
	}
	return sum * 11
}

func part2(input int) int {
	i := 1
	n := presents2(i)
	for n < input {
		i++
		n = presents2(i)
	}
	return i
}
