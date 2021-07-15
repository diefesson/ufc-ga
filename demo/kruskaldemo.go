package demo

import (
	"ag/graph"
	"ag/graph/algorithms"
)

func KruskalDemo() {
	cities := createCities()
	dc := createDistanceCalculator(cities.GetVertexDataLayer("coordinates"))
	algorithms.Kruskal(cities, dc, graph.PrintEdge)
}
