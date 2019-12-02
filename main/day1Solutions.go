package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day1() {
	input := string(GetPuzzleInput("input/day1.txt"))

	res1 := calculateBaseFuel(input)
	res2 := calculateAllFuel(input)

	fmt.Println("Day 1: solution one is " + strconv.Itoa(res1))
	fmt.Println("Day 1: solution two is " + strconv.Itoa(res2))
}

func calculateBaseFuel(input string) int {

	parts := strings.Split(input, "\n")

	var x int
	for _, v := range parts {
		intval, _ := strconv.Atoi(v)
		x += (intval / 3) - 2
	}
	return x
}

func calculateAllFuel(input string) int {
	parts := strings.Split(input, "\n")

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
