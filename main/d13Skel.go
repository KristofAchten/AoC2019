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
	paddle = "P"
	ball   = "@"
)

func day13() {
	start := time.Now()

	input := stringSliceToIntSlice(strings.Split(string(getPuzzleInput("input/day13.txt")), ","))

	fmt.Println(countBlocks(runGame(input)))
	printWorld(runGame(input))
	fmt.Println("Day 13: solution one is " + "NOT IMPLEMENTED")
	fmt.Println("Day 13: solution two is " + "NOT IMPLEMENTED")

	fmt.Printf("DAY 13 STATS: Execution took %s\n\n", time.Since(start))
}

func runGame(code []int64) [][]string {
	world := make([][]string, 24)

	for i := 0; i < 24; i++ {
		world[i] = make([]string, 50)
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

		var add string
		switch v {
		case 0:
			add = empty
		case 1:
			add = wall
		case 2:
			add = block
		case 3:
			add = paddle
		case 4:
			add = ball
		default:
			panic("Invalid object being spawned in the world")
		}

		curValue := world[y][x]

		if curValue == wall || curValue == paddle {
			continue
		}

		if curValue == block && add != ball {
			continue
		}

		world[y][x] = add
	}

	return world
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
