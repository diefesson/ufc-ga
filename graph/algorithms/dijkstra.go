package algorithms

import (
	"errors"
	"math"

	"github.com/diefesson/ufc-ga/graph"
)

func popNearest(vertices *[]int, distances []float64) int {
	index := -1
	minDinstance := math.Inf(1)
	for i, v := range *vertices {
		if distances[v] < minDinstance {
			index = i
			minDinstance = distances[v]
		}
	}
	if index == -1 {
		return -1
	} else {
		vertex := (*vertices)[index]
		(*vertices)[index] = (*vertices)[len(*vertices)-1]
		*vertices = (*vertices)[:len(*vertices)-1]
		return vertex
	}
}

func Dijkstra(g graph.Graph, from, to int, dc graph.DistanceCalculator, vp graph.VertexProcessor) (float64, error) {
	successors := make([]int, g.Capacity())
	distances := make([]float64, g.Capacity())
	unvisited := []int{}
	for i := 0; i < g.Capacity(); i++ {
		successors[i] = -1
		distances[i] = math.Inf(1)
		if g.IsPresent(i) {
			unvisited = append(unvisited, i)
		}
	}
	distances[to] = 0

	for v := popNearest(&unvisited, distances); v != -1; v = popNearest(&unvisited, distances) {
		graph.ForEdgesTo(g, v, graph.IfConnected(func(_ graph.Graph, f, t int) {
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

	for v := from; v != to; v = successors[v] {
		vp(g, v)
	}
	vp(g, to)

	return distances[from], nil
}
