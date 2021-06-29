package graph

type baseGraphImpl struct {
	size     int
	vertices []vertexImpl
}

func makeBaseGraphImpl(capacity int) baseGraphImpl {
	return baseGraphImpl{vertices: make([]vertexImpl, capacity)}
}

func (g *baseGraphImpl) Capacity() int {
	return len(g.vertices)
}

func (g *baseGraphImpl) Size() int {
	return g.size
}

func (g *baseGraphImpl) GetVertex(index int) Vertex {
	return &g.vertices[index]
}

func (g *baseGraphImpl) Add(index int) {
	vertex := g.GetVertex(index)
	if !vertex.IsPresent() {
		vertex.setPresent(true)
		g.size++
	}
}

func (g *baseGraphImpl) Remove(index int) {
	vertex := g.GetVertex(index)
	if vertex.IsPresent() {
		vertex.setPresent(false)
		g.size--
	}
}

func (g *baseGraphImpl) ForVertices(f VertexProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		f(i, g.GetVertex(i))
	}
}
