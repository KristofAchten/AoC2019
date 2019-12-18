package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const defaultCodeSize = 7000

type intcodeState struct {
	program      []int64
	ptr          int64
	relativeBase int64
	input        []int64
}

func day5() {
	start := time.Now()

	input := strings.Split(string(getPuzzleInput("input/day5.txt")), ",")

	res1 := runUntilHalt(createDefaultIntcodeState(stringSliceToIntSlice(input), []int64{1}))
	res2 := runUntilHalt(createDefaultIntcodeState(stringSliceToIntSlice(input), []int64{5}))

	printPuzzleResult(5, int(res1), int(res2))

	fmt.Printf("DAY 5 STATS: Execution took %s\n\n", time.Since(start))
}

func runUntilHalt(state intcodeState, lastResult ...int64) int64 {
	result, halt, newState := runIntCode(state)
	if halt {
		return lastResult[0]
	}

	return runUntilHalt(newState, result)
}

func runIntCode(state intcodeState) (int64, bool, intcodeState) {
	var output int64

	code := make([]int64, defaultCodeSize)
	copy(code, state.program)

	pointer := state.ptr
	relativeBase := state.relativeBase
	input := state.input

	for {
		opcode := code[pointer] % 100
		strval := strconv.Itoa(int(code[pointer]))

		var modes string
		if len(strval) > 2 {
			modes = strconv.Itoa(int(code[pointer]))[0:(len(strval) - 2)]
		}

		switch opcode {
		case 1: // Addition
			res := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2])
			code[code[pointer+3]+b2i[offSet(modes, 3)]*relativeBase] = res[0] + res[1]
			pointer += 4
		case 2: // Multiplication
			res := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2])
			code[code[pointer+3]+b2i[offSet(modes, 3)]*relativeBase] = res[0] * res[1]
			pointer += 4
		case 3: // Take input
			code[code[pointer+1]+b2i[offSet(modes, 1)]*relativeBase] = input[0]
			input = input[1:]
			pointer += 2
		case 4: // Produce output
			res := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1])
			output = res[0]
			pointer += 2
			return output, false, intcodeState{code, pointer, relativeBase, input}
		case 5: // Jump if true (~0)
			res := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2])
			if res[0] != 0 {
				pointer = res[1]
			} else {
				pointer += 3
			}
		case 6: // Jump if false (=0)
			res := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2])
			if res[0] == 0 {
				pointer = res[1]
			} else {
				pointer += 3
			}
		case 7: // Less than
			res := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2])
			if res[0] < res[1] {
				code[code[pointer+3]+b2i[offSet(modes, 3)]*relativeBase] = 1
			} else {
				code[code[pointer+3]+b2i[offSet(modes, 3)]*relativeBase] = 0
			}
			pointer += 4
		case 8: // Equals
			res := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1], code[pointer+2])
			if res[0] == res[1] {
				code[code[pointer+3]+b2i[offSet(modes, 3)]*relativeBase] = 1
			} else {
				code[code[pointer+3]+b2i[offSet(modes, 3)]*relativeBase] = 0
			}
			pointer += 4
		case 9:
			res := getValsAccordingToModes(modes, relativeBase, code, code[pointer+1])
			relativeBase += res[0]
			pointer += 2
		case 99: // Halt
			return output, true, intcodeState{code, pointer, relativeBase, input}
		default:
			panic("Your intcode program uses unsupported opcodes!")
		}
	}
}

func offSet(modes string, val int) bool {
	return len(modes) > (val-1) && modes[len(modes)-val] == '2'
}

func getValsAccordingToModes(modes string, relativeBase int64, code []int64, vals ...int64) []int64 {
	var res []int64
	for i, v := range vals {
		if len(modes) > i {
			idx := len(modes) - 1 - i
			switch modes[idx] {
			case '1': // Immediate mode
				res = append(res, v)
				continue
			case '2': // Relative mode
				res = append(res, code[relativeBase+v])
				continue
			}
		}
		// Position mode (default)
		res = append(res, code[v])
	}
	return res
}

func createDefaultIntcodeState(code []int64, input []int64) intcodeState {
	return intcodeState{
		program:      code,
		ptr:          0,
		relativeBase: 0,
		input:        input,
	}
}
