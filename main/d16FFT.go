package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func day16() {
	start := time.Now()
	input := string(getPuzzleInput("/input/day16.txt"))

	printPuzzleResult(16, phaseResult(input, 100), decodeSignal(input, 100))

	fmt.Printf("DAY 16 STATS: Execution took %s\n\n", time.Since(start))
}

func phaseResult(input string, iterations int) int {

	pattern := []int{0, 1, 0, -1}

	curString := input

	for r := 0; r < iterations; r++ {
		var newString string
		for j := range curString {
			var sum int
			for i := range curString {
				curVal, _ := strconv.Atoi(string(curString[i]))
				sum += curVal * pattern[((i+1)/(j+1))%4]
			}
			stringVal := strconv.Itoa(abs(sum) % 10)
			newString += stringVal
		}
		curString = newString
	}
	res, _ := strconv.Atoi(curString[:8])
	return res
}

func decodeSignal(input string, iterations int) int {
	input = strings.Repeat(input, 10000)
	messOffset, _ := strconv.Atoi(input[0:7])
	interestingBits := toIntSlice(input[messOffset:])

	for r := 0; r < iterations; r++ {
		var sum int
		for i := len(interestingBits) - 1; i >= 0; i-- {
			sum += interestingBits[i]
			interestingBits[i] = sum % 10
		}
	}

	var finalResult string
	for i := range interestingBits[:8] {
		finalResult += strconv.Itoa(interestingBits[i])
	}

	res, _ := strconv.Atoi(finalResult)
	return res
}

func toIntSlice(s string) []int {
	var values []int
	for i := range s {
		v, _ := strconv.Atoi(string(s[i]))
		values = append(values, v)
	}
	return values
}
