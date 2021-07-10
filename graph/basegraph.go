package graph

type baseGraphImpl struct {
	vertexCount   int
	edgeCount     int
	vertexPresent []bool
	vertexVisited []bool
	edgeConnected [][]bool
	vertexLayers  map[string]*UDDataLayer
	edgeLayers    map[string]BDDataLayer
}

func makeBaseGraphImpl(capacity int) baseGraphImpl {
	return baseGraphImpl{
		vertexPresent: make([]bool, capacity),
		vertexVisited: make([]bool, capacity),
	}
}

func (g *baseGraphImpl) Capacity() int {
	return len(g.vertexPresent)
}

func (g *baseGraphImpl) VertexCount() int {
	return g.vertexCount
}

func (g *baseGraphImpl) EdgeCount() int {
	return g.edgeCount
}

func (g *baseGraphImpl) Add(index int) {
	if !g.vertexPresent[index] {
		g.vertexPresent[index] = true
		g.vertexCount++
	}
}

func (g *baseGraphImpl) IsPresent(index int) bool {
	return g.vertexPresent[index]
}

func (g *baseGraphImpl) Remove(index int) {
	if g.vertexPresent[index] {
		g.vertexPresent[index] = false
		g.vertexCount--
	}
}

func (g *baseGraphImpl) Connect(from, to int) {
	if !g.edgeConnected[from][to] {
		g.edgeConnected[from][to] = true
		g.edgeCount++
	}
}

func (g *baseGraphImpl) IsConnected(from, to int) bool {
	return g.edgeConnected[from][to]
}

func (g *baseGraphImpl) Disconnect(from, to int) {
	if g.edgeConnected[from][to] {
		g.edgeConnected[from][to] = false
		g.edgeCount--
	}
}

func (g *baseGraphImpl) ForVertices(f VertexProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		f(i)
	}
}

func (g *baseGraphImpl) CreateVertexDataLayer(key string) *UDDataLayer {
	dataLayer := NewUnidimensionalDataLayer(g.Capacity())
	g.vertexLayers[key] = dataLayer
	return dataLayer
}

func (g *baseGraphImpl) GetVertexDataLayer(key string) *UDDataLayer {
	return g.vertexLayers[key]
}

func (g *baseGraphImpl) RemoveVertexDataLayer(key string) {
	delete(g.vertexLayers, key)
}

func (g *baseGraphImpl) GetEdgeDataLayer(key string) BDDataLayer {
	return g.edgeLayers[key]
}

func (g *baseGraphImpl) RemoveEdgeDataLayer(key string) {
	delete(g.edgeLayers, key)
}

func (g *baseGraphImpl) isVisited(index int) bool {
	return g.vertexVisited[index]
}

func (g *baseGraphImpl) setVisited(index int, visited bool) {
	g.vertexVisited[index] = visited
}
