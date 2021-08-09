package algorithms

import (
	"math"

	"github.com/diefesson/ufc-ga/graph"
)

func findStartEnd(vertices []int, from, to int) (int, int) {
	var start, end int
	for i, v := range vertices {
		if from == v {
			start = i
		}
	}
	for i, v := range vertices {
		if to == v {
			end = i
		}
	}
	return start, end
}

func calculateDistances(
	g *graph.DiGraph,
	vertices []int,
	start,
	end int,
	dc graph.DistanceCalculator,
) []float64 {
	distances := make([]float64, g.Capacity())
	for i := 0; i < len(distances); i++ {
		distances[i] = math.Inf(1)
	}
	distances[start] = 0.0
	for i := start; i < end; i++ {
		v := vertices[i]
		baseDistance := distances[v]
		graph.ForNeighbours(g, v, func(g graph.Graph, n int) {
			currentDistance := distances[n]
			calculatedDistance := baseDistance + dc(g, v, n)
			if calculatedDistance < currentDistance {
				distances[n] = calculatedDistance
			}
		})
	}
	return distances
}

func reverse(path []int) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}

func findPath(g *graph.DiGraph, distances []float64, from, to int) []int {
	path := make([]int, 0)
	for i := to; i != from; {
		path = append(path, i)
		next, distance := -1, math.Inf(1)
		graph.ForEdgesTo(g, i, graph.IfConnected(graph.SelectFrom(func(g graph.Graph, from int) {
			if distances[from] < distance {
				next = from
				distance = distances[from]
			}
		})))
		i = next
	}
	path = append(path, from)
	reverse(path)
	return path
}

func DAGSP(g *graph.DiGraph, from, to int, dc graph.DistanceCalculator) (float64, []int) {
	vertices := TopologicalOrdered(g)
	start, end := findStartEnd(vertices, from, to)
	if end < start {
		panic("Cannot calculate path because topologically 'to' is before 'from'")
	}
	distances := calculateDistances(g, vertices, start, end, dc)
	if math.IsInf(distances[to], 1) {
		panic("There's no valid path to the given vertices")
	}
	return distances[to], findPath(g, distances, from, to)
}
