package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	width  = 25
	height = 6
)

func day8() {
	start := time.Now()

	input := strings.Split(string(getPuzzleInput("input/day8.txt")), "")

	res1 := productOfOnesAndTwos(getLayers1D(input))
	finalLayer := joinLayers(getLayers1D(input))
	res2, _ := strconv.Atoi(joinStringSlice(finalLayer))

	fmt.Println("Day 8: solution one is " + strconv.Itoa(res1))
	fmt.Println("Day 8: solution two is: ")
	nicePrint(finalLayer)

	confirmPuzzleResult(8, res1, res2)

	fmt.Printf("DAY 8 STATS: Execution took %s\n\n", time.Since(start))
}

func getLayers1D(input []string) [][]string {
	size := width * height
	var layers [][]string
	for start := 0; start < len(input); start += size {
		layers = append(layers, input[start:start+size])
	}
	return layers
}

func productOfOnesAndTwos(layers [][]string) int {
	fewestZeros := 99999999
	var ones int
	var twos int
	for _, layer := range layers {
		zero, one, two := countOnesAndTwos(layer)
		if zero < fewestZeros {
			fewestZeros = zero
			ones = one
			twos = two
		}
	}
	return ones * twos
}

func countOnesAndTwos(layer []string) (int, int, int) {
	zeros := 0
	ones := 0
	twos := 0
	for _, s := range layer {
		switch s {
		case "0":
			zeros++
		case "1":
			ones++
		case "2":
			twos++
		default:
			continue
		}
	}

	return zeros, ones, twos
}

func joinLayers(layers [][]string) []string {
	joinedLayer := transparentLayer()

	for _, layer := range layers {
		for i, v := range layer {
			if v == "2" || joinedLayer[i] != "2" {
				continue
			} else {
				joinedLayer[i] = v
			}
		}
	}
	return joinedLayer
}

func nicePrint(layer []string) {
	for i, v := range layer {
		if v == "0" {
			layer[i] = " "
		} else {
			layer[i] = "X"
		}
	}

	for i := 0; i < width*height; i += width {
		fmt.Println(layer[i : i+width])
	}
}

func transparentLayer() []string {
	var transparentLayer []string
	for idx := 0; idx < width*height; idx++ {
		transparentLayer = append(transparentLayer, "2")
	}

	return transparentLayer
}
