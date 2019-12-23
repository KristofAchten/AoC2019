package main

import (
	"fmt"
	"strings"
	"time"
)

var natX int64
var natY int64
var natHistory []coords

func day23() {
	start := time.Now()

	input := stringSliceToIntSlice(strings.Split(string(getPuzzleInput("input/day23.txt")), ","))

	printPuzzleResult(23, setupNetwork(input, true), setupNetwork(input, false))

	fmt.Printf("DAY 23 STATS: Execution took %s\n\n", time.Since(start))
}

func setupNetwork(input []int64, returnOn255 bool) int {
	var computers []intcodeState
	for i := 0; i < 50; i++ {
		computers = append(computers, createDefaultIntcodeState(input, []int64{int64(i), -1}))
	}

	for {
		idle := true
		for c := range computers {
			if len(computers[c].input) < 2 {
				computers[c].input = append(computers[c].input, -1)
			} else {
				idle = false
			}
			output, _, newstate := runIntCode(computers[c])
			computers[c] = newstate

			if output == ominousEmptyInputValue {
				continue
			} else {
				idle = false
				var x, y int64
				x, _, newstate = runIntCode(newstate)
				y, _, newstate = runIntCode(newstate)
				computers[c] = newstate

				if output == 255 {
					if returnOn255 {
						return int(y)
					}
					natX = x
					natY = y
				} else {
					computers[int(output)].input = append(computers[int(output)].input, x, y)
				}
			}
		}
		if idle {
			computers[0].input = append(computers[0].input, natX, natY)
			if twice(natY) {
				return int(natY)
			}
			natHistory = append(natHistory, coords{int(natX), int(natY)})
		}
	}
}

func twice(i int64) bool {
	for _, v := range natHistory {
		if v.y == int(i) {
			return true
		}
	}
	return false
}
