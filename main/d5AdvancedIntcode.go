package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type intcodeState struct {
	program      []int
	ptr          int
	relativeBase int
	input        []int
}

func day5() {
	start := time.Now()

	input := strings.Split(string(getPuzzleInput("input/day5.txt")), ",")

	res1 := runUntilHalt(createDefaultIntcodeState(stringSliceToIntSlice(input), []int{1}))
	res2 := runUntilHalt(createDefaultIntcodeState(stringSliceToIntSlice(input), []int{5}))

	fmt.Println("Day 5: solution one is " + strconv.Itoa(res1))
	fmt.Println("Day 5: solution two is " + strconv.Itoa(res2))

	confirmPuzzleResult(5, res1, res2)

	fmt.Printf("DAY 5 STATS: Execution took %s\n\n", time.Since(start))
}

func runUntilHalt(state intcodeState, lastResult ...int) int {
	result, halt, newState := runIntCode(state)
	fmt.Println(result)
	if halt {
		return lastResult[0]
	}
	return runUntilHalt(newState, result)
}

func runIntCode(state intcodeState) (int, bool, intcodeState) {
	var output int

	code := state.program
	pointer := state.ptr
	relativeBase := state.relativeBase
	input := state.input

	for pointer < len(code) {

		opcode := code[pointer] % 100

		strval := strconv.Itoa(code[pointer])
		var modes = ""

		if len(strval) > 2 {
			modes = strconv.Itoa(code[pointer])[0:(len(strval) - 2)]
		}

		switch opcode {
		case 1: // Addition
			res, newCode := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2])
			code = newCode
			code = growCode(code, code[pointer+3]-len(code)+1)
			code[code[pointer+3]] = res[0] + res[1]
			pointer += 4
		case 2: // Multiplication
			res, newCode := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2])
			code = newCode
			code = growCode(code, code[pointer+3]-len(code)+1)
			code[code[pointer+3]] = res[0] * res[1]
			pointer += 4
		case 3: // Take input
			code = growCode(code, code[pointer+1]-len(code)+1)
			code[code[pointer+1]] = input[0]
			input = input[1:]
			pointer += 2
		case 4: // Store output
			res, newCode := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1])
			code = newCode
			output = res[0]
			pointer += 2
			return output, false, intcodeState{code, pointer, relativeBase, input}
		case 5: // Jump if true (~0)
			res, newCode := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2])
			code = newCode
			if res[0] != 0 {
				pointer = res[1]
			} else {
				pointer += 3
			}
		case 6: // Jump if false (=0)
			res, newCode := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2])
			code = newCode
			if res[0] == 0 {
				pointer = res[1]
			} else {
				pointer += 3
			}
		case 7: // Less than
			res, newCode := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2], code[pointer+3])
			code = newCode
			code = growCode(code, code[pointer+3]-len(code)+1)
			if res[0] < res[1] {
				code[code[pointer+3]] = 1
			} else {
				code[code[pointer+3]] = 0
			}
			pointer += 4
		case 8: // Equals
			res, newCode := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2], code[pointer+3])
			code = newCode
			code = growCode(code, code[pointer+3]-len(code)+1)
			if res[0] == res[1] {
				code[code[pointer+3]] = 1
			} else {
				code[code[pointer+3]] = 0
			}
			pointer += 4
		case 9:
			res, newCode := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1])
			code = newCode
			relativeBase += res[0]
			pointer += 2
		case 99: // Halt
			return output, true, intcodeState{code, pointer, relativeBase, input}
		default:
			pointer = 99999999
		}
	}
	panic("Your intcode sucks.")
}

func getValsAccordingToModes(modes string, relativeBase int, code []int, vals ...int) ([]int, []int) {
	var res []int
	for i, v := range vals {
		if len(modes) > i {
			idx := len(modes) - 1 - i
			switch modes[idx] {
			case '1':
				res = append(res, v)
				continue
			case '2':
				if relativeBase >= len(code) {
					code = growCode(code, relativeBase-len(code)+1)
				}
				res = append(res, code[relativeBase])
				continue
			}
		}
		if v >= len(code) {
			code = growCode(code, v-len(code)+1)
		}
		res = append(res, code[v])
	}
	return res, code
}

func createDefaultIntcodeState(code []int, input []int) intcodeState {
	return intcodeState{
		program:      code,
		ptr:          0,
		relativeBase: 0,
		input:        input,
	}
}

func growCode(code []int, with int) []int {
	if with <= 0 {
		return code
	}
	newCode := make([]int, len(code))
	copy(newCode, code)

	for i := 0; i < with; i++ {
		newCode = append(newCode, 0)
	}

	return newCode
}
