package demos

import (
	"github.com/diefesson/ufc-ga/graph"
	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func PrimDemo() {
	cities := createUniCities()
	dc := createDistanceCalculator(createCoordinates())
	algorithms.Prim(cities, dc, graph.PrintEdge)
}
