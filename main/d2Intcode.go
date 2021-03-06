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

	res1 := runBasicIntCode(stringSliceToIntSlice(input), 12, 2)
	res2, _ := strconv.Atoi(getVerbAndNoun(stringSliceToIntSlice(input), 19690720))

	printPuzzleResult(2, int(res1), res2)

	fmt.Printf("DAY 2 STATS: Execution took %s\n\n", time.Since(start))
}

func runBasicIntCode(code []int64, noun int64, verb int64) int64 {
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

func getVerbAndNoun(input []int64, res int64) string {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			code := make([]int64, len(input))
			copy(code, input)
			if runBasicIntCode(code, int64(i), int64(j)) == res {
				return strconv.Itoa(100*i + j)
			}
		}
	}
	return "nope, fucked"
}

func stringSliceToIntSlice(input []string) []int64 {
	output := make([]int64, len(input))

	for i, v := range input {
		output[i], _ = strconv.ParseInt(v, 10, 64)
	}

	return output
}
