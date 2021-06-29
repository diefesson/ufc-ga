package graph

type data interface{}

type Vertex struct {
	Data data
}

type Edge struct {
	Connected bool
	Data      data
}

type VertexProcessor func(index int, vertex *Vertex)

type EdgeProcessor func(from, to int, edge *Edge)

type Graph interface {
	GetVertex(index int) *Vertex
	GetEdge(from, to int) *Edge
	Neighbours(index int) []int
	ForVertices(VertexProcessor)
	ForEdges(EdgeProcessor)
}

func GetConnection(g Graph, i, j int) bool {
	return g.GetEdge(i, j).Connected
}

func SetConnection(g Graph, i, j int, connected bool) {
	g.GetEdge(i, j).Connected = connected
}

func ClearVerticesData(g Graph) {
	g.ForVertices(func(index int, vertex *Vertex) { vertex.Data = nil })
}

func ClearEdgesData(g Graph) {
	g.ForEdges(func(from, to int, edge *Edge) { edge.Data = nil })
}
