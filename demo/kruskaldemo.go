package demo

import (
	"ag/graph"
	"ag/graph/algorithms"
)

func KruskalDemo() {
	cities := createCities()
	dc := CreateDistanceCalculator(cities.GetVertexDataLayer("coordinates"))
	algorithms.Kruskal(cities, dc, graph.PrintEdge)
}
