package main

import (
	"io/ioutil"
	"os"
)

/**
Map bools to ints
*/
var b2i = map[bool]int64{false: 0, true: 1}

const notImplemented = -1

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
func reverseIntSlice(input []int64) []int64 {
	if len(input) == 0 {
		return input
	}
	return append(reverseIntSlice(input[1:]), input[0])
}

/**
Return the max value of a slice of positive integers
*/
func max(input []int64) int64 {
	var max int64
	max = 0
	for _, v := range input {
		if v > max {
			max = v
		}
	}

	return max
}

/**
Return the max value of a slice of positive integers
*/
func maxInt(one int, two int) int {
	if one >= two {
		return one
	} else {
		return two
	}
}

/**
Return all permutations of a given int slice. Based on Heap's algorithm:
https://en.wikipedia.org/wiki/Heap%27s_algorithm#Details_of_the_algorithm
*/
func permutations(vals []int64) [][]int64 {
	var allPermutations [][]int64

	heapsAlgo(vals, len(vals), &allPermutations)
	return allPermutations
}

func heapsAlgo(vals []int64, curLen int, results *[][]int64) {
	if curLen == 1 {
		result := make([]int64, len(vals))
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

/**
Join all values of a slice of strings together
*/
func joinStringSlice(vals []string) string {
	var str string
	for _, v := range vals {
		str += v
	}

	return str
}

/**
Reverse a slice of strings and return.
Again, brought to you by my functional programming nightmares
*/
func reverseStringSlice(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseStringSlice(input[1:]), input[0])
}

/**
Check if a boolean slice contains only true-vals
*/
func allTrue(input []bool) bool {
	for _, v := range input {
		if !v {
			return false
		}
	}
	return true
}

// GCD The two methods below were gracefully copy/pasted from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
