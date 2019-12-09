package main

import (
	"fmt"
	"time"
)

func day9() {
	start := time.Now()

	//input := stringSliceToIntSlice(strings.Split(string(getPuzzleInput("input/day9.txt")), ","))

	fmt.Println(runUntilHalt(createDefaultIntcodeState([]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, []int{})))
	//fmt.Println(runUntilHalt(createDefaultIntcodeState(input, []int{1})))

	fmt.Println("Day 9: solution one is " + "NOT IMPLEMENTED")
	fmt.Println("Day 9: solution two is " + "NOT IMPLEMENTED")

	fmt.Printf("DAY 9 STATS: Execution took %s\n\n", time.Since(start))
}
