package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func day1() {
	start := time.Now()

	input := string(getPuzzleInput("input/day1.txt"))
	parts := strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")

	fmt.Println("Day 1: solution one is " + strconv.Itoa(calculateBaseFuel(parts)))
	fmt.Println("Day 1: solution two is " + strconv.Itoa(calculateAllFuel(parts)))

	fmt.Printf("DAY 1 STATS: Execution took %s\n\n", time.Since(start))
}

func calculateBaseFuel(parts []string) int {
	var x int
	for _, v := range parts {
		intval, _ := strconv.Atoi(v)
		x += (intval / 3) - 2
	}
	return x
}

func calculateAllFuel(parts []string) int {
	var total int
	for _, v := range parts {
		intval, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		fuel := (intval / 3) - 2
		total += fuel

		fuel = (fuel / 3) - 2
		for fuel > 0 {
			total += fuel
			fuel = (fuel / 3) - 2
		}
	}

	return total
}
