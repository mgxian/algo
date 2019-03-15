package graph

import "testing"

func TestGraph(t *testing.T) {
	aGraph := NewGraph(8)
	aGraph.AddEdge(0, 1)
	aGraph.AddEdge(1, 2)
	aGraph.AddEdge(0, 3)
	aGraph.AddEdge(1, 4)
	aGraph.AddEdge(2, 5)
	aGraph.AddEdge(3, 4)
	aGraph.AddEdge(4, 5)
	aGraph.AddEdge(4, 6)
	aGraph.AddEdge(5, 7)
	aGraph.AddEdge(6, 7)

	t.Log("get vertex's adjacent vertexes")
	t.Log(aGraph.GetVertexAdjVertex(0))
	t.Log(aGraph.GetVertexAdjVertex(4))
	t.Log(aGraph.GetVertexAdjVertex(5))

	t.Log("find path use BFS")
	t.Log(aGraph.BFS(0, 0))
	t.Log(aGraph.BFS(0, 1))
	t.Log(aGraph.BFS(0, 4))
	t.Log(aGraph.BFS(0, 7))

	t.Log("find path use DFS")
	t.Log(aGraph.DFS(0, 0))
	t.Log(aGraph.DFS(0, 1))
	t.Log(aGraph.DFS(0, 4))
	t.Log(aGraph.DFS(0, 7))
}
