package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day1()  // https://adventofcode.com/2019/day/1
	day2()  // https://adventofcode.com/2019/day/2
	day3()  // https://adventofcode.com/2019/day/3
	day4()  // https://adventofcode.com/2019/day/4
	day5()  // https://adventofcode.com/2019/day/5
	day6()  // https://adventofcode.com/2019/day/6
	day7()  // https://adventofcode.com/2019/day/7
	day8()  // https://adventofcode.com/2019/day/8
	day9()  // https://adventofcode.com/2019/day/9
	day10() // https://adventofcode.com/2019/day/10
	day11() // https://adventofcode.com/2019/day/11
	day12() // https://adventofcode.com/2019/day/12

	fmt.Println("**********************************************")
	fmt.Printf("ELAPSED TIME OVER ALL CHALLENGES: %s \n", time.Since(start))
	fmt.Println("**********************************************")
}
