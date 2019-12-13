package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	empty  = " "
	wall   = "#"
	block  = "B"
	paddle = "="
	ball   = "O"
)

func day13() {
	start := time.Now()

	input := stringSliceToIntSlice(strings.Split(string(getPuzzleInput("input/day13.txt")), ","))

	res1 := countBlocks(runGame(input, false))
	res2 := winGame(input)

	printPuzzleResult(13, res1, int(res2))

	fmt.Printf("DAY 13 STATS: Execution took %s\n\n", time.Since(start))
}

func runGame(code []int64, print bool) [][]string {
	world := make([][]string, 24)

	for i := 0; i < 24; i++ {
		world[i] = make([]string, 45)
	}
	state := createDefaultIntcodeState(code, []int64{})

	for {
		x, halt, newstate := runIntCode(state)

		if halt {
			break
		}

		y, _, sndNewState := runIntCode(newstate)
		v, _, thrdNewState := runIntCode(sndNewState)

		state = thrdNewState
		world[y][x] = i2obj(v)
	}

	if print {
		printWorld(world)
	}
	return world
}

func winGame(code []int64) int64 {
	code[0] = 2
	state := createDefaultIntcodeState(code, []int64{})

	var score int64
	var paddleXPos int64

	for {
		x, halt, newState := runIntCode(state)

		if halt {
			break
		}

		y, _, sndNewState := runIntCode(newState)
		v, _, thrdNewState := runIntCode(sndNewState)
		state = thrdNewState

		var input int64 // Default input = 0

		if x == -1 && y == 0 {
			score = v
		} else if i2obj(v) == paddle {
			paddleXPos = x
		} else if i2obj(v) == ball { // super advanced AI
			if paddleXPos > x {
				input = -1
			} else if paddleXPos < x {
				input = 1
			}
		}

		state.input = []int64{input}
	}
	return score
}

func i2obj(i int64) string {
	switch i {
	case 0:
		return empty
	case 1:
		return wall
	case 2:
		return block
	case 3:
		return paddle
	case 4:
		return ball
	default:
		panic("Invalid object identifier")
	}
}

func countBlocks(input [][]string) int {
	var count int
	for i := range input {
		for j := range input[i] {
			if input[i][j] == block {
				count++
			}
		}
	}
	return count
}

func printWorld(input [][]string) {
	for i := range input {
		fmt.Println(input[i])
	}
}
