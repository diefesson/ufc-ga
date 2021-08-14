package algorithms

import (
	"errors"
	"math"

	"github.com/diefesson/ufc-ga/graph"
)

func BellmanFord(g *graph.DiGraph, from, to int, dc graph.DistanceCalculator, vp graph.VertexProcessor) (float64, error) {
	successors := make([]int, g.Capacity())
	distances := make([]float64, g.Capacity())
	for i := 0; i < len(distances); i++ {
		successors[i] = -1
		distances[i] = math.Inf(1)
	}
	distances[to] = 0

	for i := 0; i < g.VertexCount(); i++ {
		g.ForEdges(graph.IfConnected(func(g graph.Graph, f, t int) {
			newDistance := distances[t] + dc(g, f, t)
			if newDistance < distances[f] {
				distances[f] = newDistance
				successors[f] = t
			}
		}))
	}

	if successors[from] == -1 {
		return -1, errors.New("could not find shortest path")
	}
	for t := 0; t < g.Capacity(); t++ {
		for f := 0; f < g.Capacity(); f++ {
			if g.IsConnected(f, t) && distances[t]+dc(g, f, t) < distances[f] {
				return -1, errors.New("graph contains negative weight")
			}
		}
	}

	for v := from; v != to; v = successors[v] {
		vp(g, v)
	}
	vp(g, to)
	return distances[from], nil
}
