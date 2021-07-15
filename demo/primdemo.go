package demo

import (
	"ag/graph"
	"ag/graph/algorithms"
)

func PrimDemo() {
	cities := createCities()
	dc := createDistanceCalculator(cities.GetVertexDataLayer("coordinates"))
	algorithms.Prim(cities, dc, graph.PrintEdge)
}
