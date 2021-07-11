package graph

import (
	"math"
)

const (
	UNMARKED = iota
	MARKED
	BLOCKED
)

type kruskal struct {
	graph             *UniGraph
	disjointSet       *DisjointSet
	edgeLayer         *UniEdgeDataLayer
	calculateDistance DistanceCalculator
	onFound           EdgeProcessor
}

func makeKruskal(graph *UniGraph, cd DistanceCalculator, of EdgeProcessor) kruskal {
	if of == nil {
		of = emptyEdgeProcessor
	}
	edgeLayer := NewUniEdgeDataLayer(graph.Capacity())
	graph.ForEdges(func(_ Graph, from, to int) { edgeLayer.Set(from, to, UNMARKED) })
	return kruskal{
		graph:             graph,
		disjointSet:       NewDisjointSet(graph.Capacity()),
		edgeLayer:         edgeLayer,
		calculateDistance: cd,
		onFound:           of,
	}
}

func (k *kruskal) sweep() (int, int) {
	distance := math.MaxFloat64
	from, to := -1, -1
	k.graph.ForEdges(func(_ Graph, f, t int) {
		if k.edgeLayer.Get(f, t) == UNMARKED {
			fr := k.disjointSet.RepresentantOf(f)
			tr := k.disjointSet.RepresentantOf(t)
			if fr == tr {
				k.edgeLayer.Set(f, t, BLOCKED)
			} else if d := k.calculateDistance(k.graph, f, t); d < distance {
				distance = d
				from = f
				to = t
			}
		}
	})
	return from, to
}

func (k *kruskal) run() {
	for from, to := k.sweep(); from != -1; from, to = k.sweep() {
		k.disjointSet.Join(from, to)
		k.edgeLayer.Set(from, to, MARKED)
		k.onFound(k.graph, from, to)
	}
}

func KruskalAlgorithm(g *UniGraph, cd DistanceCalculator, of EdgeProcessor) (*DisjointSet, *UniEdgeDataLayer) {
	kruskal := makeKruskal(g, cd, of)
	kruskal.run()
	return kruskal.disjointSet, kruskal.edgeLayer
}
