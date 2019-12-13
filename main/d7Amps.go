package main

import (
	"fmt"
	"strings"
	"time"
)

func day7() {
	start := time.Now()
	input := strings.Split(string(getPuzzleInput("input/day7.txt")), ",")

	res1 := findLargestOutputForAmps(input, permutations([]int64{0, 1, 2, 3, 4}))
	res2 := findLargestOutputForAmpsWithFeedbackLoop(input, permutations([]int64{5, 6, 7, 8, 9}))

	printPuzzleResult(7, int(res1), int(res2))

	fmt.Printf("DAY 7 STATS: Execution took %s\n\n", time.Since(start))
}

func findLargestOutputForAmps(input []string, possibleInputs [][]int64) int64 {
	var allOutputs []int64

	for _, v := range possibleInputs {
		var nextInput int64
		for _, v2 := range v {
			nextInput = runUntilHalt(createDefaultIntcodeState(stringSliceToIntSlice(input), []int64{int64(v2), nextInput}))
		}
		allOutputs = append(allOutputs, nextInput)
	}

	return max(allOutputs)
}

func findLargestOutputForAmpsWithFeedbackLoop(inputCode []string, possibleInputs [][]int64) int64 {
	var allOutputs []int64
	for _, v := range possibleInputs {
		var lastOutput int64
		execCtr := 0
		codePerAmp := []intcodeState{
			createDefaultIntcodeState(stringSliceToIntSlice(inputCode), []int64{v[0], 0}),
			createDefaultIntcodeState(stringSliceToIntSlice(inputCode), []int64{v[1]}),
			createDefaultIntcodeState(stringSliceToIntSlice(inputCode), []int64{v[2]}),
			createDefaultIntcodeState(stringSliceToIntSlice(inputCode), []int64{v[3]}),
			createDefaultIntcodeState(stringSliceToIntSlice(inputCode), []int64{v[4]}),
		}

		for {
			output, halt, newState := runIntCode(codePerAmp[execCtr%5])

			if halt {
				allOutputs = append(allOutputs, lastOutput)
				break
			}

			nextAmp := codePerAmp[(execCtr+1)%5]
			codePerAmp[execCtr%5] = intcodeState{newState.program, newState.ptr, newState.relativeBase, reverseIntSlice(newState.input)}
			codePerAmp[(execCtr+1)%5] = intcodeState{nextAmp.program, nextAmp.ptr, newState.relativeBase, append(nextAmp.input, output)}

			lastOutput = output
			execCtr++
		}

	}
	return max(allOutputs)
}
