package graph

type Data interface{}

type DataHolder interface {
	Data() Data
	SetData(data Data)
}

type Edge interface {
	DataHolder
	IsConnected() bool
	setConnected(connected bool)
}

type Vertex interface {
	DataHolder
	IsPresent() bool
	setPresent(present bool)
	isVisited() bool
	setVisited(visited bool)
}

type VertexProcessor func(index int, vertex Vertex)

func emptyVertexProcessor(index int, vertex Vertex) {}

type EdgeProcessor func(from, to int, edge Edge)

func emptyEdgeProcessor(from, to int, edge Edge) {}

type Graph interface {
	Capacity() int
	Size() int
	GetVertex(index int) Vertex
	Add(index int)
	Remove(index int)
	GetEdge(from, to int) Edge
	Connect(from, to int)
	Disconnect(from, to int)
	ForNeighbours(index int, vp VertexProcessor)
	ForVertices(VertexProcessor)
	ForEdges(EdgeProcessor)
	VerifyConnected() bool
}
