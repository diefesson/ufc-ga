package graph

import (
	"container/list"
)

type Graph interface {
	Capacity() int
	VertexCount() int
	EdgeCount() int
	Add(index int)
	IsPresent(index int) bool
	Remove(index int)
	Connect(from, to int)
	IsConnected(from, to int) bool
	Disconnect(from, to int)
	ForVertices(vp VertexProcessor)
	ForEdges(ep EdgeProcessor)
	CreateVertexDataLayer(key string) *UDDataLayer
	GetVertexDataLayer(key string) *UDDataLayer
	RemoveVertexDataLayer(key string)
	CreateEdgeDataLayer(key string) BDDataLayer
	GetEdgeDataLayer(key string) BDDataLayer
	RemoveEdgeDataLayer(key string)
	VerifyConnected() bool
	Clone() Graph
}

func ForEdgesFrom(g Graph, from int, ep EdgeProcessor) {
	for to := 0; to < g.Capacity(); to++ {
		ep(g, from, to)
	}
}

func ForEdgesTo(g Graph, to int, ep EdgeProcessor) {
	for from := 0; from < g.Capacity(); from++ {
		ep(g, from, to)
	}
}

func ForNeighbours(g Graph, index int, vp VertexProcessor) {
	ForEdgesFrom(g, index, IfConnected(SelectTo(vp)))
}

func ForConnections(g Graph, index int, ep EdgeProcessor) {
	ForEdgesFrom(g, index, IfConnected(ep))
}

func Neighbours(g Graph, index int) []int {
	neighbours := make([]int, 0)
	ForNeighbours(g, index, CollectVerticesInto(&neighbours))
	return neighbours
}

func FirstPresent(g Graph) int {
	start := -1
	for i := 0; i < g.Capacity(); i++ {
		if g.IsPresent(i) {
			start = i
			break
		}
	}
	return start
}

func DepthFirst(g Graph, start int, down VertexProcessor, up VertexProcessor, ep EdgeProcessor) {
	if down == nil {
		down = EmptyVertexProcessor
	}
	if up == nil {
		up = EmptyVertexProcessor
	}
	if ep == nil {
		ep = EmptyEdgeProcessor
	}
	visited := make([]bool, g.Capacity())
	depthFirst(g, visited, start, down, up, ep)
}

func depthFirst(g Graph, visited []bool, index int, down VertexProcessor, up VertexProcessor, ep EdgeProcessor) {
	visited[index] = true
	down(g, index)
	ForNeighbours(g, index, func(g Graph, i int) {
		if !visited[i] {
			ep(g, index, i)
			depthFirst(g, visited, i, down, up, ep)
		}
	})
	up(g, index)
}

func BreadthFirst(g Graph, start int, vp VertexProcessor) {
	if vp == nil {
		vp = EmptyVertexProcessor
	}
	visited := make([]bool, g.Capacity())
	breadthFirst(g, visited, start, vp)
}

func breadthFirst(g Graph, visited []bool, start int, vp VertexProcessor) {
	nexts := list.New()
	nexts.PushBack(start)
	visited[start] = true
	for e := nexts.Front(); e != nil; e = nexts.Front() {
		index := e.Value.(int)
		nexts.Remove(e)
		vp(g, index)
		ForNeighbours(g, index, func(g Graph, i int) {
			if !visited[i] {
				visited[i] = true
				nexts.PushBack(i)
			}
		})
	}
}

func IsConnectedFrom(g Graph, start int) bool {
	count := 0
	counter := func(_ Graph, _ int) { count++ }
	DepthFirst(g, start, counter, nil, nil)
	return count == g.VertexCount()
}

func IsRoot(g Graph, index int) bool {
	if !g.IsPresent(index) {
		return false
	}
	for i := 0; i < g.Capacity(); i++ {
		if g.IsConnected(i, index) {
			return false
		}
	}
	return true
}

func IsLeaf(g Graph, index int) bool {
	if !g.IsPresent(index) {
		return false
	}
	for i := 0; i < g.Capacity(); i++ {
		if g.IsConnected(index, i) {
			return false
		}
	}
	return true
}
