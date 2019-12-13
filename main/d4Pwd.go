package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	day4min = 245318
	day4max = 765747
)

func day4() {
	start := time.Now()

	candidates := getPwd()

	res1 := len(candidates)
	res2 := len(atLeastOneDoubleWithoutSupergroup(candidates))

	printPuzzleResult(4, res1, res2)

	fmt.Printf("DAY 4 STATS: Execution took %s\n\n", time.Since(start))
}

func getPwd() []int {
	return atLeastOneDouble(monotonicRising(generateIntRange(day4min, day4max)))
}

func monotonicRising(candidates []int) []int {
	var newcandidates []int
	for _, v := range candidates {
		str := strconv.Itoa(v)
		var add = true

		for i := 0; i < len(str)-1; i++ {
			if str[i] > str[i+1] {
				add = false
				break
			}
		}

		if add {
			newcandidates = append(newcandidates, v)
		}
	}
	return newcandidates
}

func atLeastOneDouble(candidates []int) []int {
	var newcandidates []int
	for _, v := range candidates {
		str := strconv.Itoa(v)

		for i := 0; i < len(str)-1; i++ {
			if str[i] == str[i+1] {
				newcandidates = append(newcandidates, v)
				break
			}
		}
	}
	return newcandidates
}

func atLeastOneDoubleWithoutSupergroup(candidates []int) []int {
	var newcandidates []int
	for _, v1 := range candidates {
		doubles := determineDoublesInString(strconv.Itoa(v1))

		for _, d := range doubles {
			if match, _ := regexp.MatchString("["+strconv.Itoa(d)+"]{3}", strconv.Itoa(v1)); !match {
				newcandidates = append(newcandidates, v1)
				break
			}
		}
	}
	return newcandidates
}

func determineDoublesInString(str string) []int {
	var doubles []int

	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			iv, _ := strconv.Atoi(string(str[i]))
			doubles = append(doubles, iv)
		}
	}

	return doubles
}
