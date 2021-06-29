package graph

type vertexImpl struct {
	dataHolderImpl
	present bool
	visited bool
}

func (v *vertexImpl) IsPresent() bool {
	return v.present
}

func (v *vertexImpl) setPresent(present bool) {
	v.present = present
}

func (v *vertexImpl) isVisited() bool {
	return v.visited
}

func (v *vertexImpl) setVisited(visited bool) {
	v.visited = visited
}
