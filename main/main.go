package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	old := os.Stdout
	wd, _ := os.Getwd()
	f, _ := os.Create(wd + "/output.txt")

	os.Stdout = f

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
	day13() // https://adventofcode.com/2019/day/13
	day14() // https://adventofcode.com/2019/day/14
	day15() // https://adventofcode.com/2019/day/15
	day16() // https://adventofcode.com/2019/day/16
	day17() // https://adventofcode.com/2019/day/17
	//day18() // https://adventofcode.com/2019/day/18 // Nope, I'm not solving a TSP for my own sanity.
	day19() // https://adventofcode.com/2019/day/19
	day20() // https://adventofcode.com/2019/day/20
	day21() // https://adventofcode.com/2019/day/21
	//day22() // https://adventofcode.com/2019/day/22
	day23() // https://adventofcode.com/2019/day/23

	fmt.Println("**********************************************")
	fmt.Printf("ELAPSED TIME OVER ALL CHALLENGES: %s \n", time.Since(start))
	fmt.Println("**********************************************")

	_ = f.Close()
	os.Stdout = old

	fmt.Println(string(getPuzzleInput("output.txt")))
}
