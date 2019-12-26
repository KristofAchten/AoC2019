package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func day25() {
	start := time.Now()

	/*input := stringSliceToIntSlice(strings.Split(string(getPuzzleInput("input/day25.txt")), ","))
	playGame(input) */
	fmt.Println("This is a game, didn't automate playing it. No output here...")
	fmt.Println("The necessary items are: ornament, easter egg, monolith and hypercube.")
	fmt.Println("Result: 1073815584")

	fmt.Printf("DAY 25 STATS: Execution took %s\n\n", time.Since(start))
}

func playGame(input []int64) {
	reader := bufio.NewReader(os.Stdin)
	state := createDefaultIntcodeState(input, []int64{})

	for {
		var output int64
		var halt bool
		output, halt, state = runIntCode(state)

		if halt {
			break
		}

		if output == ominousEmptyInputValue {
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')

			for i := range text {
				state.input = append(state.input, int64(text[i]))
			}
		} else {
			fmt.Printf("%c", output)
		}
	}

}
