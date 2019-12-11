package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func day11() {
	start := time.Now()

	input := stringSliceToIntSlice(strings.Split(string(getPuzzleInput("input/day11.txt")), ","))
	res1 := len(runPaintBot(input, 0))
	res2 := runPaintBot(input, 1)

	fmt.Println("Day 11: solution one is " + strconv.Itoa(res1))
	fmt.Println("Day 11: solution two is")
	displayGrid(res2)

	confirmPuzzleResult(11, res1, len(res2))

	fmt.Printf("DAY 11 STATS: Execution took %s\n\n", time.Since(start))
}

func runPaintBot(code []int64, startWith int64) map[coords]int64 {
	grid := make(map[coords]int64)
	curCoords := coords{0, 0}
	direction := vec{0, -1}
	state := createDefaultIntcodeState(code, []int64{startWith})

	for {
		// Retrieve & update colour
		output, halt, newState := runIntCode(state)

		if halt {
			break
		}

		grid[curCoords] = output

		// Retrieve & update direction
		sndOutput, _, sndNewState := runIntCode(newState)

		direction = turn(direction, sndOutput)
		state = sndNewState

		// Move one step
		curCoords = coords{curCoords.x + int(direction.x), curCoords.y + int(direction.y)}

		// Update the input for the next run (default = 0 = black)
		if v, ok := grid[curCoords]; ok {
			state.input = append(state.input, v)
		} else {
			state.input = append(state.input, 0)
		}
	}

	return grid
}

func turn(curDir vec, with int64) vec {
	if curDir.x != 0 {
		switch with {
		case 0:
			return vec{0, curDir.x}
		case 1:
			return vec{0, curDir.x * -1}
		}
	} else {
		switch with {
		case 0:
			return vec{curDir.y * -1, 0}
		case 1:
			return vec{curDir.y, 0}
		}
	}
	panic("Invalid input vector")
}

func displayGrid(coords map[coords]int64) {

	minCoords, maxCoords := determineCorners(coords)
	printFriendlyMap := make(map[int][]string)

	// This isn't necessarily what I wanted but it works.
	xSize := maxCoords.x + int(math.Abs(float64(minCoords.x)))
	ySize := maxCoords.y + int(math.Abs(float64(minCoords.y)))

	for y := 0; y <= ySize; y++ {
		var valArray []string
		for x := 0; x <= xSize; x++ {
			valArray = append(valArray, " ")
		}

		printFriendlyMap[y] = valArray
	}

	for k, v := range coords {
		if v == 1 {
			printFriendlyMap[k.y][k.x+xSize] = "X"
		}
	}

	for i := 0; i < ySize; i++ {
		fmt.Println(reverseStringSlice(printFriendlyMap[i]))
	}
}

func determineCorners(coordGrid map[coords]int64) (coords, coords) {
	var minx, miny, maxx, maxy int

	for k := range coordGrid {
		if k.x < minx || minx == 0 {
			minx = k.x
		} else if k.x > maxx || maxx == 0 {
			maxx = k.x
		}

		if k.y < miny || miny == 0 {
			miny = k.y
		} else if k.y > maxy || maxy == 0 {
			maxy = k.y
		}
	}

	return coords{minx, miny}, coords{maxx, maxy}
}
