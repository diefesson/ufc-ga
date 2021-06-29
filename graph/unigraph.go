package graph

type UniGraph struct {
	baseGraphImpl
	edges [][]edgeImpl
}

func makeUniEdges(capacity int) [][]edgeImpl {
	edges := make([][]edgeImpl, capacity)
	for i := 0; i < capacity; i++ {
		edges[i] = make([]edgeImpl, capacity-i)
	}
	return edges
}

func NewUniGraph(capacity int) *UniGraph {
	return &UniGraph{
		makeBaseGraphImpl(capacity),
		makeUniEdges(capacity),
	}
}

func (g *UniGraph) Remove(index int) {
	g.baseGraphImpl.Remove(index)
	g.ForNeighbours(index, func(i int, _ Vertex) { g.GetEdge(index, i).setConnected(false) })
}

func (g *UniGraph) Connect(from, to int) {
	g.baseGraphImpl.Add(from)
	g.baseGraphImpl.Add(to)
	g.GetEdge(from, to).setConnected(true)
}

func (g *UniGraph) Disconnect(from, to int) {
	g.GetEdge(from, to).setConnected(false)
}

func (g *UniGraph) GetEdge(i, j int) Edge {
	if i > j {
		i, j = j, i
	}
	j -= i
	return &g.edges[i][j]
}

func (g *UniGraph) ForNeighbours(index int, vp VertexProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		if g.GetEdge(index, i).IsConnected() {
			vp(i, g.GetVertex(i))
		}
	}
}

func (g *UniGraph) ForEdges(f EdgeProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		for j := i; j < g.Capacity(); j++ {
			f(i, j, g.GetEdge(i, j))
		}
	}
}

func (g *UniGraph) findStart() int {
	start := -1
	for i := 0; i < g.Capacity(); i++ {
		if g.GetVertex(i).IsPresent() {
			start = i
			break
		}
	}
	return start
}

func (g *UniGraph) VerifyConnected() bool {
	if g.Size() == 0 {
		return true
	}
	start := g.findStart()
	return IsConnectedFrom(g, start)
}
