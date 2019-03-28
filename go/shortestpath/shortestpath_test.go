package shortestpath

import "testing"

func TestTopoSortByKahn(t *testing.T) {
	aGraph := NewGraph(6)
	aGraph.AddEdge(0, 1, 10)
	aGraph.AddEdge(0, 4, 15)
	aGraph.AddEdge(1, 2, 15)
	aGraph.AddEdge(1, 3, 2)
	aGraph.AddEdge(2, 5, 5)
	aGraph.AddEdge(3, 2, 1)
	aGraph.AddEdge(3, 5, 12)
	aGraph.AddEdge(4, 5, 10)
	dist, path := aGraph.ShortestPathByDijkstra(0, 5)
	t.Logf("dist: %d, path: %s", dist, path)

	dist, path = aGraph.ShortestPathByDijkstra(0, 2)
	t.Logf("dist: %d, path: %s", dist, path)

	dist, path = aGraph.ShortestPathByDijkstra(1, 2)
	t.Logf("dist: %d, path: %s", dist, path)
}
