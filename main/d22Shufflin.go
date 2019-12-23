package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func day22() {
	start := time.Now()

	input := string(getPuzzleInput("input/day22.txt"))
	parts := strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")

	printPuzzleResult(22, findCard(runCardTrick(parts, 10007), 2019), doHardMath(parts))

	fmt.Printf("DAY 22 STATS: Execution took %s\n\n", time.Since(start))
}

func runCardTrick(input []string, decksize int) []int {
	deck := createDeck(decksize)

	for _, v := range input {
		if strings.Contains(v, "deal into new stack") {
			deck = deal(deck)
		} else {
			parts := strings.Split(v, " ")
			val, _ := strconv.Atoi(parts[len(parts)-1])

			if strings.Contains(v, "deal with increment") {
				deck = dealInc(deck, val)
			} else if strings.Contains(v, "cut") {
				deck = cut(deck, val)
			} else {
				panic("Invalid command")
			}
		}
	}

	return deck
}

func doHardMath(input []string) int {
	/* Wouldn't have been able to solve this in years without help from reddit. Sorry.
	https://www.reddit.com/r/adventofcode/comments/ee0rqi/2019_day_22_solutions/fbnifwk/ */

	decksize := big.NewInt(119315717514047)
	numOfIters := big.NewInt(101741582076661)
	fA := bi(1)
	fB := bigDummy()

	for _, v := range input {
		if strings.Contains(v, "deal into new stack") {
			fA.Mul(fA, bi(-1))
			fB.Add(fB, fA)
		} else {
			parts := strings.Split(v, " ")
			val, _ := strconv.Atoi(parts[len(parts)-1])
			bigVal := bi(val)

			if strings.Contains(v, "deal with increment") {
				fA.Mul(fA, bigDummy().Exp(bigVal, bigDummy().Sub(decksize, bi(2)), decksize))
			} else if strings.Contains(v, "cut") {
				fB.Add(fB, bigDummy().Mul(bigVal, fA))
			} else {
				panic("Invalid command")
			}
		}
	}

	// This is essentially the formula as shown in the linked post
	s1 := bigDummy().Exp(fA, numOfIters, decksize)
	res := bigDummy().Mul(s1, bi(2020))

	s2 := bigDummy().Sub(bi(1), s1)
	s2.Mul(s2, modInv(fA, decksize))
	s2.Mul(s2, fB)

	res.Add(res, s2)
	res.Mod(res, decksize)

	return int(res.Int64())
}

func findCard(deck []int, card int) int {
	for c := range deck {
		if deck[c] == card {
			return c
		}
	}
	panic("Couldn't find the specified card.")
}

func createDeck(size int) []int {
	var deck []int
	for i := 0; i < size; i++ {
		deck = append(deck, i)
	}
	return deck
}

func deal(deck []int) []int {
	dealtDeck := make([]int, len(deck))
	for i := range deck {
		dealtDeck[i] = deck[len(deck)-1-i]
	}
	return dealtDeck
}

func cut(deck []int, loc int) []int {
	if loc >= 0 {
		return append(deck[loc:], deck[:loc]...)
	} else {
		return append(deck[len(deck)-abs(loc):], deck[:len(deck)-abs(loc)]...)
	}
}

func dealInc(deck []int, inc int) []int {
	dealtDeck := make([]int, len(deck))
	for c := range deck {
		dealtDeck[(c*inc)%len(deck)] = deck[c]
	}
	return dealtDeck
}
