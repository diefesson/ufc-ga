package demos

import (
	"fmt"

	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func BellmanFordDemo() {
	cities := createDiCities()
	distance, path := algorithms.BellmanFord(cities, 5, 1, createDistanceCalculator(createCoordinates()))
	fmt.Println("Distance:", distance)
	fmt.Println("Path:", path)
}
