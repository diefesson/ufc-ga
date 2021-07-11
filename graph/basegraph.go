package graph

type baseGraph struct {
	vertexCount   int
	edgeCount     int
	vertexPresent []bool
	vertexVisited []bool
	edgeConnected [][]bool
	vertexLayers  map[string]*UDDataLayer
	edgeLayers    map[string]BDDataLayer
}

func makeBaseGraph(capacity int) baseGraph {
	return baseGraph{
		vertexPresent: make([]bool, capacity),
		vertexVisited: make([]bool, capacity),
		vertexLayers:  make(map[string]*UDDataLayer),
		edgeLayers:    make(map[string]BDDataLayer),
	}
}

func (g *baseGraph) Capacity() int {
	return len(g.vertexPresent)
}

func (g *baseGraph) VertexCount() int {
	return g.vertexCount
}

func (g *baseGraph) EdgeCount() int {
	return g.edgeCount
}

func (g *baseGraph) Add(index int) {
	if !g.vertexPresent[index] {
		g.vertexPresent[index] = true
		g.vertexCount++
	}
}

func (g *baseGraph) IsPresent(index int) bool {
	return g.vertexPresent[index]
}

func (g *baseGraph) Remove(index int) {
	if g.vertexPresent[index] {
		g.vertexPresent[index] = false
		g.vertexCount--
	}
}

func (g *baseGraph) Connect(from, to int) {
	if !g.edgeConnected[from][to] {
		g.edgeConnected[from][to] = true
		g.edgeCount++
	}
}

func (g *baseGraph) IsConnected(from, to int) bool {
	return g.edgeConnected[from][to]
}

func (g *baseGraph) Disconnect(from, to int) {
	if g.edgeConnected[from][to] {
		g.edgeConnected[from][to] = false
		g.edgeCount--
	}
}

func (g *baseGraph) ForVertices(f VertexProcessor) {
	for i := 0; i < g.Capacity(); i++ {
		f(i)
	}
}

func (g *baseGraph) CreateVertexDataLayer(key string) *UDDataLayer {
	dataLayer := NewUnidimensionalDataLayer(g.Capacity())
	g.vertexLayers[key] = dataLayer
	return dataLayer
}

func (g *baseGraph) GetVertexDataLayer(key string) *UDDataLayer {
	return g.vertexLayers[key]
}

func (g *baseGraph) RemoveVertexDataLayer(key string) {
	delete(g.vertexLayers, key)
}

func (g *baseGraph) GetEdgeDataLayer(key string) BDDataLayer {
	return g.edgeLayers[key]
}

func (g *baseGraph) RemoveEdgeDataLayer(key string) {
	delete(g.edgeLayers, key)
}

func (g *baseGraph) isVisited(index int) bool {
	return g.vertexVisited[index]
}

func (g *baseGraph) setVisited(index int, visited bool) {
	g.vertexVisited[index] = visited
}
