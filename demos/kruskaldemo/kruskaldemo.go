package kruskaldemo

import (
	"ag/graph"
	"math"
)

type city struct {
	x, y float64
}

func createGraph() *graph.UniGraph {
	g := graph.NewUniGraph(5)
	cities := g.CreateVertexDataLayer("cities")
	cities.Set(0, city{0, 0})
	cities.Set(1, city{1, 0})
	cities.Set(2, city{0, 1})
	cities.Set(3, city{1, 1})
	cities.Set(4, city{2, 2})
	return g
}

func CreateDistanceCalculator(g *graph.UniGraph) graph.DistanceCalculator {
	cities := g.GetVertexDataLayer("cities")
	return func(_ *graph.UniGraph, from, to int) float64 {
		fc := cities.Get(from).(city)
		tc := cities.Get(to).(city)
		dx := tc.x - fc.x
		dy := tc.y - fc.y
		return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	}
}

func RunDemo() {
	citiesGraph := createGraph()
	dc := CreateDistanceCalculator(citiesGraph)
	graph.KruskalAlgorithm(citiesGraph, dc, graph.PrintEdge)
}
