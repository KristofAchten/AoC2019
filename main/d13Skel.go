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

	fmt.Println("Day 13: solution one is ", countBlocks(runGame(input)))
	fmt.Println("Day 13: solution two is " + "NOT IMPLEMENTED")

	fmt.Printf("DAY 13 STATS: Execution took %s\n\n", time.Since(start))
}

func runGame(code []int64) [][]string {
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

	printWorld(world)
	return world
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
		panic("Invalid object")
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
