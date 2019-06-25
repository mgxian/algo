package graph

import "testing"

func TestTopologicalSortUseKahnAlgorithm(t *testing.T) {
	graph := newDirectedWeightedGraph(6)
	graph.addEdge(0, 1, 10)
	graph.addEdge(0, 4, 15)
	graph.addEdge(1, 2, 15)
	graph.addEdge(1, 3, 2)
	graph.addEdge(2, 5, 5)
	graph.addEdge(3, 5, 12)
	graph.addEdge(3, 4, 8)
	graph.addEdge(3, 2, 6)
	graph.addEdge(1, 4, 5)
	graph.addEdge(4, 5, 10)
	got := graph.String()
	want := "(0, 1, 10) (0, 4, 15) (1, 2, 15) (1, 3, 2) (1, 4, 5) " +
		"(2, 5, 5) (3, 5, 12) (3, 4, 8) (3, 2, 6) (4, 5, 10)"
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}

	got = topologicalSortUseKahnAlgorithm(graph)
	want = "0 1 3 2 4 5"
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
