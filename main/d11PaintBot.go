package main

import (
	"fmt"
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

func displayGrid(coordinates map[coords]int64) {

	shiftedCoords, maxX, maxY := shiftArray(coordinates)
	printFriendlyMap := make(map[int][]string)

	for y := 0; y <= maxY; y++ {
		var row []string
		for x := 0; x <= maxX; x++ {
			if v, ok := shiftedCoords[coords{x, y}]; ok && v == 1 {
				row = append(row, "X")
			} else {
				row = append(row, " ")
			}
		}
		printFriendlyMap[y] = row
	}

	for i := 0; i < len(printFriendlyMap); i++ {
		fmt.Println(reverseStringSlice(printFriendlyMap[i]))
	}
}

func shiftArray(coordGrid map[coords]int64) (map[coords]int64, int, int) {
	var minX, minY, maxX, maxY int

	for k := range coordGrid {
		if k.x < minX {
			minX = k.x
		} else if k.x > maxX {
			maxX = k.x
		}

		if k.y < minY {
			minY = k.y
		} else if k.y > maxY {
			maxY = k.y
		}
	}

	if minX < 0 {
		maxX += abs(minX)
	}

	if minY < 0 {
		maxX += abs(minY)
	}

	return shiftCoordinates(coordGrid, minX, minY), maxX, maxY
}

func shiftCoordinates(grid map[coords]int64, x int, y int) map[coords]int64 {
	if x == 0 && y == 0 {
		return grid
	}

	shiftedCoords := make(map[coords]int64)

	for k, v := range grid {
		shiftedCoords[coords{k.x + abs(x), k.y + abs(y)}] = v
	}

	return shiftedCoords
}
