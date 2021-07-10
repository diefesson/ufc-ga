package graph

type DisjointSet struct {
	parent     []int
	visited    []bool
	dataLayers map[string]*UDDataLayer
}

func NewDisjointSet(size int) *DisjointSet {
	ds := &DisjointSet{
		parent:  make([]int, size),
		visited: make([]bool, size),
	}
	for i := 0; i < size; i++ {
		ds.parent[i] = i
	}
	return ds
}

func (ds *DisjointSet) CreateDataLayer(key string) *UDDataLayer {
	dataLayer := NewUnidimensionalDataLayer(ds.Size())
	ds.dataLayers[key] = dataLayer
	return dataLayer
}

func (ds *DisjointSet) GetDataLayer(key string) *UDDataLayer {
	return ds.dataLayers[key]
}

func (ds *DisjointSet) RemoveDataLayer(key string) {
	delete(ds.dataLayers, key)
}

func (ds *DisjointSet) Size() int {
	return len(ds.parent)
}

func (ds *DisjointSet) ParentOf(index int) int {
	return ds.parent[index]
}

func (ds *DisjointSet) RepresentantOf(index int) int {
	for index != ds.parent[index] {
		index = ds.parent[index]
	}
	return index
}

func (ds *DisjointSet) representantIfUnvisited(index int) int {
	if ds.visited[index] {
		return -1
	}
	if ds.parent[index] == index {
		return index
	}
	return ds.representantIfUnvisited(ds.parent[index])
}

func (ds *DisjointSet) Join(i, j int) {
	ri := ds.RepresentantOf(i)
	rj := ds.RepresentantOf(j)
	if ri != rj {
		ds.parent[ri] = rj
	}
}

func (ds *DisjointSet) Separate(index int) {
	ds.parent[index] = index
}

type NodeProcessor func(index int)

func (ds *DisjointSet) ForEachNode(np NodeProcessor) {
	for i := 0; i < len(ds.parent); i++ {
		np(i)
	}
}

func (ds *DisjointSet) clearVisited() {
	for i := 0; i < ds.Size(); i++ {
		ds.visited[i] = false
	}
}

func (ds *DisjointSet) ForEachRepresentant(np NodeProcessor) {
	ds.clearVisited()
	for i := 0; i < ds.Size(); i++ {
		r := ds.representantIfUnvisited(i)
		if r != -1 {
			np(r)
		}
	}
}
