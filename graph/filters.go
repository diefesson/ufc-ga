package graph

type VertexFilter func(g Graph, index int) bool

type EdgeFilter func(g Graph, from, to int) bool

func IfPresent(vp VertexProcessor) VertexProcessor {
	return func(g Graph, index int) {
		if g.IsPresent(index) {
			vp(g, index)
		}
	}
}

func IfConnected(ep EdgeProcessor) EdgeProcessor {
	return func(g Graph, from, to int) {
		if g.IsConnected(from, to) {
			ep(g, from, to)
		}
	}
}

func IfNotLoopback(ep EdgeProcessor) EdgeProcessor {
	return func(g Graph, from, to int) {
		if from != to {
			ep(g, from, to)
		}
	}
}
