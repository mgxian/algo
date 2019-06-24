package graph

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mgxian/algo/go/practice/list"
	"github.com/mgxian/algo/go/practice/queue"
)

type undirectedGraph struct {
	adjacencyMatrix [][]bool
}

func newUndirectedGraph(n int) *undirectedGraph {
	ug := new(undirectedGraph)
	ug.adjacencyMatrix = make([][]bool, n)
	for i := 0; i < n; i++ {
		ug.adjacencyMatrix[i] = make([]bool, n)
	}
	return ug
}

func (ug *undirectedGraph) addEdge(x, y int) {
	ug.adjacencyMatrix[x][y] = true
}

func (ug *undirectedGraph) String() string {
	result := new(strings.Builder)
	for x, line := range ug.adjacencyMatrix {
		for y, hasEdge := range line {
			if hasEdge {
				source := strconv.Itoa(x)
				destnation := strconv.Itoa(y)
				result.WriteString(" " + source + "--->" + destnation)
			}
		}
	}

	if strings.HasPrefix(result.String(), " ") {
		return result.String()[1:]
	}

	return result.String()
}

type undirectedGraphWithAdjacencyList struct {
	adjacencyList []list.List
}

func newUndirectedGraphWithAdjacencyList(n int) *undirectedGraphWithAdjacencyList {
	ug := new(undirectedGraphWithAdjacencyList)
	ug.adjacencyList = make([]list.List, n)
	for i := 0; i < n; i++ {
		ug.adjacencyList[i] = list.NewSinglyLinkedList()
	}
	return ug
}

func (ug *undirectedGraphWithAdjacencyList) addEdge(x, y int) {
	ug.adjacencyList[x].PushBack(y)
}

func (ug *undirectedGraphWithAdjacencyList) String() string {
	result := new(strings.Builder)

	for x, nodes := range ug.adjacencyList {
		p := nodes.Front()
		for p != nil {
			source := strconv.Itoa(x)
			destnation := strconv.Itoa(p.Value.(int))
			result.WriteString(" " + source + "--->" + destnation)
			p = p.Next()
		}
	}

	if strings.HasPrefix(result.String(), " ") {
		return result.String()[1:]
	}

	return result.String()
}

type edge struct {
	source     int
	destnation int
	weight     int
}

func (e edge) String() string {
	return fmt.Sprintf("(%d, %d, %d)", e.source, e.destnation, e.weight)
}

type directedWeightedGraph struct {
	adjacencyList []list.List
}

func newDirectedWeightedGraph(n int) *directedWeightedGraph {
	uwg := new(directedWeightedGraph)
	uwg.adjacencyList = make([]list.List, n)
	for i := 0; i < n; i++ {
		uwg.adjacencyList[i] = list.NewSinglyLinkedList()
	}
	return uwg
}

func (uwg *directedWeightedGraph) addEdge(x, y, w int) {
	uwg.adjacencyList[x].PushBack(edge{x, y, w})
}

func (uwg *directedWeightedGraph) String() string {
	result := new(strings.Builder)

	for _, nodes := range uwg.adjacencyList {
		p := nodes.Front()
		for p != nil {
			e := p.Value.(edge)
			result.WriteString(" " + e.String())
			p = p.Next()
		}
	}

	if strings.HasPrefix(result.String(), " ") {
		return result.String()[1:]
	}

	return result.String()
}

func (uwg *directedWeightedGraph) breadthFirstTraversal() string {
	result := new(strings.Builder)

	root := 0
	visited := make(map[int]bool, 0)
	q := queue.NewListQueue(len(uwg.adjacencyList))
	q.Enqueue(root)

	for !q.IsEmpty() {
		node := q.Dequeue().(int)
		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = true
		result.WriteString(" " + strconv.Itoa(node))
		p := uwg.adjacencyList[node].Front()
		for p != nil {
			q.Enqueue(p.Value.(edge).destnation)
			p = p.Next()
		}
	}

	if strings.HasPrefix(result.String(), " ") {
		return result.String()[1:]
	}

	return result.String()
}

func (uwg *directedWeightedGraph) dfs(node int, visited map[int]bool, result *strings.Builder) {
	if _, ok := visited[node]; ok {
		return
	}

	result.WriteString(" " + strconv.Itoa(node))
	visited[node] = true
	p := uwg.adjacencyList[node].Front()
	for p != nil {
		uwg.dfs(p.Value.(edge).destnation, visited, result)
		p = p.Next()
	}
}

func (uwg *directedWeightedGraph) depthFirstTraversal() string {
	result := new(strings.Builder)

	visited := make(map[int]bool, 0)
	uwg.dfs(0, visited, result)

	if strings.HasPrefix(result.String(), " ") {
		return result.String()[1:]
	}

	return result.String()
}
