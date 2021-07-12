package algorithms

type edge struct {
	from, to int
	distance float64
}

type edgeQueue []edge

func (eq edgeQueue) Len() int {
	return len(eq)
}

func (eq edgeQueue) Less(i, j int) bool {
	return eq[i].distance < eq[j].distance
}

func (eq edgeQueue) Swap(i, j int) {
	eq[i], eq[j] = eq[j], eq[i]
}

func (eq *edgeQueue) Push(e interface{}) {
	*eq = append(*eq, e.(edge))
}

func (eq *edgeQueue) Pop() interface{} {
	old := *eq
	n := len(old)
	e := old[n-1]
	*eq = old[0 : n-1]
	return e
}
