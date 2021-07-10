package graph

type UDDataLayerProcessor func(index int, data Data)

type UDDataLayer struct {
	data []Data
}

func NewUnidimensionalDataLayer(size int) *UDDataLayer {
	return &UDDataLayer{
		data: make([]Data, size),
	}
}

func (dl *UDDataLayer) Get(index int) Data {
	return dl.data[index]
}

func (dl *UDDataLayer) Set(index int, data Data) {
	dl.data[index] = data
}

func (dl *UDDataLayer) Size() int {
	return len(dl.data)
}
