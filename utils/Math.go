package utils

import (
	"math"
	"regexp"
	"strconv"
)

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MaxInt(nums ...int) int {
	res := math.MinInt
	for _, n := range nums {
		if n > res {
			res = n
		}
	}
	return res
}

func MinInt(nums ...int) int {
	res := math.MaxInt
	for _, n := range nums {
		if n < res {
			res = n
		}
	}
	return res
}

func PowInt(base, exp int) int {
	res := 1
	for exp > 0 {
		if exp%2 == 1 {
			res *= base
		}
		base *= base
		exp /= 2
	}
	return res
}

var numRegex = regexp.MustCompile(`\d+`)

func ExtractInts(s string) []int {
	matches := numRegex.FindAllString(s, -1)
	nums := make([]int, len(matches))
	for i, match := range matches {
		n, err := strconv.Atoi(match)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

var signedRegex = regexp.MustCompile(`-?\d+`)

func ExtractIntsSigned(s string) []int {
	matches := signedRegex.FindAllString(s, -1)
	nums := make([]int, len(matches))
	for i, match := range matches {
		n, err := strconv.Atoi(match)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}
