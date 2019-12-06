package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day1() // https://adventofcode.com/2019/day/1
	day2() // https://adventofcode.com/2019/day/2
	day3() // https://adventofcode.com/2019/day/3
	day4() // https://adventofcode.com/2019/day/4
	day5() // https://adventofcode.com/2019/day/5
	day6() // https://adventofcode.com/2019/day/6
	day7() // https://adventofcode.com/2019/day/7

	fmt.Println("**********************************************")
	fmt.Printf("ELAPSED TIME OVER ALL CHALLENGES: %s \n", time.Since(start))
	fmt.Println("**********************************************")
}
