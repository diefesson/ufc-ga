package demos

import (
	"fmt"

	"github.com/diefesson/ufc-ga/graph"
	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func DijkstraDemo() {
	cities := createDiCities()
	dc := createDistanceCalculator(createCoordinates())
	path := []int{}
	distance, err := algorithms.Dijkstra(cities, 5, 1, dc, graph.CollectVerticesInto(&path))
	if err != nil {
		panic(err)
	}
	fmt.Println("Distance:", distance)
	fmt.Println("Path:", path)
}
