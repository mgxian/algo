package graph

import "testing"

func TestUndirectedGraphAdjacencyMatrix(t *testing.T) {
	graph := newUndirectedGraph(6)
	graph.addEdge(0, 1)
	graph.addEdge(0, 4)
	graph.addEdge(1, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 3)
	graph.addEdge(2, 5)
	graph.addEdge(3, 5)
	graph.addEdge(4, 5)
	got := graph.String()
	want := "0--->1 0--->4 1--->2 1--->3 2--->3 2--->5 3--->5 4--->5"
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestUndirectedGraphAdjacencyList(t *testing.T) {
	graph := newUndirectedGraphWithAdjacencyList(6)
	graph.addEdge(0, 1)
	graph.addEdge(0, 4)
	graph.addEdge(1, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 3)
	graph.addEdge(2, 5)
	graph.addEdge(3, 5)
	graph.addEdge(4, 5)
	got := graph.String()
	want := "0--->1 0--->4 1--->2 1--->3 2--->3 2--->5 3--->5 4--->5"
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestDirectedWeightedGraph(t *testing.T) {
	graph := newDirectedWeightedGraph(6)
	graph.addEdge(0, 1, 10)
	graph.addEdge(0, 4, 15)
	graph.addEdge(1, 2, 15)
	graph.addEdge(1, 3, 2)
	graph.addEdge(2, 3, 1)
	graph.addEdge(2, 5, 5)
	graph.addEdge(3, 5, 12)
	graph.addEdge(4, 5, 10)
	got := graph.String()
	want := "(0, 1, 10) (0, 4, 15) (1, 2, 15) (1, 3, 2) (2, 3, 1) (2, 5, 5) (3, 5, 12) (4, 5, 10)"
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
