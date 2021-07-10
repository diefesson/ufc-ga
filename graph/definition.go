package graph

type Data interface{}

type Graph interface {
	Capacity() int
	VertexCount() int
	EdgeCount() int
	Add(index int)
	IsPresent(index int) bool
	Remove(index int)
	Connect(from, to int)
	IsConnected(from, to int) bool
	Disconnect(from, to int)
	ForNeighbours(index int, vp VertexProcessor)
	ForVertices(vp VertexProcessor)
	ForEdges(ep EdgeProcessor)
	CreateVertexDataLayer(key string) *UDDataLayer
	GetVertexDataLayer(key string) *UDDataLayer
	RemoveVertexDataLayer(key string)
	CreateEdgeDataLayer(key string) BDDataLayer
	GetEdgeDataLayer(key string) BDDataLayer
	RemoveEdgeDataLayer(key string)
	VerifyConnected() bool
	setVisited(index int, visited bool)
	isVisited(index int) bool
}
