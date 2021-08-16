package demos

import (
	"fmt"

	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func DAGSPDemo() {
	graph := createDiCities()
	dc := createDistanceCalculator(createCoordinates())
	distance, path := algorithms.DAGSP(graph, 5, 1, dc)
	fmt.Println("Distance:", distance)
	fmt.Println("Path:", path)
}
