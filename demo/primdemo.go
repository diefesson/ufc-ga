package demo

import (
	"github.com/diefesson/ufc-ga/graph"
	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func PrimDemo() {
	cities := createCities()
	dc := createDistanceCalculator(cities.GetVertexDataLayer("coordinates"))
	algorithms.Prim(cities, dc, graph.PrintEdge)
}
