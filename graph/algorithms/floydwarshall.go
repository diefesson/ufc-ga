package algorithms

import (
	"errors"
	"math"

	"github.com/diefesson/ufc-ga/graph"
)

func FloydWarshall(g graph.Graph, dc graph.DistanceCalculator) ([][]float64, [][]int, error) {
	distances := make([][]float64, g.Capacity())
	sucessors := make([][]int, g.Capacity())
	for i := 0; i < g.Capacity(); i++ {
		distances[i] = make([]float64, g.Capacity())
		sucessors[i] = make([]int, g.Capacity())
	}

	for i := 0; i < g.Capacity(); i++ {
		for j := 0; j < g.Capacity(); j++ {
			if g.IsConnected(i, j) {
				distances[i][j] = dc(g, i, j)
				sucessors[i][j] = j
			} else {
				distances[i][j] = math.Inf(1)
				sucessors[i][j] = -1
			}
		}
	}

	for k := 0; k < g.Capacity(); k++ {
		for i := 0; i < g.Capacity(); i++ {
			for j := 0; j < g.Capacity(); j++ {
				if distances[i][k]+distances[k][j] < distances[i][j] {
					distances[i][j] = distances[i][k] + distances[k][j]
					sucessors[i][j] = k
				}
			}
		}
	}

	for i := 0; i < g.Capacity(); i++ {
		if distances[i][i] < 0 {
			return nil, nil, errors.New("negative cycle detected")
		}
	}

	return distances, sucessors, nil
}

func FloydWarshallWalk(g graph.Graph, sucessors [][]int, from, to int, vp graph.VertexProcessor) {
	for v := from; v != to; v = sucessors[v][to] {
		vp(g, v)
	}
	vp(g, to)
}
