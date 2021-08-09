package algorithms

import (
	"container/list"
	"math"

	"github.com/diefesson/ufc-ga/graph"
)

func calculateDepthsFrom(g *graph.DiGraph, depths []int, start int) {
	depth := 0
	down := func(_ graph.Graph, index int) {
		if depths[index] < depth {
			depths[index] = depth
		}
		depth++
	}
	up := func(_ graph.Graph, index int) {
		depth--
	}
	graph.DepthFirst(g, start, down, up, nil)
}

func InTopologicalOrder(g *graph.DiGraph, vp graph.VertexProcessor) {
	depths := make([]int, g.Capacity())
	visited := make([]bool, g.Capacity())
	nexts := list.New()

	for i := 0; i < len(depths); i++ {
		depths[i] = math.MinInt64
	}
	graph.ForRoots(g, func(_ graph.Graph, r int) {
		calculateDepthsFrom(g, depths, r)
		nexts.PushBack(r)
		visited[r] = true
	})

	for e := nexts.Front(); e != nil; e = nexts.Front() {
		v := e.Value.(int)
		vp(g, v)
		graph.ForNeighbours(g, v, func(_ graph.Graph, n int) {
			if !visited[n] && depths[n] == depths[v]+1 {
				nexts.PushBack(n)
				visited[n] = true
			}
		})
		nexts.Remove(e)
	}
}

func TopologicalOrdered(g *graph.DiGraph) []int {
	vertices := make([]int, 0, g.VertexCount())
	InTopologicalOrder(g, graph.CollectVerticesInto(&vertices))
	return vertices
}
