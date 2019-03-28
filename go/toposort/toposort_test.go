package toposort

import "testing"

func TestTopoSortByKahn(t *testing.T) {
	aGraph := NewGraph(6)
	aGraph.AddEdge(0, 2)
	aGraph.AddEdge(1, 2)
	aGraph.AddEdge(2, 4)
	aGraph.AddEdge(3, 4)
	aGraph.AddEdge(4, 5)
	t.Log(aGraph.TopoSortByKahn())
}

func TestTopoSortByDFS(t *testing.T) {
	aGraph := NewGraph(6)
	aGraph.AddEdge(0, 2)
	aGraph.AddEdge(1, 2)
	aGraph.AddEdge(2, 4)
	aGraph.AddEdge(3, 4)
	aGraph.AddEdge(4, 5)
	t.Log(aGraph.TopoSortByDFS())
}
