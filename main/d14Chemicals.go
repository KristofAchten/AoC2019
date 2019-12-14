package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type formula struct {
	produces int
	reqs     []requirement
}

type requirement struct {
	quantity int
	id       string
}

var formulas = make(map[string]formula)
var inventory = make(map[string]int)

func day14() {
	start := time.Now()

	processInput(string(getPuzzleInput("input/day14.txt")))

	printPuzzleResult(14, howMuch("FUEL", 1), maxFuel(1000000000000, 0, 2000000))

	fmt.Printf("DAY 14 STATS: Execution took %s\n\n", time.Since(start))
}

func howMuch(want string, amount int) int {

	if want == "ORE" {
		return amount
	}

	have, _ := inventory[want]
	if have >= amount {
		inventory[want] = have - amount
		return 0
	} else if have != 0 {
		inventory[want] = 0
		amount -= have
	}

	form := formulas[want]
	multiplier := int(math.Ceil(float64(amount) / float64(form.produces)))
	var totalCost int
	for _, r := range form.reqs {
		get := r.quantity
		totalCost += howMuch(r.id, get*multiplier)
	}

	inventory[want] += (multiplier * form.produces) - amount
	return totalCost
}

func maxFuel(max int, curmin int, curmax int) int {
	if curmin == curmax-1 {
		return curmin
	}

	half := ((curmax - curmin) / 2) + curmin
	res := howMuch("FUEL", half)

	if res > max {
		return maxFuel(max, curmin, half)
	} else {
		return maxFuel(max, half, curmax)
	}
}

func processInput(input string) {
	parts := strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")

	for _, p := range parts {
		formParts := strings.Split(p, " => ")
		var reqStrins = strings.Split(formParts[0], ", ")

		var reqs []requirement
		for _, r := range reqStrins {
			vAndID := strings.Split(r, " ")
			v, _ := strconv.Atoi(vAndID[0])
			reqs = append(reqs, requirement{v, vAndID[1]})
		}

		vAndId := strings.Split(formParts[1], " ")
		v, _ := strconv.Atoi(vAndId[0])
		formulas[vAndId[1]] = formula{v, reqs}
	}
}
