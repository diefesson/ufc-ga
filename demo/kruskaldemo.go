package demo

import (
	"github.com/diefesson/ufc-ga/graph"
	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func KruskalDemo() {
	cities := createCities()
	dc := createDistanceCalculator(cities.GetVertexDataLayer("coordinates"))
	algorithms.Kruskal(cities, dc, graph.PrintEdge)
}
