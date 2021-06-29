package graph

type DiGraph struct {
	vertices []Vertex
	edges    [][]Edge
}

func makeDiEdges(vertexCount int) [][]Edge {
	edges := make([][]Edge, vertexCount)
	for i := 0; i < vertexCount; i++ {
		edges[i] = make([]Edge, vertexCount)
	}
	return edges
}

func NewDiGraph(size int) *DiGraph {
	return &DiGraph{vertices: make([]Vertex, size), edges: makeDiEdges(size)}
}

func (g *DiGraph) GetVertex(index int) *Vertex {
	return &g.vertices[index]
}

func (g *DiGraph) GetEdge(from, to int) *Edge {
	return &g.edges[from][to]
}

func (g *DiGraph) From(index int) []int {
	from := make([]int, 0)
	for i := 0; i < len(g.edges); i++ {
		if GetConnection(g, index, i) {
			from = append(from, i)
		}
	}
	return from
}

func (g *DiGraph) To(index int) []int {
	to := make([]int, 0)
	for i := 0; i < len(g.edges); i++ {
		if GetConnection(g, i, index) {
			to = append(to, i)
		}
	}
	return to
}

func (g *DiGraph) Neighbours(index int) []int {
	return g.From(index)
}

func (g *DiGraph) ForVertices(f VertexProcessor) {
	for i := 0; i < len(g.vertices); i++ {
		f(i, g.GetVertex(i))
	}
}

func (g *DiGraph) ForEdges(f EdgeProcessor) {
	for i := 0; i < len(g.vertices); i++ {
		for j := 0; j < len(g.vertices); j++ {
			f(i, j, g.GetEdge(i, j))
		}
	}
}
