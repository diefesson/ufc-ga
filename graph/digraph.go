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
	ForEdgesFrom(g, index, IfConnected(Disconnect))
	ForEdgesTo(g, index, IfConnected(Disconnect))
}

func (g *DiGraph) Connect(from, to int) {
	g.baseGraph.Add(from)
	g.baseGraph.Add(to)
	g.baseGraph.Connect(from, to)
}

func (g *DiGraph) ForVertices(vp VertexProcessor) {
	g.baseGraph.ForVertices(vpCompat(g, vp))
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

func (g *DiGraph) Clone() Graph {
	c := NewDiGraph(g.Capacity())
	transferBaseData(&c.baseGraph, &g.baseGraph)
	return c
}
