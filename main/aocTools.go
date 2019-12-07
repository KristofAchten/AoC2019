package main

import (
	"io/ioutil"
	"os"
)

/**
Get the puzzle input from the file at a specified path
*/
func getPuzzleInput(path string) []byte {
	wd, _ := os.Getwd()
	file, err := os.Open(wd + "/" + path)
	if err != nil {
		panic(err)
	}

	conts, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return conts
}

/**
Absolute value for integers.
*/
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/**
Generate a slice of subsequent integers from f to t inclusive.
*/
func generateIntRange(f int, t int) []int {
	var fullRange []int
	for i := f; i < t; i++ {
		fullRange = append(fullRange, i)
	}

	return fullRange
}

/**
Reverse a slice of integers and return.
Brought to you by my functional programming nightmares
*/
func reverseIntSlice(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseIntSlice(input[1:]), input[0])
}

/**
Return the max value of a slice of positive integers
*/
func max(input []int) int {
	max := 0
	for _, v := range input {
		if v > max {
			max = v
		}
	}

	return max
}

/**
Return all permutations of a given int slice. Based on Heap's algorithm:
https://en.wikipedia.org/wiki/Heap%27s_algorithm#Details_of_the_algorithm
*/
func permutations(vals []int) [][]int {
	var allPermutations [][]int

	heapsAlgo(vals, len(vals), &allPermutations)
	return allPermutations
}

func heapsAlgo(vals []int, curLen int, results *[][]int) {
	if curLen == 1 {
		result := make([]int, len(vals))
		copy(result, vals)

		*results = append(*results, result)
	} else {
		for i := range vals {

			heapsAlgo(vals, curLen-1, results)

			if curLen%2 == 0 {
				vals[0], vals[curLen-1] = vals[curLen-1], vals[0]
			} else {
				vals[i], vals[curLen-1] = vals[curLen-1], vals[i]
			}
		}
	}
}
