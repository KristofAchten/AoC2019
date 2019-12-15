package main

import (
	"fmt"
	"strings"
	"time"
)

type intcodeNode struct {
	val      int64
	children []intcodeNode
}

var visited map[coords]bool

func day15() {
	start := time.Now()

	visited = make(map[coords]bool)

	input := strings.Split(string(getPuzzleInput("input/day15.txt")), ",")
	fullMaze := buildTree(createDefaultIntcodeState(stringSliceToIntSlice(input), []int64{2}), coords{0, 0}, 0)

	printPuzzleResult(15, searchOxygen(fullMaze, 0), notImplemented)

	fmt.Printf("DAY 15 STATS: Execution took %s\n\n", time.Since(start))
}

func buildTree(input intcodeState, loc coords, steps int) intcodeNode {
	dirs := []int64{1, 2, 3, 4}
	var children []intcodeNode
	visited[loc] = true

	output, _, newState := runIntCode(input)
	steps++
	var backtrack bool
	if output == 0 {
		backtrack = true
	}

	if !backtrack {
		for _, v := range dirs {
			newloc := updateLoc(loc, v)

			if _, ok := visited[newloc]; !ok {
				children = append(children, buildTree(intcodeState{newState.program, newState.ptr, newState.relativeBase, []int64{v}}, newloc, steps))
			}
		}
	}

	return intcodeNode{output, children}
}

func searchOxygen(node intcodeNode, steps int) int {
	steps++

	if node.val == 2 {
		return steps
	}

	var stepsReq int
	for _, v := range node.children {
		stepsReq += searchOxygen(v, steps)
	}

	return stepsReq
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
