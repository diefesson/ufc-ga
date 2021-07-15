package demo

import (
	"github.com/diefesson/graph-algorithms/graph"
	"github.com/diefesson/graph-algorithms/graph/algorithms"
)

func KruskalDemo() {
	cities := createCities()
	dc := createDistanceCalculator(cities.GetVertexDataLayer("coordinates"))
	algorithms.Kruskal(cities, dc, graph.PrintEdge)
}
