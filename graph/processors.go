package graph

import "fmt"

type VertexProcessor func(g Graph, index int)

func emptyVertexProcessor(g Graph, index int) {}

type EdgeProcessor func(g Graph, from, to int)

func emptyEdgeProcessor(g Graph, from, to int) {}

type NodeProcessor func(index int)

type DistanceCalculator func(g *UniGraph, from, to int) float64

func DisconnectFrom(from int) VertexProcessor {
	return func(g Graph, index int) {
		g.Disconnect(from, index)
	}
}

func DisconnectTo(to int) VertexProcessor {
	return func(g Graph, index int) {
		g.Disconnect(index, to)
	}
}

func ConnectProcessor(g Graph, from, to int) {
	g.Connect(from, to)
}

func PrintVertex(_ Graph, index int) {
	fmt.Println(index)
}

func PrintEdge(_ Graph, from, to int) {
	fmt.Println(from, to)
}
