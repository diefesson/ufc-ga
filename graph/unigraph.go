package graph

type UniGraph struct {
	vertices []Vertex
	edges    [][]Edge
}

func makeUniEdges(size int) [][]Edge {
	edges := make([][]Edge, size)
	for i := 0; i < size; i++ {
		edges[i] = make([]Edge, size-i)
	}
	return edges
}

func NewUniGraph(size int) *UniGraph {
	return &UniGraph{vertices: make([]Vertex, size), edges: makeUniEdges(size)}
}

func (g *UniGraph) GetVertex(index int) *Vertex {
	return &g.vertices[index]
}

func (g *UniGraph) GetEdge(i, j int) *Edge {
	if i > j {
		i, j = j, i
	}
	j -= i
	return &g.edges[i][j]
}

func (g *UniGraph) Neighbours(index int) []int {
	neighbours := make([]int, 0)
	for i := 0; i < len(g.vertices); i++ {
		if GetConnection(g, index, i) {
			neighbours = append(neighbours, i)
		}
	}
	return neighbours
}

func (g *UniGraph) ForVertices(f VertexProcessor) {
	for i := 0; i < len(g.vertices); i++ {
		f(i, g.GetVertex(i))
	}
}

func (g *UniGraph) ForEdges(f EdgeProcessor) {
	for i := 0; i < len(g.vertices); i++ {
		for j := i; j < len(g.vertices); j++ {
			f(i, j, g.GetEdge(i, j))
		}
	}
}
