package graph

type edgeImpl struct {
	dataHolderImpl
	connected bool
}

func (e *edgeImpl) IsConnected() bool {
	return e.connected
}

func (e *edgeImpl) setConnected(connected bool) {
	e.connected = connected
}
