package graph

import "fmt"

func solveUniIndex(from, to int) (int, int) {
	if from > to {
		from, to = to, from
	}
	return from, to - from
}

//TODO update to use data oriented information
func PrintVertex(index int) {
	fmt.Println(index)
}

//TODO update to use data oriented information
func PrintEdge(from, to int) {
	fmt.Println(from, to)
}

//TODO update to use data oriented information
func PrintGraph(g Graph) {
	g.ForVertices(PrintVertex)
	g.ForEdges(PrintEdge)
}
