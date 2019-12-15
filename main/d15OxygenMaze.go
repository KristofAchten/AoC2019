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

var (
	visited   map[coords]bool
	maze      [][]string
	oxygenPos coords
)

const (
	// All of the numbers below originate from trial and error.
	shiftX   = 21
	shiftY   = 22
	mazeDim  = 41
	startDir = 2

	oxygenSystem = "@"
)

func day15() {
	start := time.Now()

	initializeMaze()

	input := strings.Split(string(getPuzzleInput("input/day15.txt")), ",")
	mazeTree := buildMaze(createDefaultIntcodeState(stringSliceToIntSlice(input), []int64{startDir}), coords{shiftX, shiftY}, 0)
	maze[shiftY][shiftX] = "X"

	fmt.Println("The maze looks like this (@ = oxygen system, # = wall, X = starting point):")
	drawMaze()

	printPuzzleResult(15, searchOxygen(mazeTree, 0), floodOxygen(oxygenPos, 0))

	fmt.Printf("DAY 15 STATS: Execution took %s\n\n", time.Since(start))
}

func initializeMaze() {
	visited = make(map[coords]bool)
	maze = make([][]string, mazeDim)
	for i := 0; i < mazeDim; i++ {
		maze[i] = make([]string, mazeDim)
		for j := 0; j < mazeDim; j++ {
			maze[i][j] = "#"
		}
	}

}

func floodOxygen(curPos coords, curStep int) int {
	curVal := maze[curPos.y][curPos.x]

	if curVal == empty {
		maze[curPos.y][curPos.x] = ball
	} else if curVal != oxygenSystem {
		return curStep - 1
	}

	var newResult int
	for _, v := range []int64{1, 2, 3, 4} {
		newResult = maxInt(newResult, floodOxygen(updateLoc(curPos, v), curStep+1))
	}
	return maxInt(newResult, curStep)
}

func buildMaze(input intcodeState, loc coords, steps int) cell {
	var desc []cell
	var backtrack bool

	visited[loc] = true
	output, _, newState := runIntCode(input)
	maze[loc.y][loc.x] = i2viz(int(output))

	if output == 0 {
		backtrack = true
	}

	if !backtrack {
		for _, v := range []int64{1, 2, 3, 4} {
			newloc := updateLoc(loc, v)
			if _, ok := visited[newloc]; !ok {
				desc = append(desc, buildMaze(intcodeState{newState.program, newState.ptr, newState.relativeBase, []int64{v}}, newloc, steps+1))
			}
		}
	}

	return cell{output, desc, loc}
}

func searchOxygen(node cell, steps int) int {
	steps++

	if node.val == 2 {
		oxygenPos = node.physPos
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
		return oxygenSystem
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
		panic("Provided value does not map to a valid directional change")
	}
}

func drawMaze() {
	for x := range maze {
		fmt.Println(maze[x])
	}
}
