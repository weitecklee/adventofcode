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
	puzzleInput := strings.Split(string(data), "\n")
	rooms := createRooms(puzzleInput)
	realRooms := verifyRooms(rooms)
	fmt.Println(part1(realRooms))
	fmt.Println(part2(realRooms))
}

type Room struct {
	name     string
	sectorId int
	checksum string
}

var roomRegex = regexp.MustCompile(`^([a-z\-]+)\-(\d+)\[(\w+)\]$`)

func NewRoom(s string) *Room {
	match := roomRegex.FindStringSubmatch(s)
	if match == nil {
		panic(fmt.Sprintf("Could not match room regex with: %s", s))
	}
	sectorId, err := strconv.Atoi(match[2])
	if err != nil {
		panic(err)
	}
	return &Room{match[1], sectorId, match[3]}
}

func (r *Room) IsReal() bool {
	letterMap := make(map[rune]int, 26)
	for _, ch := range r.name {
		if ch != '-' {
			letterMap[ch]++
		}
	}
	var top5 [5]rune
	for ch, n := range letterMap {
		for i, r := range top5 {
			if n > letterMap[r] || (n == letterMap[r] && ch < r) {
				copy(top5[i+1:], top5[i:])
				top5[i] = ch
				break
			}
		}
	}
	for i, ch := range r.checksum {
		if top5[i] != ch {
			return false
		}
	}
	return true
}

func (r *Room) DecryptedName() string {
	var sb strings.Builder
	for _, ch := range r.name {
		sb.WriteRune(rotateLetter(ch, r.sectorId))
	}
	return sb.String()
}

func createRooms(puzzleInput []string) []*Room {
	rooms := make([]*Room, len(puzzleInput))
	for i, line := range puzzleInput {
		rooms[i] = NewRoom(line)
	}
	return rooms
}

func verifyRooms(rooms []*Room) []*Room {
	var realRooms []*Room
	for _, room := range rooms {
		if room.IsReal() {
			realRooms = append(realRooms, room)
		}
	}
	return realRooms
}

func rotateLetter(letter rune, n int) rune {
	if letter == '-' {
		return ' '
	}
	return (letter-'a'+rune(n))%26 + 'a'
}

func part1(rooms []*Room) int {
	res := 0
	for _, room := range rooms {
		res += room.sectorId
	}
	return res
}

func part2(rooms []*Room) int {
	target := "northpole object storage"
	for _, room := range rooms {
		if room.DecryptedName() == target {
			return room.sectorId
		}
	}
	return -1
}
