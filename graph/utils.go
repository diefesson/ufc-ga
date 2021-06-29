package graph

import "fmt"

func PrintVertex(index int, vertex Vertex) {
	fmt.Println(index, ":", vertex)
}

func PrintEdge(from, to int, edge Edge) {
	fmt.Println(from, to, ":", edge)
}

func PrintGraph(g Graph) {
	g.ForVertices(PrintVertex)
	g.ForEdges(PrintEdge)
}
