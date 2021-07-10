package graph

type VertexProcessor func(index int)

func emptyVertexProcessor(index int) {}

type EdgeProcessor func(from, to int)

func emptyEdgeProcessor(from, to int) {}
