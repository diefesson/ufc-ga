package algorithms

import (
	"ag/graph"
	"math"
)

const (
	UNMARKED = iota
	MARKED
	BLOCKED
)

type kruskal struct {
	graph             *graph.UniGraph
	groups            *graph.DisjointSet
	markedEdges       *graph.UniEdgeDataLayer
	calculateDistance graph.DistanceCalculator
	onFound           graph.EdgeProcessor
}

func makeKruskal(g *graph.UniGraph, cd graph.DistanceCalculator, of graph.EdgeProcessor) kruskal {
	if of == nil {
		of = graph.EmptyEdgeProcessor
	}
	markedEdges := graph.NewUniEdgeDataLayer(g.Capacity())
	g.ForEdges(func(_ graph.Graph, from, to int) { markedEdges.Set(from, to, UNMARKED) })
	return kruskal{
		graph:             g,
		groups:            graph.NewDisjointSet(g.Capacity()),
		markedEdges:       markedEdges,
		calculateDistance: cd,
		onFound:           of,
	}
}

func (k *kruskal) sweep() (int, int) {
	distance := math.MaxFloat64
	from, to := -1, -1
	s := func(_ graph.Graph, f, t int) {
		if k.markedEdges.Get(f, t) == UNMARKED {
			fr := k.groups.RepresentantOf(f)
			tr := k.groups.RepresentantOf(t)
			if fr == tr {
				k.markedEdges.Set(f, t, BLOCKED)
			} else if d := k.calculateDistance(k.graph, f, t); d < distance {
				distance = d
				from = f
				to = t
			}
		}
	}
	k.graph.ForEdges(graph.IfConnected(s))
	return from, to
}

func (k *kruskal) run() {
	for from, to := k.sweep(); from != -1; from, to = k.sweep() {
		k.groups.Join(from, to)
		k.markedEdges.Set(from, to, MARKED)
		k.onFound(k.graph, from, to)
	}
}

func Kruskal(g *graph.UniGraph, cd graph.DistanceCalculator, of graph.EdgeProcessor) *graph.UniEdgeDataLayer {
	k := makeKruskal(g, cd, of)
	k.run()
	return k.markedEdges
}
