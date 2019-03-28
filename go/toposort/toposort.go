package toposort

import (
	"container/list"
)

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
func (g *Graph) AddEdge(s, t int) {
	g.adj[s].PushBack(t)
}

// TopoSortByKahn toposort use kahn
func (g *Graph) TopoSortByKahn() []int {
	inDegree := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		for e := g.adj[i].Front(); e != nil; e = e.Next() {
			v := e.Value.(int)
			inDegree[v]++
		}
	}

	queue := list.New()
	for i, v := range inDegree {
		if v == 0 {
			queue.PushBack(i)
		}
	}

	result := make([]int, 0)
	for queue.Len() != 0 {
		e := queue.Front()
		v := e.Value.(int)
		result = append(result, v)
		queue.Remove(e)

		for e := g.adj[v].Front(); e != nil; e = e.Next() {
			t := e.Value.(int)
			inDegree[t]--
			if inDegree[t] == 0 {
				queue.PushBack(t)
			}
		}
	}

	return result
}

// TopoSortByDFS toposort use dfs
func (g *Graph) TopoSortByDFS() (result []int) {
	inverseAdj := make([]list.List, g.v)
	for i := 0; i < g.v; i++ {
		for e := g.adj[i].Front(); e != nil; e = e.Next() {
			v := e.Value.(int)
			inverseAdj[v].PushBack(i)
		}
	}

	visited := make([]bool, g.v)
	for i := 0; i < g.v; i++ {
		if visited[i] == false {
			visited[i] = true
			result = dfs(i, inverseAdj, visited, result)
		}
	}

	return
}

func dfs(vertex int, inverseAdj []list.List, visited []bool, result []int) []int {
	for e := inverseAdj[vertex].Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		if visited[v] == true {
			continue
		}
		dfs(v, inverseAdj, visited, result)
	}
	result = append(result, vertex)
	return result
}
