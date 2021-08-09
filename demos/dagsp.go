package demos

import (
	"fmt"

	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func DAGSPDemo() {
	graph := createDiCities()
	dc := createDistanceCalculator(createCoordinates())
	distance, path := algorithms.DAGSP(graph, 0, 3, dc)
	fmt.Println("distance: ", distance)
	for i, v := range path {
		fmt.Println(i, ":", v)
	}
}
