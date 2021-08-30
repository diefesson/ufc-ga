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

func ForPathEdges(path []int, f func(f int, t int)) {
	for i := 0; i < len(path)-1; i++ {
		f(path[i], path[i+1])
	}
}
