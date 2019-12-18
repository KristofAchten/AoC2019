package main

import (
	"fmt"
	"strings"
	"time"
)

/**
I manually traced the scaffolding because it was easier.
*/
var seq = [5]string{
	"A,A,B,C,C,A,C,B,C,B",
	"L,4,L,4,L,6,R,10,L,6",
	"L,12,L,6,R,10,L,6",
	"R,8,R,10,L,6",
	"n",
}

var scaffolding map[coords]int64

func day17() {
	start := time.Now()

	input := stringSliceToIntSlice(strings.Split(string(getPuzzleInput("/input/day17.txt")), ","))
	scaffolding = make(map[coords]int64)

	buildScaffolding(createDefaultIntcodeState(input, []int64{}))
	fmt.Println("The scaffolding looks like this:")
	printScaffolding()

	printPuzzleResult(17, calculateAlignmentParams(), startSucking(createDefaultIntcodeState(input, []int64{})))

	fmt.Printf("DAY 17 STATS: Execution took %s\n\n", time.Since(start))
}

func startSucking(state intcodeState) int {
	state.program[0] = 2

	var input []int64
	for _, in := range seq {
		for _, c := range in {
			input = append(input, int64(c))
		}
		input = append(input, int64('\n'))
	}

	state.input = input
	return int(runUntilHalt(state))
}

func calculateAlignmentParams() int {
	sum := 0
	for row := 0; row < 35; row++ {
		for col := 0; col < 45; col++ {
			val := scaffolding[coords{col, row}]

			if val == '#' {
				up := scaffolding[coords{col, row - 1}]
				down := scaffolding[coords{col, row + 1}]
				left := scaffolding[coords{col - 1, row}]
				right := scaffolding[coords{col + 1, row}]
				if up == val && down == val && left == val && right == val {
					sum += row * col
				}
			}
		}
	}
	return sum
}

func buildScaffolding(state intcodeState) intcodeState {
	row := 0
	col := 0
	for {
		output, halt, newstate := runIntCode(state)

		if halt {
			break
		}

		if output == 10 {
			row++
			col = 0
		} else {
			scaffolding[coords{col, row}] = output
			col++
		}

		state = newstate
	}
	return state
}

func printScaffolding() {
	for row := 0; row < 35; row++ {
		for col := 0; col < 45; col++ {
			fmt.Printf("%c", rune(scaffolding[coords{col, row}]))
		}
		fmt.Println()
	}
}
