package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func day7() {
	start := time.Now()
	input := strings.Split(string(getPuzzleInput("input/day7.txt")), ",")

	res1 := findLargestOutputForAmps(input, permutations([]int{0, 1, 2, 3, 4}))
	res2 := findLargestOutputForAmpsWithFeedbackLoop(input, permutations([]int{5, 6, 7, 8, 9}))

	fmt.Println("Day 7: solution one is " + strconv.Itoa(res1))
	fmt.Println("Day 7: solution two is " + strconv.Itoa(res2))

	confirmPuzzleResult(7, res1, res2)

	fmt.Printf("DAY 7 STATS: Execution took %s\n\n", time.Since(start))
}

func findLargestOutputForAmps(input []string, possibleInputs [][]int) int {
	var allOutputs []int

	for _, v := range possibleInputs {
		var nextInput int
		for _, v2 := range v {
			nextInput = runUntilHalt(createDefaultIntcodeState(stringSliceToIntSlice(input), []int{v2, nextInput}))
		}
		allOutputs = append(allOutputs, nextInput)
	}

	return max(allOutputs)
}

func findLargestOutputForAmpsWithFeedbackLoop(inputCode []string, possibleInputs [][]int) int {
	var allOutputs []int
	for _, v := range possibleInputs {
		var lastOutput int
		execCtr := 0
		codePerAmp := []intcodeState{
			{stringSliceToIntSlice(inputCode), 0, []int{v[0], 0}},
			{stringSliceToIntSlice(inputCode), 0, []int{v[1]}},
			{stringSliceToIntSlice(inputCode), 0, []int{v[2]}},
			{stringSliceToIntSlice(inputCode), 0, []int{v[3]}},
			{stringSliceToIntSlice(inputCode), 0, []int{v[4]}},
		}

		for {
			output, halt, newState := runIntCode(codePerAmp[execCtr%5])

			if halt {
				allOutputs = append(allOutputs, lastOutput)
				break
			}

			nextAmp := codePerAmp[(execCtr+1)%5]
			codePerAmp[execCtr%5] = intcodeState{newState.program, newState.ptr, reverseIntSlice(newState.input)}
			codePerAmp[(execCtr+1)%5] = intcodeState{nextAmp.program, nextAmp.ptr, append(nextAmp.input, output)}

			lastOutput = output
			execCtr++
		}

	}
	return max(allOutputs)
}
