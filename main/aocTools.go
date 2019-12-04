package main

import (
	"io/ioutil"
	"os"
)

/**
Get the puzzle input from the file at a specified path
*/
func GetPuzzleInput(path string) []byte {
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
func Abs(x int) int {
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
