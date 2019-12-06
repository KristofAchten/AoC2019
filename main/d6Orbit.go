package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type tree struct {
	root node
}

type node struct {
	val      string
	children []node
}

func day6() {
	start := time.Now()

	input := string(getPuzzleInput("input/test.txt"))
	parts := strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")

	fmt.Println("Day 6: solution one is " + strconv.Itoa(determineAllOrbits(parts)))
	fmt.Println("Day 6: solution two is " + "NOT IMPLEMENTED")

	fmt.Printf("DAY 6 STATS: Execution took %s\n\n", time.Since(start))
}

func determineAllOrbits(parts []string) int {
	rootnode := node{"COM", []node{}}
	tree := tree{rootnode}

	for _, v := range parts {
		center := strings.Split(v, ")")[0]
		orbiter := strings.Split(v, ")")[1]

		newNode := node{orbiter, []node{}}
		parent := searchTree(tree, center)

	}

	return countAllOrbits(alldata)
}

func searchTree(tree tree, str string) *node {
	for _, v := range tree.root.children {
		if v.val == str {
			return &v
		}
		tree := tree{v}
		return searchTree(tree{v}, str)
	}
	return nil
}
