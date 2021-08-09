package algorithms

import (
	"math"

	"github.com/diefesson/ufc-ga/graph"
)

func BellmanFord(g *graph.DiGraph, from, to int, dc graph.DistanceCalculator) (float64, []int) {
	path := make([]int, 0)
	predecessors := make([]int, g.Capacity())
	for i := 0; i < len(predecessors); i++ {
		predecessors[i] = -1
	}
	distances := make([]float64, g.Capacity())
	for i := 0; i < len(distances); i++ {
		distances[i] = math.Inf(1)
	}
	distances[from] = 0

	for i := 0; i < g.VertexCount(); i++ {
		g.ForEdges(graph.IfConnected(func(_ graph.Graph, from, to int) {
			baseDistance := distances[from]
			currentDistance := distances[to]
			newDistance := baseDistance + dc(g, from, to)
			if newDistance < currentDistance {
				distances[to] = newDistance
				predecessors[to] = from
			}
		}))
	}

	for v := to; v != -1; v = predecessors[v] {
		path = append(path, v)
	}
	reverse(path)
	return distances[to], path
}
