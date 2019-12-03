package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Coords struct {
	x     int
	y     int
	steps int
}

func day3() {
	start := time.Now()

	input := strings.Split(string(GetPuzzleInput("input/day3.txt")), "\n")
	overlap := overlap(traceWires(input))

	fmt.Println("Day 3: solution one is " + strconv.Itoa(determineOverlapWithSmallestManhattanDistance(overlap)))
	fmt.Println("Day 3: solution two is " + strconv.Itoa(determineOverlapWithLeastSteps(overlap)))

	fmt.Printf("DAY 3 STATS: Execution took %s\n\n", time.Since(start))

}

func traceWires(input []string) []([]Coords) {
	var coordSets []([]Coords)

	for _, v := range input {
		cmds := strings.Split(v, ",")

		var visited []Coords
		var curCoord = Coords{0, 0, 0}

		for i, sv := range cmds {
			steps, _ := strconv.Atoi(string(sv[1:]))

			switch string(sv[0]) {
			case "U":
				curCoord, visited = visit(i, 0, +1, steps, visited, curCoord)
			case "D":
				curCoord, visited = visit(i, 0, -1, steps, visited, curCoord)
			case "L":
				curCoord, visited = visit(i, -1, 0, steps, visited, curCoord)
			case "R":
				curCoord, visited = visit(i, +1, 0, steps, visited, curCoord)
			default:
				panic("Invalid input " + string(sv[0]) + ": only U(p), D(own), L(eft) and R(ight) are supported.")
			}
		}
		coordSets = append(coordSets, visited)
	}
	return coordSets
}

func visit(wire int, x int, y int, steps int, visited []Coords, curCoords Coords) (Coords, []Coords) {
	curX := curCoords.x
	curY := curCoords.y
	curSteps := curCoords.steps

	for i := 0; i < steps; i++ {
		cur := Coords{curX, curY, curSteps}
		visited = append(visited, cur)

		curX += x
		curY += y
		curSteps += 1
	}
	return Coords{curCoords.x + (steps * x), curCoords.y + (steps * y), curCoords.steps + steps}, visited
}

func determineOverlapWithSmallestManhattanDistance(overlap []Coords) int {

	var curBest = 9999999999
	for _, v := range overlap {
		val := Abs(v.x) + Abs(v.y)
		if val != 0 && val < curBest {
			curBest = val
		}
	}

	return curBest
}

func determineOverlapWithLeastSteps(overlap []Coords) int {
	var curBest = 9999999999
	for _, v := range overlap {
		if v.steps != 0 && v.steps < curBest {
			curBest = v.steps
		}
	}
	return curBest
}

/**
Helper functions
*/

func overlap(coordSets []([]Coords)) []Coords {
	var overlap []Coords
	for _, v := range coordSets[0] {
		if contains(coordSets[1], v) {
			overlap = append(overlap, Coords{v.x, v.y, v.steps + getSteps(coordSets[1], v)})
		}
	}
	return overlap
}

func contains(slice []Coords, val Coords) bool {
	for _, v := range slice {
		if v.x == val.x && v.y == val.y {
			return true
		}
	}
	return false
}

func getSteps(coordSets []Coords, val Coords) int {
	for _, v := range coordSets {
		if v.x == val.x && v.y == val.y {
			return v.steps
		}
	}
	panic("Coordinates not found in the provided slice.")
}
