package graph

func solveUniIndex(from, to int) (int, int) {
	if from > to {
		from, to = to, from
	}
	return from, to - from
}

func PrintGraph(g Graph) {
	g.ForVertices(PrintVertex)
	g.ForEdges(PrintEdge)
}
