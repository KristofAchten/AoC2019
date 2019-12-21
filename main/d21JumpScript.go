package main

import (
	"fmt"
	"strings"
	"time"
)

const base = "NOT A J\n" +
	"NOT B T\n" +
	"OR T J\n" +
	"NOT C T\n" +
	"OR T J\n" +
	"AND D J\n"

const p1 = base +
	"WALK\n"

const p2 = base +
	"NOT H T\n" +
	"NOT T T\n" +
	"OR E T\n" +
	"AND T J\n" +
	"RUN\n"

func day21() {
	start := time.Now()

	input := stringSliceToIntSlice(strings.Split(string(getPuzzleInput("input/day21.txt")), ","))
	printPuzzleResult(21, runJumpBot(input, p1, false), runJumpBot(input, p2, false))

	fmt.Printf("DAY 21 STATS: Execution took %s\n\n", time.Since(start))
}

func runJumpBot(code []int64, cmds string, print bool) int {
	var input []int64
	for _, in := range cmds {
		input = append(input, int64(in))
	}

	state := createDefaultIntcodeState(code, input)
	var result int64
	for {
		var halt bool
		output, halt, newState := runIntCode(state)

		if halt {
			break
		}

		if print {
			fmt.Printf("%c", output)
		}
		result = output
		state = newState
	}
	return int(result)
}
