package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Coords struct {
	x     int
	y     int
	steps int
}

func day3() {
	input := strings.Split(string(GetPuzzleInput("input/day3.txt")), "\n")
	runFull := runWires(input)

	fmt.Println("Day 3: solution one is " + strconv.Itoa(determineClosestOverlap(runFull)))
	fmt.Println("Day 2: solution two is " + strconv.Itoa(determineFewestDistance(determineOverlap(runFull), input)))

}

func runWires(input []string) []([]Coords) {
	coords, _ := runWiresWithSteps(input, Coords{})
	return coords
}

func runWiresWithSteps(input []string, overlap Coords) ([]([]Coords), int) {
	var coordSets []([]Coords)
	var totalSteps = 0
	for _, v := range input {
		cmds := strings.Split(v, ",")
		var visited []Coords
		var curCoords = Coords{0, 0, 0}

		for i, sv := range cmds {
			steps, _ := strconv.Atoi(string(sv[1:]))
			totalSteps += steps

			switch string(sv[0]) {
			case "U":
				curCoords, visited = visit(i, 0, +1, steps, visited, curCoords)
			case "D":
				curCoords, visited = visit(i, 0, -1, steps, visited, curCoords)
			case "L":
				curCoords, visited = visit(i, -1, 0, steps, visited, curCoords)
			case "R":
				curCoords, visited = visit(i, +1, 0, steps, visited, curCoords)
			default:
				panic("Cmd not supported :'( " + string(sv[0]))
			}
			if overlap.x == curCoords.x && overlap.y == curCoords.y {
				break
			}
		}
		coordSets = append(coordSets, visited)
	}
	return coordSets, totalSteps
}

func visit(wire int, x int, y int, steps int, visited []Coords, curCoords Coords) (Coords, []Coords) {
	if x == 0 {
		fixedX := curCoords.x
		curY := curCoords.y
		curSteps := curCoords.steps
		for i := 0; i < steps; i++ {
			cur := Coords{fixedX, curY, curSteps + i}
			if !contains(visited, cur) {
				visited = append(visited, cur)
				curY += y
			}
		}
		return Coords{curCoords.x, curCoords.y + (steps * y), curSteps + steps}, visited
	} else {
		fixedY := curCoords.y
		curX := curCoords.x
		curSteps := curCoords.steps
		for i := 0; i < steps; i++ {
			cur := Coords{curX, fixedY, curSteps + i}
			if !contains(visited, cur) {
				visited = append(visited, cur)
				curX += x
			}
		}
		return Coords{curCoords.x + (steps * x), curCoords.y, curSteps + steps}, visited
	}
	panic("shouldn't be here")
}

func determineOverlap(coordSets []([]Coords)) []Coords {
	var overlap []Coords
	if len(coordSets) < 2 {
		panic("Not enough coordsets provided!")
	}

	for _, v := range coordSets[0] {
		if contains(coordSets[1], v) {
			overlap = append(overlap, v)
		}
	}
	return overlap
}

func determineClosestOverlap(coordSets []([]Coords)) int {
	overlap := determineOverlap(coordSets)
	var curBest = 99999999
	for _, v := range overlap {
		val := Abs(v.x) + Abs(v.y)
		if val != 0 && val < curBest {
			curBest = val
		}
	}

	return curBest
}

func determineFewestDistance(overlap []Coords, input []string) int {
	var curBest = 9999999
	for _, v := range overlap {
		if v.steps != 0 && v.steps < curBest {
			curBest = v.steps
		}
	}
	return curBest
}

func contains(slice []Coords, val Coords) bool {
	for _, v := range slice {
		if v.x == val.x && v.y == val.y {
			return true
		}
	}
	return false
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
