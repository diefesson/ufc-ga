package demos

import (
	"fmt"

	"github.com/diefesson/ufc-ga/graph"
	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func FloydWarshallDemo() {
	cities := createDiCities()
	dc := createDistanceCalculator(createCoordinates())

	distances, sucessors, err := algorithms.FloydWarshall(cities, dc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Distances:")
	for _, r := range distances {
		fmt.Println(r)
	}
	fmt.Println("Sucessors:")
	for _, r := range sucessors {
		fmt.Println(r)
	}

	path := []int{}
	algorithms.FloydWarshallWalk(cities, sucessors, 5, 1, graph.CollectVerticesInto(&path))
	fmt.Println("Path:", path)
}
