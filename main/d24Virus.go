package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

var virusGrid [5][5]bool
var visitedViruses map[[5][5]bool]bool

func day24() {
	start := time.Now()

	visitedViruses = make(map[[5][5]bool]bool)

	input := string(getPuzzleInput("input/day24.txt"))
	parts := strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")

	parseVirusGrid(parts)

	printPuzzleResult(24, getBioDiversityOfFirstRecurrence(), notImplemented)

	fmt.Printf("DAY 24 STATS: Execution took %s\n\n", time.Since(start))
}

func getBioDiversityOfFirstRecurrence() int {
	stopAtfirstRecurrentMatch()

	var sum int
	for i := range virusGrid {
		for j := range virusGrid {
			if virusGrid[i][j] {
				val := (i * 5) + j
				sum += int(math.Pow(2, float64(val)))
			}
		}
	}
	return sum
}

func stopAtfirstRecurrentMatch() {
	for {
		grow()
		if _, ok := visitedViruses[virusGrid]; ok {
			return
		}
		visitedViruses[virusGrid] = true
	}
}

func parseVirusGrid(parts []string) {
	for i := range parts {
		for j := range parts[i] {
			if parts[i][j] == '.' {
				virusGrid[i][j] = false
			} else {
				virusGrid[i][j] = true
			}
		}
	}
	visitedViruses[virusGrid] = true
}

func printVirusGrid() {
	for i := range virusGrid {
		for j := range virusGrid {
			if virusGrid[i][j] == true {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func grow() {
	var newGrid [5][5]bool
	for i := range virusGrid {
		for j := range virusGrid {
			if virusGrid[i][j] {
				newGrid[i][j] = hasExactlyOneNeighbour(i, j)
			} else {
				newGrid[i][j] = hasOneOrTwoNeighbours(i, j)
			}
		}
	}
	virusGrid = newGrid
}

func hasOneOrTwoNeighbours(i int, j int) bool {
	neighbours := getNeighbours(i, j)

	cnt := 0
	for _, v := range neighbours {
		if v {
			cnt++
		}
	}

	return cnt > 0 && cnt < 3
}

func hasExactlyOneNeighbour(i int, j int) bool {
	neighbours := getNeighbours(i, j)

	cnt := 0
	for _, v := range neighbours {
		if v {
			cnt++
		}
	}

	return cnt == 1
}

func getNeighbours(i int, j int) []bool {
	var neighbours []bool

	//up
	if i > 0 {
		neighbours = append(neighbours, virusGrid[i-1][j])
	}

	//down
	if i < 4 {
		neighbours = append(neighbours, virusGrid[i+1][j])
	}

	//left
	if j > 0 {
		neighbours = append(neighbours, virusGrid[i][j-1])
	}

	//down
	if j < 4 {
		neighbours = append(neighbours, virusGrid[i][j+1])
	}

	return neighbours
}
