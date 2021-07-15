package demo

import (
	"github.com/diefesson/graph-algorithms/graph"
	"github.com/diefesson/graph-algorithms/graph/algorithms"
)

func PrimDemo() {
	cities := createCities()
	dc := createDistanceCalculator(cities.GetVertexDataLayer("coordinates"))
	algorithms.Prim(cities, dc, graph.PrintEdge)
}
