package graph

type VertexFilter func(g Graph, index int) bool

type EdgeFilter func(g Graph, from, to int) bool

func FilterVertices(vf VertexFilter, vp VertexProcessor) VertexProcessor {
	return func(g Graph, index int) {
		if vf(g, index) {
			vp(g, index)
		}
	}
}

func FilterEdges(ef EdgeFilter, ep EdgeProcessor) EdgeProcessor {
	return func(g Graph, from, to int) {
		if ef(g, from, to) {
			ep(g, from, to)
		}
	}
}

func PresentVertexFilter(g Graph, index int) bool {
	return g.IsPresent(index)
}

func ConnectedEdgeFilter(g Graph, from, to int) bool {
	return g.IsConnected(from, to)
}
