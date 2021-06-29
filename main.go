package main

import (
	"ag/graph"
	"fmt"
)

func main() {
	var i graph.Graph
	i = graph.NewDiGraph(10)
	i.Connect(0, 1)
	fmt.Println(i.Size())
	fmt.Println(i.VerifyConnected())
}
