package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func day9() {
	start := time.Now()

	input := stringSliceToIntSlice(strings.Split(string(getPuzzleInput("input/day9.txt")), ","))

	res1 := int(runUntilHalt(createDefaultIntcodeState(input, []int64{1})))
	res2 := int(runUntilHalt(createDefaultIntcodeState(input, []int64{2})))

	fmt.Println("Day 9: solution one is " + strconv.Itoa(res1))
	fmt.Println("Day 9: solution two is " + strconv.Itoa(res2))

	confirmPuzzleResult(9, res1, res2)

	fmt.Printf("DAY 9 STATS: Execution took %s\n\n", time.Since(start))
}
