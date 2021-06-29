package graph

type dataHolderImpl struct {
	data Data
}

func (dh *dataHolderImpl) Data() Data {
	return dh.data
}

func (dh *dataHolderImpl) SetData(data Data) {
	dh.data = data
}
