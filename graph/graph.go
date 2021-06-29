package graph

import (
	"container/list"
)

func SetVerticesData(g Graph, data Data) {
	g.ForVertices(func(index int, vertex Vertex) { vertex.SetData(data) })
}

func SetEdgesData(g Graph, data Data) {
	g.ForEdges(func(from, to int, edge Edge) { edge.SetData(data) })
}

func clearVisited(g Graph) {
	g.ForVertices(func(index int, vertex Vertex) { vertex.setVisited(false) })
}

func ForPresentVertices(g Graph, vp VertexProcessor) {
	g.ForVertices(func(index int, vertex Vertex) {
		if vertex.IsPresent() {
			vp(index, vertex)
		}
	})
}

func ForConnectedEdges(g Graph, ep EdgeProcessor) {
	g.ForEdges(func(from, to int, edge Edge) {
		if edge.IsConnected() {
			ep(from, to, edge)
		}
	})
}

func Neighbours(g Graph, index int) []int {
	neighbours := make([]int, 0)
	g.ForNeighbours(index, func(i int, vertex Vertex) { neighbours = append(neighbours, i) })
	return neighbours
}

func DepthFirst(g Graph, start int, vp VertexProcessor, ep EdgeProcessor) {
	if vp == nil {
		vp = emptyVertexProcessor
	}
	if ep == nil {
		ep = emptyEdgeProcessor
	}
	clearVisited(g)
	depthFirst(g, start, vp, ep)
}

func depthFirst(g Graph, index int, vp VertexProcessor, ep EdgeProcessor) {
	vertex := g.GetVertex(index)
	vertex.setVisited(true)
	vp(index, vertex)
	for _, i := range Neighbours(g, index) {
		nextVertex := g.GetVertex(i)
		if !nextVertex.isVisited() {
			nextEdge := g.GetEdge(index, i)
			ep(index, i, nextEdge)
			depthFirst(g, i, vp, ep)
		}
	}
}

func BreadthFirst(g Graph, start int, vp VertexProcessor) {
	if vp == nil {
		vp = emptyVertexProcessor
	}
	clearVisited(g)
	breadthFirst(g, start, vp)
}

func breadthFirst(g Graph, start int, vp VertexProcessor) {
	nexts := list.New()
	nexts.PushBack(start)
	g.GetVertex(start).setVisited(true)
	for e := nexts.Front(); e != nil; e = nexts.Front() {
		index := e.Value.(int)
		vertex := g.GetVertex(index)
		nexts.Remove(e)
		vp(index, vertex)
		for _, i := range Neighbours(g, index) {
			if !g.GetVertex(i).isVisited() {
				g.GetVertex(i).setVisited(true)
				nexts.PushBack(i)
			}
		}
	}
}

func IsConnectedFrom(g Graph, start int) bool {
	count := 0
	counter := func(_ int, vertex Vertex) { count++ }
	DepthFirst(g, start, counter, nil)
	return count == g.Size()
}
