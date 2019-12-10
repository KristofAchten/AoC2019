package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type vec struct {
	x, y float64
}

func day10() {
	start := time.Now()

	input := string(getPuzzleInput("input/test.txt"))
	parts := strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")

	coord, best := findBestAsteroidparts(parts)
	fmt.Println(coord, best)

	fmt.Println("Day 10: solution one is " + "NOT IMPLEMENTED")
	fmt.Println("Day 10: solution two is " + "NOT IMPLEMENTED")

	fmt.Printf("DAY 10 STATS: Execution took %s\n\n", time.Since(start))
}

func getAsteroidCoordingates(parts []string) []coords {
	var asteroids []coords

	for y, v1 := range parts {
		for x, v2 := range strings.Split(v1, "") {
			if v2 == "#" {
				asteroids = append(asteroids, coords{x, y})
			}
		}
	}
	return asteroids
}

func findBestAsteroidparts(parts []string) (coords, int) {

	asteroids := getAsteroidCoordingates(parts)
	var curBest int
	var curBestVec coords

	for _, a1 := range asteroids {
		var unitvecs []vec
		for _, a2 := range asteroids {
			if a1 == a2 {
				continue
			}

			v := normalize(vec{float64(a2.x - a1.x), float64(a2.y - a1.y)})
			if !contains(unitvecs, v) {
				unitvecs = append(unitvecs, v)
			}
		}

		if len(unitvecs) > curBest {
			curBest = len(unitvecs)
			curBestVec = a1
		}
	}

	return curBestVec, curBest
}

func normalize(v vec) vec {
	nf := math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
	return vec{v.x / nf, v.y / nf}
}

func contains(vals []vec, val vec) bool {
	for _, v := range vals {
		if v == val {
			return true
		}
	}
	return false
}
