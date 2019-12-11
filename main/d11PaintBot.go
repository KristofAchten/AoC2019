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
	displayGrid(runPaintBot(input, 1))

	fmt.Println("Day 11: solution one is " + strconv.Itoa(res1))
	fmt.Println("Day 11: solution two is " + "NOT IMPLEMENTED")

	fmt.Printf("DAY 11 STATS: Execution took %s\n\n", time.Since(start))
}

func runPaintBot(code []int64, startWith int64) map[coords]int64 {
	grid := make(map[coords]int64)
	curCoords := coords{0, 0}
	direction := vec{0, -1}
	state := createDefaultIntcodeState(code, []int64{startWith})

	for {
		output, halt, newState := runIntCode(state)

		if halt {
			break
		}

		grid[curCoords] = output

		sndOutput, sndHalt, sndNewState := runIntCode(newState)

		if sndHalt {
			break
		}

		direction = turn(direction, sndOutput)
		state = sndNewState
		curCoords = coords{curCoords.x + int(direction.x), curCoords.y + int(direction.y)}

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
	var minx int
	var miny int
	var maxx int
	var maxy int

	for k := range coords {
		if k.x < minx {
			minx = k.x
		} else if k.x > maxx {
			maxx = k.x
		}

		if k.y < miny {
			miny = k.y
		} else if k.y > maxy {
			maxy = k.y
		}
	}

	xSize := maxx + int(math.Abs(float64(minx)))
	ySize := maxy + int(math.Abs(float64(miny)))
	fmt.Println(xSize, ySize)

	printmap := make(map[int][]string)
	for y := 0; y <= ySize; y++ {
		var valArray []string
		for x := 0; x <= xSize; x++ {
			valArray = append(valArray, " ")
		}
		printmap[y] = valArray
	}

	for k, v := range coords {
		if v == 1 {
			printmap[k.y][k.x+xSize] = "X"
		}
	}

	for i := 0; i <= ySize; i++ {
		fmt.Println(printmap[i])
	}
}
