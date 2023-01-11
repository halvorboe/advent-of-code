package day16

import (
	"sort"
)

func buildGraph(valves []Valve) []Node {
	// sort order
	sort.Slice(valves, func(i, j int) bool {
		if valves[i].From == "AA" {
			return true
		} else if valves[j].From == "AA" {
			return false
		} else {
			return valves[i].Flow > valves[j].Flow
		}
	})

	// generate mapping
	mapping := make(map[string]int)
	valvesWithFlow := 0
	for i, valve := range valves {
		mapping[valve.From] = i
		if valve.Flow > 0 || valve.From == "AA" {
			valvesWithFlow += 1
		}
	}

	// convert to integers
	edges := make([][]int, len(valves))
	for i, valve := range valves {
		edges[i] = make([]int, 0)
		for _, to := range valve.To {
			edges[i] = append(edges[i], mapping[to])
		}
	}

	graph := floydWarshall(edges)

	// exclude 0 valves
	nodes := make([]Node, 0)
	for i := 0; i < valvesWithFlow; i++ {
		if valves[i].Flow == 0 && valves[i].From != "AA" {
			panic("should not happen")
		}
		nodes = append(nodes, Node{
			id:        valves[i].From,
			distances: graph[i][:valvesWithFlow],
			flow:      valves[i].Flow,
		})
	}

	return nodes
}

func floydWarshall(nodes [][]int) [][]int {
	dist := make([][]int, len(nodes))
	for i := range dist {
		di := make([]int, len(nodes))
		for j := range di {
			di[j] = INF
		}
		di[i] = 0
		dist[i] = di
	}
	for u, node := range nodes {
		for _, v := range node {
			dist[u][v] = 1
		}
	}
	for k, dk := range dist {
		for _, di := range dist {
			for j, dij := range di {
				if d := di[k] + dk[j]; dij > d {
					di[j] = d
				}
			}
		}
	}
	// open the valve, we would never move to a valve without
	// opening it
	for i := range dist {
		for j := range dist[i] {
			dist[i][j] += 1
		}
	}
	return dist
}
