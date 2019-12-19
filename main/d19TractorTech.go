package main

import (
	"fmt"
	"strings"
	"time"
)

func day19() {
	start := time.Now()

	input := stringSliceToIntSlice(strings.Split(string(getPuzzleInput("input/day19.txt")), ","))

	// The startAt coords have been optimized to reduce execution time. Set them to {0, 0} for full generation. This takes a few minutes
	printPuzzleResult(19,
		mapBeam(input, 50, -1, coords{0, 0}),
		mapBeam(input, 2200, 100, coords{1725, 2060}))
	fmt.Printf("DAY 19 STATS: Execution took %s\n\n", time.Since(start))
}

func mapBeam(input []int64, beamLen int, surface int, startAt coords) int {
	beamMap := make([][]int, beamLen)
	for i := range beamMap {
		beamMap[i] = make([]int, beamLen)
	}

	count := 0
	state := createDefaultIntcodeState(input, []int64{})
	for i := startAt.x; i < beamLen; i++ {
		var width int
		for j := startAt.y; j < beamLen; j++ {
			state.input = []int64{int64(i), int64(j)}
			out := runUntilHalt(state)
			beamMap[i][j] = int(out)
			if out == 0 {
				width = 0
			} else {
				width++
				count++
				if surface > 0 && width == surface {
					if i-surface+1 >= 0 && beamMap[i-surface+1][j] == 1 {
						return ((i - surface + 1) * 10000) + (j - surface + 1)
					}
				}
			}
		}
	}

	return count
}
