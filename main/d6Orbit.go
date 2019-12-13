package main

import (
	"fmt"
	"strings"
	"time"
)

type node struct {
	val      string
	children []string
	parent   string
}

func day6() {
	start := time.Now()

	input := string(getPuzzleInput("input/day6.txt"))
	parts := strings.Split(strings.Replace(input, "\r\n", "\n", -1), "\n")

	pseudoTree := createNodeMap(parts)

	res1 := determineAllOrbits(pseudoTree)
	res2 := shortestPath(pseudoTree, "YOU", "SAN") - 2

	printPuzzleResult(6, res1, res2)

	fmt.Printf("DAY 6 STATS: Execution took %s\n\n", time.Since(start))
}

func createNodeMap(parts []string) map[string]node {
	pseudoTree := make(map[string]node)
	pseudoTree["COM"] = node{"COM", []string{}, ""}

	for _, v := range parts {
		center := strings.Split(v, ")")[0]
		orbiter := strings.Split(v, ")")[1]

		var orbiternode node
		var centernode node

		// Orbiter already exists in the list (=> has been added previously): correctly set its parent to the center-node.
		// Else => create "empty" node with parent set correctly
		if on, onExists := pseudoTree[orbiter]; onExists {
			orbiternode = on
			orbiternode.parent = center
		} else {
			orbiternode = node{orbiter, []string{}, center}
		}

		// Center mass already exists in the list (=> has been added previously): add the orbiter to its list of children
		// Else => create "empty" node with children set correctly
		if cn, cnExists := pseudoTree[center]; cnExists {
			centernode = cn
			centernode.children = append(centernode.children, orbiternode.val)
		} else {
			centernode = node{center, []string{orbiter}, ""}
		}

		// Update the list, because fuck golang referencing
		pseudoTree[orbiter] = orbiternode
		pseudoTree[center] = centernode
	}
	return pseudoTree
}

func determineAllOrbits(pseudoTree map[string]node) int {
	return countAllOrbits("COM", pseudoTree, 0)
}

func countAllOrbits(node string, nodes map[string]node, depth int) int {
	curNode := nodes[node]

	if len(curNode.children) == 0 {
		// Leafnode: return the depth (= length of it's path to the root = # of direct and indirect orbits)
		return depth
	} else {
		// Internal node: recursively sum up depth of children and add the current depth to it (to include the weight of this node)
		total := depth
		for _, v := range curNode.children {
			total += countAllOrbits(v, nodes, depth+1)
		}
		return total
	}
}

func shortestPath(pseudoTree map[string]node, source string, target string) int {
	// This really just is slightly modified Dijkstra if I remember correctly

	dists := make(map[string]int)
	visited := make(map[string]bool)

	dists[source] = 0
	var cur string

	// This is necessary because delete(map, key) doesn't actually delete anything, it just.. invalidates?
	for len(visited) < len(pseudoTree) {
		cur = getUnvisitedNodeWithLeastDistance(dists, visited)

		if cur == target {
			break
		}

		// Check all links (= children + parent) and update distances
		for _, v := range append(pseudoTree[cur].children, pseudoTree[cur].parent) {
			newDist := dists[cur] + 1
			existingDist, exists := dists[v]

			if !exists || newDist < existingDist {
				dists[v] = newDist
			}
		}
	}

	return dists[target]
}

func getUnvisitedNodeWithLeastDistance(dists map[string]int, visited map[string]bool) string {
	minDist := 9999999
	var minDistNode string

	for k, v := range dists {
		_, ok := visited[k]

		if !ok && v < minDist {
			minDist = v
			minDistNode = k
		}
	}

	visited[minDistNode] = true
	return minDistNode
}
