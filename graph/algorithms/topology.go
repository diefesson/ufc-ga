package algorithms

import (
	"container/list"

	"github.com/diefesson/ufc-ga/graph"
)

const (
	INCLUDED = iota
	EXCLUDED
)

func InTopologicalOrder(g *graph.DiGraph, vp graph.VertexProcessor) {
	c := g.Clone()
	nexts := list.New()

	graph.ForRoots(c, func(_ graph.Graph, i int) {
		nexts.PushBack(i)
	})

	for e := nexts.Front(); e != nil; e = nexts.Front() {
		n := e.Value.(int)
		nexts.Remove(e)
		neighbours := graph.Neighbours(c, n)
		c.Remove(n)
		vp(g, n)
		for _, i := range neighbours {
			if graph.IsRoot(c, i) {
				nexts.PushBack(i)
			}
		}
	}
}

func TopologicalOrdered(g *graph.DiGraph) []int {
	vertices := make([]int, 0, g.VertexCount())
	InTopologicalOrder(g, graph.CollectVerticesInto(&vertices))
	return vertices
}
