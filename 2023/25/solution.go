package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
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
	componentMap := parseInput(strings.Split(string(data), "\n"))
	fmt.Println(solve(componentMap))
}

/*
  Strategy is to take 1000 random pairs of nodes and find
  the shortest path between them. We tally up all the
  node-to-node connections in each path and find the three
  most common connections. Presumably, these three are the
  most critical connections that we want to cut.
  Seems to work well enough, 1000 is a big enough number that
  the same three connections are found each time.
  For the record, my input had 1458 nodes, which makes
  1458 * 1457 / 2 = 1062153 pairs.
*/

type Component struct {
	name      string
	neighbors map[*Component]struct{}
}

func addLink(c1, c2 *Component) {
	c1.neighbors[c2] = struct{}{}
	c2.neighbors[c1] = struct{}{}
}

func removeLink(c1, c2 *Component) {
	delete(c1.neighbors, c2)
	delete(c2.neighbors, c1)
}

func parseInput(data []string) map[string]*Component {
	componentMap := make(map[string]*Component)
	componentRegex := regexp.MustCompile(`[a-z]+`)
	for _, line := range data {
		matches := componentRegex.FindAllString(line, -1)
		components := make([]*Component, len(matches))
		for i, match := range matches {
			name := match
			if _, ok := componentMap[name]; !ok {
				componentMap[name] = &Component{name, map[*Component]struct{}{}}
			}
			components[i] = componentMap[name]
		}
		for _, component := range components[1:] {
			addLink(components[0], component)
		}
	}
	return componentMap
}

func pickRandom[T any](sl []T, n int) []T {
	if len(sl) == 0 || n <= 0 {
		return nil
	}
	res := make([]T, 0, n)
	for range n {
		res = append(res, sl[rand.Intn(len(sl))])
	}
	return res
}

type QueueEntry struct {
	component *Component
	path      []*Component
}

func shortestPath(c1, c2 *Component) []*Component {
	if c1 == c2 {
		return []*Component{c1, c2}
	}
	visited := make(map[*Component]struct{})
	visited[c1] = struct{}{}
	queue := []QueueEntry{{c1, []*Component{c1}}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.component == c2 {
			return curr.path
		}
		for neighbor := range curr.component.neighbors {
			if _, ok := visited[neighbor]; ok {
				continue
			}
			visited[neighbor] = struct{}{}
			path2 := slices.Clone(curr.path)
			path2 = append(path2, neighbor)
			queue = append(queue, QueueEntry{neighbor, path2})
		}
	}
	return nil
}

func linksFromPath(path []*Component) [][2]*Component {
	links := make([][2]*Component, len(path)-1)
	for i := range path[1:] {
		link := [2]*Component{path[i], path[i+1]}
		if path[i].name > path[i+1].name {
			link = [2]*Component{path[i+1], path[i]}
		}
		links[i] = link
	}
	return links
}

func solve(componentMap map[string]*Component) int {
	components := make([]*Component, 0, len(componentMap))
	for _, component := range componentMap {
		components = append(components, component)
	}

	pairs := make([][2]*Component, 0, 1000)
	var pair [2]*Component
	for range 1000 {
		pairSlice := pickRandom(components, 2)
		copy(pair[:], pairSlice)
		pairs = append(pairs, pair)
	}

	linkCountMap := make(map[[2]*Component]int)
	for _, pair := range pairs {
		path := shortestPath(pair[0], pair[1])
		links := linksFromPath(path)
		for _, link := range links {
			linkCountMap[link]++
		}
	}

	var top3 [3][2]*Component
	for link, count := range linkCountMap {
		for i, pair := range top3 {
			if count > linkCountMap[pair] {
				copy(top3[i+1:], top3[i:])
				top3[i] = link
				break
			}
		}
	}
	for _, pair := range top3 {
		removeLink(pair[0], pair[1])
	}

	queue := []*Component{components[0]}
	group1 := make(map[*Component]struct{})
	group1[queue[0]] = struct{}{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for neighbor := range curr.neighbors {
			if _, ok := group1[neighbor]; ok {
				continue
			}
			group1[neighbor] = struct{}{}
			queue = append(queue, neighbor)
		}
	}

	return (len(componentMap) - len(group1)) * len(group1)
}
