package graph

import (
	"github.com/mgxian/algo/go/linkedlist/singly"
	"github.com/mgxian/algo/go/queue"
)

// Graph is a graph data structure
type Graph struct {
	vertexCount int
	adjList     []*singly.LinkedList
}

// NewGraph create a new graph
func NewGraph(vertexCount int) *Graph {
	adjList := make([]*singly.LinkedList, vertexCount)
	for i := 0; i < vertexCount; i++ {
		adjList[i] = singly.NewLinkedList()
	}
	return &Graph{
		vertexCount: vertexCount,
		adjList:     adjList,
	}
}

// AddEdge add a edge into the graph
func (g *Graph) AddEdge(s, t int) {
	g.adjList[s].InsertToTail(t)
	g.adjList[t].InsertToTail(s)
}

// GetVertexAdjVertex get vertex's edges
func (g *Graph) GetVertexAdjVertex(vertex int) []int {
	var ret []int
	for _, v := range g.adjList[vertex].Traverse() {
		v, _ := v.(int)
		ret = append(ret, v)
	}
	return ret
}

// BFS is breadth first search
func (g *Graph) BFS(s, t int) (vertexes []int) {
	if s == t {
		vertexes = append(vertexes, s, t)
		return
	}

	visited := make([]bool, g.vertexCount)
	prev := make([]int, g.vertexCount)
	for i := range prev {
		prev[i] = -1
	}
	visited[s] = true

	aQueue := queue.NewSimpleQueue(g.vertexCount)
	aQueue.EnQueue(s)
	for !aQueue.IsEmpty() {
		vertex, _ := aQueue.DeQueue()
		currentVertex := vertex.(int)
		for _, v := range g.GetVertexAdjVertex(currentVertex) {
			if !visited[v] {
				prev[v] = currentVertex
				if v == t {
					return getPathVertexes(t, prev)
				}
				visited[v] = true
				aQueue.EnQueue(v)
			}
		}
	}
	return
}

// DFS is depth first search
func (g *Graph) DFS(s, t int) (vertexes []int) {
	if s == t {
		vertexes = append(vertexes, s, t)
		return
	}

	visited := make([]bool, g.vertexCount)
	prev := make([]int, g.vertexCount)
	for i := range prev {
		prev[i] = -1
	}

	g.recurDFS(s, t, visited, prev)
	return getPathVertexes(t, prev)
}

func (g *Graph) recurDFS(s, t int, visited []bool, prev []int) {
	if s == t {
		return
	}

	visited[s] = true
	for _, v := range g.GetVertexAdjVertex(s) {
		if !visited[v] {
			prev[v] = s
			g.recurDFS(v, t, visited, prev)
		}
	}
}

func getPathVertexes(t int, prev []int) (vertexes []int) {
	var tmp []int
	tmp = append(tmp, t)
	p := prev[t]
	for p != -1 {
		tmp = append(tmp, p)
		p = prev[p]
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		vertexes = append(vertexes, tmp[i])
	}

	return
}
