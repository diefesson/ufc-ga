package algorithms

import (
	"math"

	"github.com/diefesson/ufc-ga/graph"
)

func findAugmentedPath(g graph.Graph, capacity [][]float64, flow [][]float64, source, sink int) []int {
	path := []int{}
	capacityFilter := func(_ graph.Graph, f, t int) bool {
		return capacity[f][t]-flow[f][t] > 0
	}
	err := graph.BreadthFirstSearch(g, source, sink, capacityFilter, graph.CollectVerticesInto(&path))
	if err != nil {
		return nil
	} else {
		return path
	}
}

func findBottleneck(path []int, capacity [][]float64, flow [][]float64) float64 {
	bottleneck := math.Inf(1)
	graph.ForPathEdges(path, func(f, t int) {
		remaining := capacity[f][t] - flow[f][t]
		if remaining < bottleneck {
			bottleneck = remaining
		}
	})
	return bottleneck
}

func updateFlow(path []int, flow [][]float64, addition float64) {
	graph.ForPathEdges(path, func(f, t int) {
		flow[f][t] += addition
	})
}

func EdmondsKarp(g graph.Graph, capacity [][]float64, source, sink int) (float64, [][]float64) {
	flow := make([][]float64, g.Capacity())
	for i := 0; i < g.Capacity(); i++ {
		flow[i] = make([]float64, g.Capacity())
	}
	for {
		path := findAugmentedPath(g, capacity, flow, source, sink)
		if path == nil {
			break
		}
		bottleneck := findBottleneck(path, capacity, flow)
		updateFlow(path, flow, bottleneck)
	}
	total := 0.0
	graph.ForNeighbours(g, source, func(_ graph.Graph, n int) {
		total += flow[source][n]
	})
	return total, flow
}
