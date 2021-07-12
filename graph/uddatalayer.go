package graph

type Data interface{}

type UDDataLayerProcessor func(index int, data Data)

type UDDataLayer struct {
	data []Data
}

func NewUnidimensionalDataLayer(size int) *UDDataLayer {
	return &UDDataLayer{
		data: make([]Data, size),
	}
}

func (dl *UDDataLayer) SetAll(d Data) {
	for i := 0; i < len(dl.data); i++ {
		dl.data[i] = d
	}
}

func (dl *UDDataLayer) Get(index int) Data {
	return dl.data[index]
}

func (dl *UDDataLayer) Set(index int, d Data) {
	dl.data[index] = d
}

func (dl *UDDataLayer) Size() int {
	return len(dl.data)
}
