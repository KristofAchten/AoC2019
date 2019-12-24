package main

import (
	"fmt"
	"strconv"
)

type resultSet struct {
	part1 int
	part2 int
}

var results = []resultSet{
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
	{1909, 249},
	{11384, 452582583272768},
	{462, 23981},
	{907302, 1670299},
	{412, 418},
	{59281788, 96062868},
	{3448, 762405},
	{notImplemented, notImplemented},
	{110, 17302065},
	{526, 6292},
	{19349939, 1142412777},
	{4096, 78613970589919},
	{17714, 10982},
	{19516944, notImplemented},
}

func printPuzzleResult(day int, resultP1 int, resultP2 int) {

	fmt.Println("Day", day, "- solution one is", resultP1)
	fmt.Println("Day", day, "- solution two is", resultP2)

	var messages []string

	if day > len(results) {
		fmt.Println("Not all solutions implemented. Could not confirm results!")
		return
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
