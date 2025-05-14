package main

import (
	"fmt"
	"math"
)

func findLowestCost(costs map[string]float64, seen map[string]bool) string {
	lowestCost := math.Inf(1)
	lowestCostNode := ""
	for node := range costs {
		cost := costs[node]
		if cost < lowestCost && !seen[node] {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode

}

func main() {
	graph := map[string]map[string]float64{
		"start": {"a": 5, "b": 2},
		"a":     {"d": 4, "c": 2},
		"b":     {"a": 8, "c": 7},
		"c":     {"fin": 1},
		"d":     {"c": 6, "fin": 3},
		"fin":   {},
	}
	costs := map[string]float64{
		"a":   5,
		"b":   2,
		"c":   math.Inf(1),
		"d":   math.Inf(1),
		"fin": math.Inf(1),
	}
	parents := map[string]string{
		"a":   "start",
		"b":   "start",
		"fin": "",
	}
	seen := map[string]bool{}

	node := findLowestCost(costs, seen)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for neighbor := range neighbors {
			newCost := cost + neighbors[neighbor]
			if costs[neighbor] > newCost {
				costs[neighbor] = newCost
				parents[neighbor] = node
			}
		}
		seen[node] = true
		node = findLowestCost(costs, seen)
	}

	path := []string{}
	current := "fin"
	for current != "" {
		path = append([]string{current}, path...)
		current = parents[current]
	}

	fmt.Println(path)
	fmt.Println(costs)
	fmt.Println(costs["fin"])

}
