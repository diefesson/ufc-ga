package main

import (
	"ag/graph"
	"fmt"
)

type Inter interface {
	foo() string
}

type A struct {
	data int
}

type B struct {
	data float32
}

func (a *A) foo() string {
	return fmt.Sprint(a.data)
}

func (b *B) foo() string {
	return fmt.Sprint(b.data)
}

func main() {
	var i graph.Graph
	i = graph.NewDiGraph(10)
	i.GetEdge(0, 1).Connected = true
	i.GetEdge(1, 0).Connected = true
	i.GetEdge(1, 2).Connected = true
	fmt.Println(i.Neighbours(2))
}
