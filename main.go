package main

import (
	"ag/graph"
	"fmt"
)

func main() {
	g := graph.NewUniGraph(3)
	g.ForEdges(func(from, to int) { g.Connect(from, to) })
	fmt.Println(g.VerifyConnected())
}
