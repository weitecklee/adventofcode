package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	root := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(part1(root))
	fmt.Println(part2(root))
}

type Directory struct {
	name        string
	location    *Directory
	directories map[string]*Directory
	files       map[string]*File
	size        int
}

func NewDirectory(name string, location *Directory) *Directory {
	return &Directory{name, location, make(map[string]*Directory), make(map[string]*File), 0}
}

func (d *Directory) Size() int {
	if d.size == 0 {
		for _, file := range d.files {
			d.size += file.size
		}
		for _, directory := range d.directories {
			d.size += directory.Size()
		}
	}
	return d.size
}

type File struct {
	name     string
	size     int
	location *Directory
}

func parseInput(data []string) *Directory {
	root := NewDirectory("/", nil)
	curr := root
	for _, line := range data[1:] {
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == ".." {
					curr = curr.location
				} else {
					curr = curr.directories[parts[2]]
				}
			}
		} else if parts[0] == "dir" {
			curr.directories[parts[1]] = NewDirectory(parts[1], curr)
		} else {
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			curr.files[parts[1]] = &File{parts[1], size, curr}
		}
	}
	return root
}

func part1(root *Directory) int {
	queue := []*Directory{root}
	res := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Size() <= 100000 {
			res += curr.Size()
		}
		for _, directory := range curr.directories {
			queue = append(queue, directory)
		}
	}
	return res
}

func part2(root *Directory) int {
	totalSpace := 70000000
	neededSpace := 30000000
	usedSpace := root.Size()
	freeSpace := totalSpace - usedSpace
	res := totalSpace
	queue := []*Directory{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Size()+freeSpace >= neededSpace && curr.Size() < res {
			res = curr.Size()
		}
		for _, directory := range curr.directories {
			queue = append(queue, directory)
		}
	}
	return res
}
