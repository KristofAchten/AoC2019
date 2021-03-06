package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var donutMaze [][]string
var donutPairs map[string][]mazeStep
var coordToID map[coords]string

type stepsPair struct {
	c coords
	s int
}

type stepsPair3D struct {
	c coord3D
	s int
}

type mazeStep struct {
	c coords
	i int
}

func day20() {
	start := time.Now()
	donutPairs = make(map[string][]mazeStep)
	coordToID = make(map[coords]string)

	input := string(getPuzzleInput("/input/day20.txt"))
	parts := strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")
	parseInput(parts)
	fmt.Println("This is what the maze looks like: ")
	printDonutMaze()

	printPuzzleResult(20, searchRoute(), searchRouteWithDepth())

	fmt.Printf("DAY 20 STATS: Execution took %s\n\n", time.Since(start))
}

func printDonutMaze() {
	for _, v := range donutMaze {
		fmt.Println(v)
	}
}

func searchRoute() int {
	var queue []stepsPair
	goal := donutPairs["ZZ"][0]
	queue = append(queue, stepsPair{donutPairs["AA"][0].c, 0})
	visitedInDonut := make(map[coords]bool)

	for len(queue) > 0 {
		val := queue[0]
		visitedInDonut[val.c] = true
		queue = queue[1:]

		if val.c == goal.c {
			return val.s
		}

		id, ok := coordToID[val.c]
		if ok && id != "AA" {
			var tp coords
			for _, c := range donutPairs[id] {
				if c.c != val.c {
					tp = c.c
				}
			}
			queue = append(queue, stepsPair{tp, val.s + 1})
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
			queue = append(queue, stepsPair{v, val.s + 1})
		}
	}
	panic("Couldn't find any route to ZZ")
}

func searchRouteWithDepth() int {
	var queue []stepsPair3D
	goal := coord3D{donutPairs["ZZ"][0].c.x, donutPairs["ZZ"][0].c.y, 0}
	queue = append(queue, stepsPair3D{coord3D{donutPairs["AA"][0].c.x, donutPairs["AA"][0].c.y, 0}, 0})
	visitedInDonut := make(map[coord3D]bool)

	for len(queue) > 0 {
		val := queue[0]
		queue = queue[1:]

		if val.c == goal {
			return val.s
		}

		coords2D := coords{val.c.x, val.c.y}
		id, ok := coordToID[coords2D]
		if ok && id != "AA" && id != "ZZ" && !visitedInDonut[val.c] {
			var tp mazeStep
			var dx int
			for _, c := range donutPairs[id] {
				if c.c != coords2D {
					tp = c
				} else {
					dx = c.i
				}
			}

			if val.c.z != 0 || dx != -1 {
				newCoord3d := coord3D{tp.c.x, tp.c.y, val.c.z + dx}
				queue = append(queue, stepsPair3D{newCoord3d, val.s + 1})
			}
		}
		visitedInDonut[val.c] = true

		neighbours := []coord3D{
			{val.c.x - 1, val.c.y, val.c.z},
			{val.c.x + 1, val.c.y, val.c.z},
			{val.c.x, val.c.y - 1, val.c.z},
			{val.c.x, val.c.y + 1, val.c.z},
		}

		for _, v := range neighbours {
			if donutMaze[v.y][v.x] != "." || visitedInDonut[v] {
				continue
			}
			queue = append(queue, stepsPair3D{v, val.s + 1})
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
					var i int
					if y < 5 {
						i = -1
					} else {
						i = 1
					}
					donutPairs[strVal] = append(donutPairs[strVal], mazeStep{coords{x, y}, i})
					coordToID[coords{x, y}] = strVal
				}
				if ok, _ := regexp.MatchString("[A-Z]", string(down)); ok {
					down2 := input[y+2][x]
					strVal := string(down) + string(down2)
					var i int
					if y > height-5 {
						i = -1
					} else {
						i = 1
					}
					donutPairs[strVal] = append(donutPairs[strVal], mazeStep{coords{x, y}, i})
					coordToID[coords{x, y}] = strVal
				}
				if ok, _ := regexp.MatchString("[A-Z]", string(left)); ok {
					left2 := input[y][x-2]
					strVal := string(left2) + string(left)
					var i int
					if x < 5 {
						i = -1
					} else {
						i = 1
					}
					donutPairs[strVal] = append(donutPairs[strVal], mazeStep{coords{x, y}, i})
					coordToID[coords{x, y}] = strVal
				}
				if ok, _ := regexp.MatchString("[A-Z]", string(right)); ok {
					right2 := input[y][x+2]
					strVal := string(right) + string(right2)
					var i int
					if x > width-5 {
						i = -1
					} else {
						i = 1
					}
					donutPairs[strVal] = append(donutPairs[strVal], mazeStep{coords{x, y}, i})
					coordToID[coords{x, y}] = strVal
				}

			}
		}
	}
}
