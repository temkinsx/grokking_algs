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
		"start": {"a": 10},
		"a":     {"c": 20},
		"b":     {"a": 1},
		"c":     {"b": 1, "fin": 30},
		"fin":   {},
	}
	costs := map[string]float64{
		"a":   10,
		"b":   math.Inf(1),
		"c":   math.Inf(1),
		"fin": math.Inf(1),
	}
	parents := map[string]string{
		"a":   "start",
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
