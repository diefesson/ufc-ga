package graph

type DiGraph struct {
	baseGraph
}

func makeDiEdges(vertexCount int) [][]bool {
	edges := make([][]bool, vertexCount)
	for i := 0; i < vertexCount; i++ {
		edges[i] = make([]bool, vertexCount)
	}
	return edges
}

func NewDiGraph(capacity int) *DiGraph {
	base := makeBaseGraph(capacity)
	base.edgeConnected = makeDiEdges(capacity)
	return &DiGraph{base}
}

func (g *DiGraph) Remove(index int) {
	g.baseGraph.Remove(index)
	g.ForFrom(index, DisconnectFrom(index))
	g.ForTo(index, DisconnectTo(index))
}

func (g *DiGraph) Connect(from, to int) {
	g.baseGraph.Add(from)
	g.baseGraph.Add(to)
	g.baseGraph.Connect(from, to)
}

func (g *DiGraph) ForFrom(index int, vp VertexProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		if g.IsConnected(index, i) {
			vp(g, i)
		}
	}
}

func (g *DiGraph) ForNeighbours(index int, vp VertexProcessor) {
	g.ForFrom(index, vp)
}

func (g *DiGraph) ForTo(index int, vp VertexProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		if g.IsConnected(i, index) {
			vp(g, i)
		}
	}
}

func (g *DiGraph) ForVertices(f VertexProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		f(g, i)
	}
}

func (g *DiGraph) ForEdges(ep EdgeProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		for j := 0; j < g.Capacity(); j++ {
			ep(g, i, j)
		}
	}
}

func (g *DiGraph) CreateEdgeDataLayer(key string) BDDataLayer {
	dataLayer := NewDiEdgeDataLayer(g.Capacity())
	g.edgeLayers[key] = dataLayer
	return dataLayer
}

func (g *DiGraph) VerifyConnected() bool {
	for i := 0; i < g.Capacity(); i++ {
		if g.IsPresent(i) && !IsConnectedFrom(g, i) {
			return false
		}
	}
	return true
}
