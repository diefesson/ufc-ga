package demos

import (
	"fmt"

	"github.com/diefesson/ufc-ga/graph"
	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func BellmanFordDemo() {
	cities := createDiCities()
	path := make([]int, 0)
	distance, err := algorithms.BellmanFord(cities, 5, 1, createDistanceCalculator(createCoordinates()), graph.CollectVerticesInto(&path))
	if err != nil {
		panic(err)
	}
	fmt.Println("Distance:", distance)
	fmt.Println("Path:", path)
}
