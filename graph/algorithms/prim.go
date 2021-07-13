package algorithms

import (
	"ag/graph"
	"container/heap"
)

type prim struct {
	graph             *graph.UniGraph
	markedVertices    *graph.UDDataLayer
	markedEdges       *graph.UniEdgeDataLayer
	nextEdges         edgeQueue
	calculateDistance graph.DistanceCalculator
	onFound           graph.EdgeProcessor
}

func makePrim(g *graph.UniGraph, cd graph.DistanceCalculator, of graph.EdgeProcessor) prim {
	addedVertices := graph.NewUnidimensionalDataLayer(g.Capacity())
	edgeLayer := graph.NewUniEdgeDataLayer(g.Capacity())
	addedVertices.SetAll(UNMARKED)
	edgeLayer.SetAll(UNMARKED)
	return prim{
		graph:             g,
		markedVertices:    addedVertices,
		markedEdges:       edgeLayer,
		nextEdges:         make(edgeQueue, 0),
		calculateDistance: cd,
		onFound:           of,
	}
}

func enqueueEdges(p *prim, from int) graph.VertexProcessor {
	return func(_ graph.Graph, to int) {
		if p.markedEdges.Get(from, to) == UNMARKED {
			heap.Push(&p.nextEdges, edge{from, to, p.calculateDistance(p.graph, from, to)})
		}
	}
}

func (p *prim) nextEdge() edge {
	e := heap.Pop(&p.nextEdges).(edge)
	for p.markedVertices.Get(e.to) == MARKED {
		e = heap.Pop(&p.nextEdges).(edge)
	}
	return e
}

func (p *prim) run() {
	if p.graph.VertexCount() == 0 {
		return
	}
	start := graph.FirstPresent(p.graph)
	p.markedVertices.Set(start, MARKED)
	graph.ForNeighbours(p.graph, start, enqueueEdges(p, start))
	for i := 0; i < p.graph.VertexCount()-1; i++ {
		e := p.nextEdge()
		p.markedVertices.Set(e.to, MARKED)
		p.markedEdges.Set(e.from, e.to, MARKED)
		p.onFound(p.graph, e.from, e.to)
		graph.ForNeighbours(p.graph, e.to, enqueueEdges(p, e.to))
	}
}

func Prim(g *graph.UniGraph, dc graph.DistanceCalculator, of graph.EdgeProcessor) *graph.UniEdgeDataLayer {
	p := makePrim(g, dc, of)
	p.run()
	return p.markedEdges
}
