package graph

type DiGraph struct {
	baseGraphImpl
	edges [][]edgeImpl
}

func makeDiEdges(vertexCount int) [][]edgeImpl {
	edges := make([][]edgeImpl, vertexCount)
	for i := 0; i < vertexCount; i++ {
		edges[i] = make([]edgeImpl, vertexCount)
	}
	return edges
}

func NewDiGraph(capacity int) *DiGraph {
	return &DiGraph{
		makeBaseGraphImpl(capacity),
		makeDiEdges(capacity),
	}
}

func (g *DiGraph) Remove(index int) {
	g.baseGraphImpl.Remove(index)
	g.ForFrom(index, func(i int, vertex Vertex) { g.GetEdge(index, i).setConnected(false) })
	g.ForTo(index, func(i int, vertex Vertex) { g.GetEdge(i, index).setConnected(false) })
}

func (g *DiGraph) Connect(from, to int) {
	g.baseGraphImpl.Add(from)
	g.baseGraphImpl.Add(to)
	g.GetEdge(from, to).setConnected(true)
}

func (g *DiGraph) Disconnect(from, to int) {
	g.GetEdge(from, to).setConnected(false)
}

func (g *DiGraph) GetEdge(from, to int) Edge {
	return &g.edges[from][to]
}

func (g *DiGraph) ForFrom(index int, vp VertexProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		if g.GetEdge(index, i).IsConnected() {
			vp(i, g.GetVertex(i))
		}
	}
}

func (g *DiGraph) ForNeighbours(index int, vp VertexProcessor) {
	g.ForFrom(index, vp)
}

func (g *DiGraph) ForTo(index int, vp VertexProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		if g.GetEdge(i, index).IsConnected() {
			vp(i, g.GetVertex(i))
		}
	}
}

func (g *DiGraph) ForEdges(f EdgeProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		for j := 0; j < g.Capacity(); j++ {
			f(i, j, g.GetEdge(i, j))
		}
	}
}

func (g *DiGraph) VerifyConnected() bool {
	for i := 0; i < g.Capacity(); i++ {
		if g.GetVertex(i).IsPresent() && !IsConnectedFrom(g, i) {
			return false
		}
	}
	return true
}
