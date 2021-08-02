package demos

import (
	"github.com/diefesson/ufc-ga/graph"
	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func KruskalDemo() {
	cities := createUniCities()
	dc := createDistanceCalculator(createCoordinates())
	algorithms.Kruskal(cities, dc, graph.PrintEdge)
}
