package main

import (
	"fmt"
	"strings"
	"time"
)

type cell struct {
	val        int64
	neighbours []cell
	physPos    coords
}

var visited map[coords]bool
var maze [][]string

func day15() {
	start := time.Now()

	visited = make(map[coords]bool)
	maze = make([][]string, 41)
	for i := 0; i < 41; i++ {
		maze[i] = make([]string, 41)
		for j := 0; j < 41; j++ {
			maze[i][j] = "#"
		}
	}

	input := strings.Split(string(getPuzzleInput("input/day15.txt")), ",")
	fullMaze := buildTree(createDefaultIntcodeState(stringSliceToIntSlice(input), []int64{2}), coords{21, 22}, 0)

	printPuzzleResult(15, searchOxygen(fullMaze, 0), notImplemented)

	fmt.Println("The maze looked like this (@ = oxygen system, # = wall):")
	drawMaze()

	fmt.Printf("DAY 15 STATS: Execution took %s\n\n", time.Since(start))
}

func buildTree(input intcodeState, loc coords, steps int) cell {
	dirs := []int64{1, 2, 3, 4}
	var neightbours []cell
	visited[loc] = true

	output, _, newState := runIntCode(input)

	maze[loc.y][loc.x] = i2viz(int(output))

	steps++
	var backtrack bool
	if output == 0 {
		backtrack = true
	}

	if !backtrack {
		for _, v := range dirs {
			newloc := updateLoc(loc, v)

			if _, ok := visited[newloc]; !ok {
				neightbours = append(neightbours, buildTree(intcodeState{newState.program, newState.ptr, newState.relativeBase, []int64{v}}, newloc, steps))
			}
		}
	}

	return cell{output, neightbours, loc}
}

func searchOxygen(node cell, steps int) int {
	steps++

	if node.val == 2 {
		return steps
	}

	var stepsReq int
	for _, v := range node.neighbours {
		stepsReq += searchOxygen(v, steps)
	}

	return stepsReq
}

func i2viz(input int) string {
	switch input {
	case 0:
		return wall
	case 1:
		return empty
	case 2:
		return "@"
	default:
		panic("Wrong input")
	}
}

func updateLoc(loc coords, v int64) coords {
	switch v {
	case 1: // North
		return coords{loc.x, loc.y - 1}
	case 2: // South
		return coords{loc.x, loc.y + 1}
	case 3: // West
		return coords{loc.x - 1, loc.y}
	case 4: // East
		return coords{loc.x + 1, loc.y}
	default:
		panic("aaaaah")
	}
}

func drawMaze() {
	for x := range maze {
		fmt.Println(maze[x])
	}
}
