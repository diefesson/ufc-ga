package graph

type BDDataLayerProcessor func(i, j int, data Data)

type BDDataLayer interface {
	Get(i, j int) Data
	Set(i, j int, data Data)
	Size() int
}

type UniEdgeDataLayer struct {
	data [][]Data
}

func NewUniEdgeDataLayer(size int) *UniEdgeDataLayer {
	data := make([][]Data, size)
	for i := 0; i < size; i++ {
		data[i] = make([]Data, size-i)
	}
	return &UniEdgeDataLayer{data}
}

func (dl *UniEdgeDataLayer) Get(i, j int) Data {
	if i > j {
		i, j = j, i
	}
	return dl.data[i][j]
}

func (dl *UniEdgeDataLayer) Set(i, j int, data Data) {
	if i > j {
		i, j = j, i
	}
	dl.data[i][j] = data
}

func (dl *UniEdgeDataLayer) Size() int {
	return len(dl.data)
}

type DiEdgeDataLayer struct {
	data [][]Data
}

func NewDiEdgeDataLayer(size int) *DiEdgeDataLayer {
	data := make([][]Data, size)
	for i := 0; i < size; i++ {
		data[i] = make([]Data, size)
	}
	return &DiEdgeDataLayer{data}
}

func (dl *DiEdgeDataLayer) Get(from, to int) Data {
	return dl.data[from][to]
}

func (dl *DiEdgeDataLayer) Set(from, to int, data Data) {
	dl.data[from][to] = data
}

func (dl *DiEdgeDataLayer) Size() int {
	return len(dl.data)
}
