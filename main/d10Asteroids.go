package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

type vec struct {
	x, y float64
}

func day10() {
	start := time.Now()

	input := string(getPuzzleInput("input/day10.txt"))
	parts := strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")

	bestCoords, res1, info := findBestAsteroidparts(parts)
	vaporizedCoord := findVaporizedAsteroid(bestCoords, info, 200)
	res2 := vaporizedCoord.x*100 + vaporizedCoord.y

	fmt.Println("Day 10: solution one is " + strconv.Itoa(res1))
	fmt.Println("Day 10: solution two is " + strconv.Itoa(res2))

	confirmPuzzleResult(10, res1, res2)

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

func findBestAsteroidparts(parts []string) (coords, int, map[vec][]coords) {

	asteroids := getAsteroidCoordingates(parts)
	var curBest int
	var curBestVec coords
	var curBestInfo map[vec][]coords

	for _, a1 := range asteroids {
		var units2vec = make(map[vec][]coords)
		for _, a2 := range asteroids {
			if a1 == a2 {
				continue
			}

			v := normalize(vec{float64(a2.x - a1.x), float64(a2.y - a1.y)})
			units2vec[v] = append(units2vec[v], a2)
		}

		if len(units2vec) > curBest {
			curBest = len(units2vec)
			curBestVec = a1
			curBestInfo = units2vec
		}
	}

	return curBestVec, curBest, curBestInfo
}

func normalize(v vec) vec {
	nf := math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
	tv := vec{v.x / nf, v.y / nf}

	// icky rounding to 4 decimal places...
	return vec{math.Round(tv.x*10000) / 10000, math.Round(tv.y*10000) / 10000}
}

func findVaporizedAsteroid(center coords, info map[vec][]coords, target int) coords {
	sortedAngles := make(map[float64][]coords)

	for k, v := range info {
		// 1. Convert unit-vectors (= keys in info-map) to -Y-based angles and use this as index in a new map
		angle := negYBasedAngle(k)
		sortedAngles[angle] = v

		// 2. Sort per key based on distance to the object (furthest away = end of map)
		helper := func(i1 int, i2 int) bool {
			v1 := sortedAngles[angle][i1]
			v2 := sortedAngles[angle][i2]
			return closerThan(v1, v2, center)
		}
		sort.Slice(sortedAngles[angle], helper)
	}

	// 3. Sort the map (keys) to be used in the search later on.
	var keys []float64
	for k := range sortedAngles {
		keys = append(keys, k)
	}

	sort.Float64s(keys)

	// 4. Iterate over the keys of the map: if its value is not an empty slice -> one step closer to target. Remove 0-indexed val.
	var curIt int

	for {
		for _, v := range keys {
			if len(sortedAngles[v]) == 0 {
				continue
			}
			curIt++
			if curIt == target {
				return sortedAngles[v][0]
			} else {
				sortedAngles[v] = sortedAngles[v][1:]
			}

		}
	}
}

func negYBasedAngle(v vec) float64 {
	angle := math.Atan2(v.y, v.x) + (math.Pi / 2)
	if angle < 0 {
		angle += 2 * math.Pi
	}

	return angle
}

func closerThan(v1 coords, v2 coords, center coords) bool {
	d1 := math.Sqrt(math.Pow(float64(center.x-v1.x), 2) + math.Pow(float64(center.y-v1.y), 2))
	d2 := math.Sqrt(math.Pow(float64(center.x-v2.x), 2) + math.Pow(float64(center.y-v2.y), 2))

	return d1 < d2
}
