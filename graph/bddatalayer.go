package graph

type BDDataLayerProcessor func(i, j int, data Data)

type BDDataLayer interface {
	Get(i, j int) Data
	Set(i, j int, data Data)
	SetAll(d Data)
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

func (dl *UniEdgeDataLayer) Get(from, to int) Data {
	from, to = solveUniIndex(from, to)
	return dl.data[from][to]
}

func (dl *UniEdgeDataLayer) Set(from, to int, data Data) {
	from, to = solveUniIndex(from, to)
	dl.data[from][to] = data
}

func (dl *UniEdgeDataLayer) SetAll(d Data) {
	for i := 0; i < len(dl.data); i++ {
		for j := 0; j < len(dl.data[i]); j++ {
			dl.data[i][j] = d
		}
	}
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

func (dl *DiEdgeDataLayer) SetAll(d Data) {
	for i := 0; i < dl.Size(); i++ {
		for j := 0; j < dl.Size(); j++ {
			dl.data[i][j] = d
		}
	}
}

func (dl *DiEdgeDataLayer) Size() int {
	return len(dl.data)
}
