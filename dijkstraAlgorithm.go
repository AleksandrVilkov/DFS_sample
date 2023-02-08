package main

import (
	"fmt"
	"math"
)

// Dijkstra Algorithm in Go

// createGraph returns a graph represented as a map of maps
func createGraph() map[string]map[string]int {
	graph := make(map[string]map[string]int)
	graph["start"] = map[string]int{"a": 6, "b": 2}
	graph["a"] = map[string]int{"fin": 1}
	graph["b"] = map[string]int{"a": 3, "fin": 5}
	graph["fin"] = map[string]int{}
	return graph
}

// createCosts returns a costs table represented as a map
func createCosts() map[string]int {
	costs := make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = math.MaxInt32
	return costs
}

// createParents returns a parents table represented as a map
func createParents() map[string]string {
	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""
	return parents
}

// Dijkstra finds the shortest path from the start node to the finish node
func Dijkstra(graph map[string]map[string]int, costs map[string]int, parents map[string]string) ([]string, int) {
	processed := []string{}
	node := findLowestCostNode(costs)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for n := range neighbors {
			newCost := cost + neighbors[n]
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed = append(processed, node)
		node = findLowestCostNode(costs)
	}
	return processed, costs["fin"]
}

// findLowestCostNode finds the node with the lowest cost
func findLowestCostNode(costs map[string]int) string {
	lowestCost := math.MaxInt32
	lowestCostNode := ""
	for node := range costs {
		cost := costs[node]
		if cost < lowestCost && node != "fin" {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func main() {
	graph := createGraph()
	costs := createCosts()
	parents := createParents()
	processed, cost := Dijkstra(graph, costs, parents)
	fmt.Println("Shortest path is", processed)
	fmt.Println("Cost is", cost)
}
