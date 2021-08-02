package demos

import (
	"github.com/diefesson/ufc-ga/graph"
	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func TopologyDemo() {
	cities := createDiCities()
	algorithms.InTopologicalOrder(cities, graph.PrintVertex)
}
