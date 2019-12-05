package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func day5() {
	start := time.Now()

	input := strings.Split(string(getPuzzleInput("input/day5.txt")), ",")
	o1 := runAdvancedIntcode(stringSliceToIntSlice(input), []int{1})
	o2 := runAdvancedIntcode(stringSliceToIntSlice(input), []int{5})

	fmt.Println("Day 5: solution one is " + strconv.Itoa(o1[len(o1)-1]))
	fmt.Println("Day 5: solution two is " + strconv.Itoa(o2[len(o2)-1]))

	fmt.Printf("DAY 5 STATS: Execution took %s\n\n", time.Since(start))
}

func runAdvancedIntcode(code []int, input []int) []int {
	var output []int
	pointer := 0

	for pointer < len(code) {

		opcode := code[pointer] % 100

		strval := strconv.Itoa(code[pointer])
		var modes = ""

		if len(strval) > 2 {
			modes = strconv.Itoa(code[pointer])[0:(len(strval) - 2)]
		}

		switch opcode {
		case 1: // Addition
			res := getValsAccordingToModes(modes, code, code[pointer+1], code[pointer+2])
			code[code[pointer+3]] = res[0] + res[1]
			pointer += 4
		case 2: // Multiplication
			res := getValsAccordingToModes(modes, code, code[pointer+1], code[pointer+2])
			code[code[pointer+3]] = res[0] * res[1]
			pointer += 4
		case 3: // Take input
			code[code[pointer+1]] = input[0]
			input = input[1:]
			pointer += 2
		case 4: // Store output
			res := getValsAccordingToModes(modes, code, code[pointer+1])
			output = append(output, res[0])
			pointer += 2
		case 5: // Jump if true (~0)
			res := getValsAccordingToModes(modes, code, code[pointer+1], code[pointer+2])
			if res[0] != 0 {
				pointer = res[1]
			} else {
				pointer += 3
			}
		case 6: // Jump if false (=0)
			res := getValsAccordingToModes(modes, code, code[pointer+1], code[pointer+2])
			if res[0] == 0 {
				pointer = res[1]
			} else {
				pointer += 3
			}
		case 7: // Less than
			res := getValsAccordingToModes(modes, code, code[pointer+1], code[pointer+2], code[pointer+3])
			if res[0] < res[1] {
				code[code[pointer+3]] = 1
			} else {
				code[code[pointer+3]] = 0
			}
			pointer += 4
		case 8: // Equals
			res := getValsAccordingToModes(modes, code, code[pointer+1], code[pointer+2], code[pointer+3])
			if res[0] == res[1] {
				code[code[pointer+3]] = 1
			} else {
				code[code[pointer+3]] = 0
			}
			pointer += 4
		default:
			pointer = 99999999
		}
	}
	return output
}

func getValsAccordingToModes(modes string, code []int, vals ...int) []int {
	var res []int
	for i, v := range vals {
		if len(modes) > i {
			if modes[(len(modes)-1)-i] == '1' {
				res = append(res, v)
				continue
			}
		}
		res = append(res, code[v])
	}
	return res
}
