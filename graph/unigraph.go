package graph

type UniGraph struct {
	baseGraph
}

func makeUniEdges(capacity int) [][]bool {
	edges := make([][]bool, capacity)
	for i := 0; i < capacity; i++ {
		edges[i] = make([]bool, capacity-i)
	}
	return edges
}

func NewUniGraph(capacity int) *UniGraph {
	base := makeBaseGraph(capacity)
	base.edgeConnected = makeUniEdges(capacity)
	return &UniGraph{base}
}

func (g *UniGraph) Remove(index int) {
	g.baseGraph.Remove(index)
	g.ForNeighbours(index, func(i int) { g.Disconnect(index, i) })
}

func (g *UniGraph) Connect(from, to int) {
	g.baseGraph.Add(from)
	g.baseGraph.Add(to)
	from, to = solveUniIndex(from, to)
	g.baseGraph.Connect(from, to)
}

func (g *UniGraph) IsConnected(from, to int) bool {
	from, to = solveUniIndex(from, to)
	return g.baseGraph.IsConnected(from, to)
}

func (g *UniGraph) Disconnect(from, to int) {
	from, to = solveUniIndex(from, to)
	g.baseGraph.Disconnect(from, to)
}

func (g *UniGraph) ForNeighbours(index int, vp VertexProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		if g.IsConnected(index, i) {
			vp(i)
		}
	}
}

func (g *UniGraph) ForEdges(f EdgeProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		for j := i; j < g.Capacity(); j++ {
			f(i, j)
		}
	}
}

func (g *UniGraph) findStart() int {
	start := -1
	for i := 0; i < g.Capacity(); i++ {
		if g.IsPresent(i) {
			start = i
			break
		}
	}
	return start
}

func (g *UniGraph) CreateEdgeDataLayer(key string) BDDataLayer {
	dataLayer := NewUniEdgeDataLayer(g.Capacity())
	g.edgeLayers[key] = dataLayer
	return dataLayer
}

func (g *UniGraph) VerifyConnected() bool {
	if g.EdgeCount() == 0 {
		return true
	}
	start := g.findStart()
	return IsConnectedFrom(g, start)
}
