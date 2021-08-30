package graph

type VertexFilter func(g Graph, index int) bool

type EdgeFilter func(g Graph, from, to int) bool

func PassVertexFilter(g Graph, index int) bool {
	return true
}

func PassEdgeFilter(g Graph, from, to int) bool {
	return true
}

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

func IfRoot(vp VertexProcessor) VertexProcessor {
	return func(g Graph, index int) {
		if IsRoot(g, index) {
			vp(g, index)
		}
	}
}

func IfLeaf(vp VertexProcessor) VertexProcessor {
	return func(g Graph, index int) {
		if IsLeaf(g, index) {
			vp(g, index)
		}
	}
}
