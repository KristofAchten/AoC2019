package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	day4min = 245318
	day4max = 765747
)

func day4() {
	start := time.Now()

	candidates := getPwd()
	fmt.Println("Day 4: solution one is " + strconv.Itoa(len(candidates)))
	fmt.Println("Day 4: solution two is " + strconv.Itoa(len(oneDoubleWithoutSupergroup(candidates))))

	fmt.Printf("DAY 4 STATS: Execution took %s\n\n", time.Since(start))
}

func getPwd() []int {
	return oneDouble(monotonicRising(generateIntRange(day4min, day4max)))
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

func oneDouble(candidates []int) []int {
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

func oneDoubleWithoutSupergroup(candidates []int) []int {
	var newcandidates []int
	for _, v1 := range candidates {
		doubles := determineDoublesInString(strconv.Itoa(v1))

		var add bool
		for _, d := range doubles {
			c := strconv.Itoa(d)
			strval := strconv.Itoa(v1)
			add = true

			for i := 3; i < 6; i++ {
				repStr := strings.Repeat(c, i)
				if strings.Contains(strval, repStr) {
					add = false
					break
				}
			}

			if add {
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
