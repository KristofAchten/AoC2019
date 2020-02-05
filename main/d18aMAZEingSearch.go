package main

import (
	"fmt"
	"strings"
	"time"
)

func day18() {
	start := time.Now()

	input := string(getPuzzleInput("/input/day18.txt"))
	parts := strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")
	s, m, doors, keys := parseMaze(parts)

	fmt.Println(s)
	fmt.Println(m)
	fmt.Println(doors)
	fmt.Println(keys)

	fmt.Printf("DAY 9 STATS: Execution took %s\n\n", time.Since(start))
}

func parseMaze(parts []string) (coords, map[coords]bool, map[coords]int, map[coords]int) {

	doormaze := make(map[coords]bool)
	doors := make(map[coords]int)
	keys := make(map[coords]int)

	var y int
	var start coords

	for _, l := range parts {
		var x int
		for _, c := range l {
			switch c {
			case '@':
				start = coords{x, y}
			case '.':
				doormaze[coords{x, y}] = true
			case '#':
				doormaze[coords{x, y}] = false
			default:
				doormaze[coords{x, y}] = true
				if c < 'a' {
					doors[coords{x, y}] = int(c - 'A')
				} else {
					keys[coords{x, y}] = int(c - 'a')
				}
			}
			x++
		}
		y++
	}

	return start, doormaze, doors, keys
}
