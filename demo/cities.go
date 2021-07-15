package demo

import (
	"ag/graph"
	"math"
)

type city struct {
	x, y float64
}

func createCities() *graph.UniGraph {
	cities := graph.NewUniGraph(5)
	cities.ForEdges(graph.Connect)
	coordinates := cities.CreateVertexDataLayer("coordinates")
	coordinates.Set(0, city{0, 0})
	coordinates.Set(1, city{1, 0})
	coordinates.Set(2, city{0, 1})
	coordinates.Set(3, city{1, 1})
	coordinates.Set(4, city{2, 2})
	return cities
}

func createDistanceCalculator(coordinates *graph.UDDataLayer) graph.DistanceCalculator {
	return func(_ *graph.UniGraph, from, to int) float64 {
		fc := coordinates.Get(from).(city)
		tc := coordinates.Get(to).(city)
		dx := tc.x - fc.x
		dy := tc.y - fc.y
		return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	}
}
