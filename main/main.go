package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	day1()
	day2()
	day3()

	fmt.Println("**********************************************")
	fmt.Printf("ELAPSED TIME OVER ALL CHALLENGES: %s \n", time.Since(start))
	fmt.Println("**********************************************")
}
