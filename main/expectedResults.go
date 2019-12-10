package main

import (
	"fmt"
	"strconv"
)

type resultSet struct {
	part1 int
	part2 int
}

func confirmPuzzleResult(day int, resultP1 int, resultP2 int) {
	var messages []string
	results := []resultSet{
		{3210097, 4812287},
		{8017076, 3146},
		{280, 10554},
		{1079, 699},
		{9431221, 1409363},
		{147807, 229},
		{567045, 39016654},
		{1965, 9223372036854775807},
		{4288078517, 69256},
		{247, 1919},
	}

	expectedResults := results[day-1]

	if resultP1 != expectedResults.part1 {
		messages = append(messages, "Day "+strconv.Itoa(day)+" part 1: unexpected results! "+
			"Expected "+strconv.Itoa(expectedResults.part1)+", actual "+strconv.Itoa(resultP1))
	}

	if resultP2 != expectedResults.part2 {
		messages = append(messages, "Day "+strconv.Itoa(day)+" part 2: unexpected results! "+
			"Expected "+strconv.Itoa(expectedResults.part2)+", actual "+strconv.Itoa(resultP2))
	}

	if len(messages) == 0 {
		fmt.Println("Test succesful: all results matched")
	} else {
		for _, message := range messages {
			fmt.Println(message)
		}
	}

}
