package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Coords struct {
	x int
	y int
}

func day3() {
	start := time.Now()

	input := strings.Split(string(getPuzzleInput("input/day3.txt")), "\n")
	overlap, allSteps := overlap(traceWires(input))

	fmt.Println("Day 3: solution one is " + strconv.Itoa(determineOverlapWithSmallestManhattanDistance(overlap)))
	fmt.Println("Day 3: solution two is " + strconv.Itoa(determineOverlapWithLeastSteps(allSteps)))

	fmt.Printf("DAY 3 STATS: Execution took %s\n\n", time.Since(start))
}

func traceWires(input []string) [][]Coords {
	var coordSets [][]Coords

	for _, v := range input {
		cmds := strings.Split(v, ",")

		var visited []Coords
		var curCoord = Coords{0, 0}

		for _, sv := range cmds {
			steps, _ := strconv.Atoi(sv[1:])

			switch string(sv[0]) {
			case "U":
				curCoord, visited = visit(0, +1, steps, visited, curCoord)
			case "D":
				curCoord, visited = visit(0, -1, steps, visited, curCoord)
			case "L":
				curCoord, visited = visit(-1, 0, steps, visited, curCoord)
			case "R":
				curCoord, visited = visit(+1, 0, steps, visited, curCoord)
			default:
				panic("Invalid input " + string(sv[0]) + ": only U(p), D(own), L(eft) and R(ight) are supported.")
			}
		}
		coordSets = append(coordSets, visited)
	}
	return coordSets
}

func visit(x int, y int, steps int, visited []Coords, curCoords Coords) (Coords, []Coords) {
	curX := curCoords.x
	curY := curCoords.y

	for i := 0; i < steps; i++ {
		cur := Coords{curX, curY}
		visited = append(visited, cur)

		curX += x
		curY += y
	}
	return Coords{curCoords.x + (steps * x), curCoords.y + (steps * y)}, visited
}

func determineOverlapWithSmallestManhattanDistance(overlap []Coords) int {

	var curBest = 9999999999
	for _, v := range overlap {
		val := abs(v.x) + abs(v.y)
		if val != 0 && val < curBest {
			curBest = val
		}
	}

	return curBest
}

func determineOverlapWithLeastSteps(totalSteps []int) int {
	var curBest = 9999999999
	for _, v := range totalSteps {
		if v != 0 && v < curBest {
			curBest = v
		}
	}
	return curBest
}

/**
Helper functions
*/

func overlap(coordSets [][]Coords) ([]Coords, []int) {
	var overlaps []Coords
	var totalSteps []int

	m := sliceToMap(coordSets[0])

	for idx2, v2 := range coordSets[1] {
		fv, ok := m[v2]
		if ok {
			overlaps = append(overlaps, v2)
			totalSteps = append(totalSteps, fv+idx2)
		}
	}

	return overlaps, totalSteps
}

func sliceToMap(set []Coords) map[Coords]int {
	m := make(map[Coords]int)

	for idx1, v1 := range set {
		m[v1] = idx1
	}

	return m
}
