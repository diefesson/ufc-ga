package graph

import "fmt"

type VertexProcessor func(g Graph, index int)

func EmptyVertexProcessor(g Graph, index int) {}

type EdgeProcessor func(g Graph, from, to int)

func EmptyEdgeProcessor(g Graph, from, to int) {}

type NodeProcessor func(index int)

type DistanceCalculator func(g Graph, from, to int) float64

func vpCompat(g Graph, vp VertexProcessor) func(*baseGraph, int) {
	return func(_ *baseGraph, i int) { vp(g, i) }
}

func Add(g Graph, index int) {
	g.Add(index)
}

func Remove(g Graph, index int) {
	g.Remove(index)
}

func Connect(g Graph, from, to int) {
	g.Connect(from, to)
}

func Disconnect(g Graph, from, to int) {
	g.Disconnect(from, to)
}

func SelectFrom(vp VertexProcessor) EdgeProcessor {
	return func(g Graph, from, to int) {
		vp(g, from)
	}
}

func SelectTo(vp VertexProcessor) EdgeProcessor {
	return func(g Graph, from, to int) {
		vp(g, to)
	}
}

func PrintVertex(_ Graph, index int) {
	fmt.Println(index)
}

func PrintEdge(_ Graph, from, to int) {
	fmt.Println(from, to)
}
