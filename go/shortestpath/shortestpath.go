package shortestpath

import (
	"container/list"
	"fmt"
	"math"
)

// Edge is a graph edge
type Edge struct {
	source   int
	terminal int
	weight   int
}

// NewEdge create a new edge
func NewEdge(s, t, w int) *Edge {
	return &Edge{
		source:   s,
		terminal: t,
		weight:   w,
	}
}

// Graph is struct for graph
type Graph struct {
	v   int
	adj []list.List
}

// NewGraph create a new graph
func NewGraph(v int) *Graph {
	return &Graph{
		v:   v,
		adj: make([]list.List, v),
	}
}

// AddEdge add a edge from s to t to the graph
func (g *Graph) AddEdge(s, t, w int) {
	g.adj[s].PushBack(NewEdge(s, t, w))
}

// ShortestPathByDijkstra get shortest from s to t path use Dijkstra
func (g *Graph) ShortestPathByDijkstra(s, t int) (int, string) {
	vertexDist := make([]int, g.v)
	preVertex := make([]int, g.v)
	visited := make([]bool, g.v)
	for i := range vertexDist {
		vertexDist[i] = math.MaxInt64
		preVertex[i] = -1
	}

	vertexDist[s] = 0
	for {
		v := getNotVisitedMinDistVertex(vertexDist, visited)
		// fmt.Println(v)
		if v == -1 {
			path := generatePath(preVertex, t)
			return vertexDist[t], path
		}

		for e := g.adj[v].Front(); e != nil; e = e.Next() {
			edge := e.Value.(*Edge)
			t, w := edge.terminal, edge.weight
			if vertexDist[v]+w < vertexDist[t] {
				vertexDist[t] = vertexDist[v] + w
				preVertex[t] = v
			}
		}
		visited[v] = true
	}
}

func getNotVisitedMinDistVertex(vertexDist []int, visited []bool) int {
	min := math.MaxInt64
	index := -1
	for v, dist := range vertexDist {
		if visited[v] == false && dist < min {
			min = dist
			index = v
		}
	}

	return index
}

func generatePath(preVertex []int, t int) string {
	result := ""
	q := t
	for preVertex[q] != -1 {
		result = fmt.Sprintf(" --> %d", q) + result
		q = preVertex[q]
	}
	result = fmt.Sprintf("%d", q) + result

	return result
}
