package demos

import (
	"fmt"

	"github.com/diefesson/ufc-ga/graph"
	"github.com/diefesson/ufc-ga/graph/algorithms"
)

func EdmondsKarpDemo() {
	cities := createDiCities()
	bandwidths := createBandwidths(cities)
	total, flows := algorithms.EdmondsKarp(cities, bandwidths, 3, 1)
	cities.ForEdges(func(g graph.Graph, f, t int) {
		fmt.Println(
			"edge:", f, t,
			"capacity:", bandwidths[f][t],
			"flow:", flows[f][t], "/", flows[t][f],
		)
	})
	fmt.Println("total flow:", total)
}
