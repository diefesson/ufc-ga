package demos

import (
	"math"

	"github.com/diefesson/ufc-ga/graph"
)

type city struct {
	x, y float64
}

func createCoordinates() *graph.UDDataLayer {
	coordinates := graph.NewUnidimensionalDataLayer(6)
	coordinates.Set(0, city{-2, -2})
	coordinates.Set(1, city{-1, -1})
	coordinates.Set(2, city{1, -1})
	coordinates.Set(3, city{1, 1})
	coordinates.Set(4, city{-1, 1})
	coordinates.Set(5, city{2, 2})
	return coordinates
}

func createUniCities() *graph.UniGraph {
	cities := graph.NewUniGraph(6)
	cities.ForEdges(graph.Connect)
	return cities
}

func createDiCities() *graph.DiGraph {
	cities := graph.NewDiGraph(6)
	cities.Connect(0, 1)
	cities.Connect(5, 3)
	cities.Connect(3, 2)
	cities.Connect(3, 4)
	cities.Connect(2, 1)
	cities.Connect(4, 1)
	return cities
}

func createDistanceCalculator(coordinates *graph.UDDataLayer) graph.DistanceCalculator {
	return func(_ graph.Graph, from, to int) float64 {
		fc := coordinates.Get(from).(city)
		tc := coordinates.Get(to).(city)
		dx := tc.x - fc.x
		dy := tc.y - fc.y
		return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	}
}
