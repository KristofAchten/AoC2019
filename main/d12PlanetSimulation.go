package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type coord3D struct {
	x, y, z int
}

type vec3D struct {
	dx, dy, dz int
}

type planet struct {
	position coord3D
	velocity vec3D
}

func day12() {
	start := time.Now()

	input := parsePlanetInput(string(getPuzzleInput("/input/day12.txt")))

	res1 := simulate(input, 1000)
	res2 := findLoop(input)

	fmt.Println("Day 12: solution one is " + strconv.Itoa(res1))
	fmt.Println("Day 12: solution two is " + strconv.Itoa(res2))

	confirmPuzzleResult(12, res1, res2)

	fmt.Printf("DAY 12 STATS: Execution took %s\n\n", time.Since(start))
}

func parsePlanetInput(input string) []planet {
	input = strings.ReplaceAll(input, "<", "")
	input = strings.ReplaceAll(input, ">", "")
	input = strings.ReplaceAll(input, "=", "")
	input = strings.ReplaceAll(input, "x", "")
	input = strings.ReplaceAll(input, "y", "")
	input = strings.ReplaceAll(input, "z", "")
	input = strings.ReplaceAll(input, " ", "")

	input = strings.Replace(input, "\r\n", "\n", -1)
	params := strings.Split(input, "\n")

	var inputSet []planet

	for _, v := range params {
		finalParams := stringSliceToIntSlice(strings.Split(v, ","))
		inputSet = append(inputSet, planet{
			coord3D{int(finalParams[0]), int(finalParams[1]), int(finalParams[2])},
			vec3D{0, 0, 0},
		})
	}

	return inputSet
}

func simulate(input []planet, target int) int {

	run := 0
	for run < target {
		simulateStep(input)
		run++
	}

	total := 0
	for i := 0; i < len(input); i++ {
		position := input[i].position
		velocity := input[i].velocity

		potE := abs(position.x) + abs(position.y) + abs(position.z)
		kinE := abs(velocity.dx) + abs(velocity.dy) + abs(velocity.dz)

		total += potE * kinE
	}

	return total
}

func simulateStep(input []planet) {

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			p1 := input[i].position
			p2 := input[j].position

			if p1.x > p2.x {
				input[i].velocity.dx -= 1
				input[j].velocity.dx += 1
			} else if p1.x < p2.x {
				input[i].velocity.dx += 1
				input[j].velocity.dx -= 1
			}

			if p1.y > p2.y {
				input[i].velocity.dy -= 1
				input[j].velocity.dy += 1
			} else if p1.y < p2.y {
				input[i].velocity.dy += 1
				input[j].velocity.dy -= 1
			}

			if p1.z > p2.z {
				input[i].velocity.dz -= 1
				input[j].velocity.dz += 1
			} else if p1.z < p2.z {
				input[i].velocity.dz += 1
				input[j].velocity.dz -= 1
			}
		}
	}

	for i := 0; i < len(input); i++ {
		input[i].position = movePlanet(input[i])
	}
}

func movePlanet(p planet) coord3D {
	return coord3D{
		p.position.x + p.velocity.dx,
		p.position.y + p.velocity.dy,
		p.position.z + p.velocity.dz,
	}
}

func findLoop(input []planet) int {
	originalValues := make([]planet, len(input))
	doneSearching := make([]bool, 3)
	stepsRequired := make([]int, 3)

	copy(originalValues, input)

	steps := 0
	for !allTrue(doneSearching) {
		simulateStep(input)
		steps++

		for i := 0; i < 3; i++ {
			if doneSearching[i] {
				continue
			}

			done := true
			for i2, v := range input {
				if !equalOnAxis(v, originalValues[i2], i) {
					done = false
				}
			}

			if done {
				doneSearching[i] = true
				stepsRequired[i] = steps
			}
		}
	}

	return LCM(stepsRequired[0], stepsRequired[1], stepsRequired[2])
}

func equalOnAxis(p planet, p2 planet, axis int) bool {
	if axis == 0 { // X-Axis
		return p.position.x == p2.position.x && p.velocity.dx == p2.velocity.dx
	}

	if axis == 1 { // Y-Axis
		return p.position.y == p2.position.y && p.velocity.dy == p2.velocity.dy
	}

	if axis == 2 { // Z-Axis
		return p.position.z == p2.position.z && p.velocity.dz == p2.velocity.dz
	}

	panic("Invalid input for the axis parameter.")
}
