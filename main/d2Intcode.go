package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func day2() {
	start := time.Now()

	input := strings.Split(string(getPuzzleInput("input/day2.txt")), ",")

	fmt.Println("Day 2: solution one is " + strconv.Itoa(runBasicIntCode(stringSliceToIntSlice(input), 12, 2)))
	fmt.Println("Day 2: solution two is " + getVerbAndNoun(stringSliceToIntSlice(input), 19690720))

	fmt.Printf("DAY 2 STATS: Execution took %s\n\n", time.Since(start))
}

func runBasicIntCode(code []int, noun int, verb int) int {
	pointer := 0

	code[1] = noun
	code[2] = verb

	for pointer < len(code) {
		val := code[pointer]
		switch val {
		case 1:
			code[code[pointer+3]] = code[code[pointer+1]] + code[code[pointer+2]]
			pointer += 4
		case 2:
			code[code[pointer+3]] = code[code[pointer+1]] * code[code[pointer+2]]
			pointer += 4
		default:
			pointer = 99999
		}
	}

	return code[0]
}

func getVerbAndNoun(input []int, res int) string {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			code := make([]int, len(input))
			copy(code, input)
			if runBasicIntCode(code, i, j) == res {
				return strconv.Itoa(100*i + j)
			}
		}
	}
	return "nope, fucked"
}

func stringSliceToIntSlice(input []string) []int {
	output := make([]int, len(input))

	for i, v := range input {
		output[i], _ = strconv.Atoi(v)
	}

	return output
}
