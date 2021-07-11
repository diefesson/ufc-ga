package graph

// Graphs

type VertexProcessor func(index int)

func emptyVertexProcessor(index int) {}

type EdgeProcessor func(from, to int)

func emptyEdgeProcessor(from, to int) {}

// Disjoint set

type NodeProcessor func(index int)

// Kruskal algorithm

type DistanceCalculator func(g *UniGraph, from, to int) float64
