package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input202014.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	memory := map[string]int{}
	var mask string
	re := regexp.MustCompile(`\S*$`)
	re2 := regexp.MustCompile(`\d+`)
	for _, line := range input {
		if strings.Contains(line, "mask") {
			mask = re.FindString(line)
		} else {
			matches := re2.FindAllString(line, -1)
			mem, _ := strconv.Atoi(matches[1])
			memBinary := strconv.FormatInt(int64(mem), 2)
			for len(memBinary) < 36 {
				memBinary = "0" + memBinary
			}
			masked := ""
			for i := range mask {
				if mask[i] == "X"[0] {
					masked += string(memBinary[i])
				} else {
					masked += string(mask[i])
				}
			}
			maskedVal, _ := strconv.ParseInt(masked, 2, 64)
			memory[matches[0]] = int(maskedVal)
		}
	}
	sum := 0
	for _, value := range memory {
		sum += value
	}
	return sum
}

func possibleAddresses(maskedAddress string) []int {
	floats := []int{}
	for i, c := range maskedAddress {
		if string(c) == "X" {
			floats = append(floats, i)
		}
	}
	addresses := []string{maskedAddress}
	for _, floating := range floats {
		addresses2 := []string{}
		for _, address := range addresses {
			addresses2 = append(addresses2, address[:floating]+"0"+address[floating+1:])
			addresses2 = append(addresses2, address[:floating]+"1"+address[floating+1:])
		}
		addresses = addresses2
	}
	addressesInt := []int{}
	for _, address := range addresses {
		addressInt, _ := strconv.ParseInt(address, 2, 64)
		addressesInt = append(addressesInt, int(addressInt))
	}
	return addressesInt
}

func part2(input []string) int {
	memory := map[int]int{}
	var mask string
	re := regexp.MustCompile(`\S*$`)
	re2 := regexp.MustCompile(`\d+`)
	for _, line := range input {
		if strings.Contains(line, "mask") {
			mask = re.FindString(line)
		} else {
			matches := re2.FindAllString(line, -1)
			address, _ := strconv.Atoi(matches[0])
			addressBinary := strconv.FormatInt(int64(address), 2)
			for len(addressBinary) < 36 {
				addressBinary = "0" + addressBinary
			}
			maskedAddress := ""
			for i := range mask {
				if mask[i] == "0"[0] {
					maskedAddress += string(addressBinary[i])
				} else {
					maskedAddress += string(mask[i])
				}
			}
			possible := possibleAddresses(maskedAddress)
			val, _ := strconv.Atoi(matches[1])
			for _, addr := range possible {
				memory[addr] = val
			}
		}
	}
	sum := 0
	for _, value := range memory {
		sum += value
	}
	return sum
}
