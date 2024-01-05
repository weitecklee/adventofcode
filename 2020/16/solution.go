package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	inputFilePath := filepath.Join(dirname, "input.txt")
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	valids, yourTicket, nearbyTickets, fields := parseInput(input)
	errorRate, validTickets := part1(valids, nearbyTickets)
	fmt.Println(errorRate)
	fmt.Println(part2(valids, validTickets, yourTicket, fields))
}

func parseInput(input []string) ([1000]map[string]bool, []int, [][]int, []string) {
	valids := [1000]map[string]bool{}
	for i := range valids {
		valids[i] = map[string]bool{}
	}
	i := 0
	re := regexp.MustCompile(`\d+`)
	reField := regexp.MustCompile(`.*:`)
	fields := []string{}
	for len(input[i]) > 0 {
		numsS := re.FindAllString(input[i], -1)
		field := reField.FindString(input[i])
		numsI := [4]int{}
		for j := range numsS {
			numsI[j], _ = strconv.Atoi(numsS[j])
		}
		for j := numsI[0]; j <= numsI[1]; j++ {
			valids[j][field] = true
		}
		for j := numsI[2]; j <= numsI[3]; j++ {
			valids[j][field] = true
		}
		fields = append(fields, field)
		i++
	}
	i += 2
	yourTicket := []int{}
	numsS := re.FindAllString(input[i], -1)
	for _, numS := range numsS {
		numI, _ := strconv.Atoi(numS)
		yourTicket = append(yourTicket, numI)
	}
	i += 3
	nearbyTickets := [][]int{}
	for i < len(input) {
		numsS := re.FindAllString(input[i], -1)
		currRow := []int{}
		for _, numS := range numsS {
			numI, _ := strconv.Atoi(numS)
			currRow = append(currRow, numI)
		}
		nearbyTickets = append(nearbyTickets, currRow)
		i++
	}
	return valids, yourTicket, nearbyTickets, fields
}

func part1(valids [1000]map[string]bool, nearbyTickets [][]int) (int, [][]int) {
	sum := 0
	validTickets := [][]int{}
	for _, row := range nearbyTickets {
		valid := true
		for _, val := range row {
			if len(valids[val]) == 0 {
				sum += val
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, row)
		}
	}
	return sum, validTickets
}

func part2(valids [1000]map[string]bool, validTickets [][]int, yourTicket []int, fields []string) int {
	ticketFields := [20]map[string]bool{}
	for i := range ticketFields {
		ticketFields[i] = map[string]bool{}
		for _, field := range fields {
			ticketFields[i][field] = true
		}
	}
	for _, row := range validTickets {
		for i, val := range row {
			for _, field := range fields {
				if !valids[val][field] {
					ticketFields[i][field] = false
				}
			}
		}
	}
	determinedFields := map[string]int{}
	determinedPositions := map[int]string{}
	for i := 0; i < 20; i++ {
		determinedPositions[i] = ""
	}
	for len(determinedFields) < 20 {
		determinedField := ""
		for i, ticketField := range ticketFields {
			if determinedPositions[i] != "" {
				continue
			}
			possibleFields := 0
			for field, valid := range ticketField {
				if valid {
					determinedField = field
					possibleFields++
				}
			}
			if possibleFields == 1 {
				determinedFields[determinedField] = i
				determinedPositions[i] = determinedField
				break
			}
		}
		for i, ticketField := range ticketFields {
			if determinedPositions[i] != "" {
				continue
			}
			ticketField[determinedField] = false
		}
	}
	prod := 1
	for field, pos := range determinedFields {
		if strings.Contains(field, "departure") {
			prod *= yourTicket[pos]
		}
	}
	return prod
}
