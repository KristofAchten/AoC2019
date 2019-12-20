package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var donutMaze [][]string
var donutPairs map[string][]coords
var visitedInDonut map[coords]bool
var coordToID map[coords]string

type depthPair struct {
	c coords
	d int
}

func day20() {
	start := time.Now()
	donutPairs = make(map[string][]coords)
	coordToID = make(map[coords]string)
	visitedInDonut = make(map[coords]bool)

	input := string(getPuzzleInput("/input/day20.txt"))
	parts := strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")
	parseInput(parts)
	fmt.Println("This is what the maze looks like: ")
	printDonutMaze()

	printPuzzleResult(20, searchRoute(), notImplemented)

	fmt.Printf("DAY 20 STATS: Execution took %s\n\n", time.Since(start))
}

func printDonutMaze() {
	for _, v := range donutMaze {
		fmt.Println(v)
	}
}

func searchRoute() int {
	var queue []depthPair
	goal := donutPairs["ZZ"][0]
	queue = append(queue, depthPair{donutPairs["AA"][0], 0})

	for len(queue) > 0 {
		val := queue[0]
		visitedInDonut[val.c] = true
		queue = queue[1:]

		if val.c == goal {
			return val.d
		}

		id, ok := coordToID[val.c]
		if ok && id != "AA" {
			var tp coords
			for _, c := range donutPairs[id] {
				if c != val.c {
					tp = c
				}
			}
			if visited[tp] {
				continue
			}
			queue = append(queue, depthPair{tp, val.d + 1})
		}

		neighbours := []coords{
			{val.c.x - 1, val.c.y},
			{val.c.x + 1, val.c.y},
			{val.c.x, val.c.y - 1},
			{val.c.x, val.c.y + 1},
		}

		for _, v := range neighbours {
			if donutMaze[v.y][v.x] != "." || visitedInDonut[v] {
				continue
			}
			queue = append(queue, depthPair{v, val.d + 1})
		}
	}
	panic("Couldn't find any route to ZZ")
}

func parseInput(input []string) {
	width := len(input[2]) + 2
	height := len(input)

	donutMaze = make([][]string, height)
	for i := range donutMaze {
		donutMaze[i] = make([]string, width)
		for j := range donutMaze[i] {
			donutMaze[i][j] = " "
		}
	}

	for y := range input {
		for x := range input[y] {
			curVal := input[y][x]
			donutMaze[y][x] = string(curVal)

			if string(curVal) == "." {
				up := input[y-1][x]
				down := input[y+1][x]
				left := input[y][x-1]
				right := input[y][x+1]

				if ok, _ := regexp.MatchString("[A-Z]", string(up)); ok {
					up2 := input[y-2][x]
					strVal := string(up2) + string(up)
					donutPairs[strVal] = append(donutPairs[strVal], coords{x, y})
					coordToID[coords{x, y}] = strVal
				}
				if ok, _ := regexp.MatchString("[A-Z]", string(down)); ok {
					down2 := input[y+2][x]
					strVal := string(down) + string(down2)
					donutPairs[strVal] = append(donutPairs[strVal], coords{x, y})
					coordToID[coords{x, y}] = strVal
				}
				if ok, _ := regexp.MatchString("[A-Z]", string(left)); ok {
					left2 := input[y][x-2]
					strVal := string(left2) + string(left)
					donutPairs[strVal] = append(donutPairs[strVal], coords{x, y})
					coordToID[coords{x, y}] = strVal
				}
				if ok, _ := regexp.MatchString("[A-Z]", string(right)); ok {
					right2 := input[y][x+2]
					strVal := string(right) + string(right2)
					donutPairs[strVal] = append(donutPairs[strVal], coords{x, y})
					coordToID[coords{x, y}] = strVal
				}

			}
		}
	}
}
