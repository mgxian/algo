package graph

import (
	"strconv"
	"strings"

	"github.com/mgxian/algo/go/practice/list"
)

func topologicalSortUseKahnAlgorithm(graph *directedWeightedGraph) string {
	indegrees := make(map[int]int, 0)
	visited := make(map[int]bool, 0)

	for node := range graph.adjacencyList {
		indegrees[node] = 0
	}

	for node, edges := range graph.adjacencyList {
		visited[node] = false
		p := edges.Front()
		for p != nil {
			e := p.Value.(edge)
			indegrees[e.destnation]++
			p = p.Next()
		}
	}

	unVisitedZeroIndegreeNode := func() int {
		for node, indegree := range indegrees {
			if indegree == 0 && visited[node] != true {
				return node
			}
		}
		return -1
	}

	doneNode := func(node int) {
		p := graph.adjacencyList[node].Front()
		for p != nil {
			dnode := p.Value.(edge).destnation
			indegrees[dnode]--
			p = p.Next()
		}
	}

	result := new(strings.Builder)
	for {
		node := unVisitedZeroIndegreeNode()
		if node == -1 {
			break
		}
		visited[node] = true
		result.WriteString(" " + strconv.Itoa(node))
		doneNode(node)
	}

	if strings.HasPrefix(result.String(), " ") {
		return result.String()[1:]
	}
	return result.String()
}

func dfs(node int, reversedAjacencyList map[int]list.List, visited map[int]bool, output *strings.Builder) {
	if visited[node] {
		return
	}

	p := reversedAjacencyList[node].Front()
	for p != nil {
		e := p.Value.(edge)
		if !visited[e.destnation] {
			dfs(e.destnation, reversedAjacencyList, visited, output)
		}
		p = p.Next()
	}

	visited[node] = true
	output.WriteString(" " + strconv.Itoa(node))
}

func topologicalSortUseDFSAlgorithm(graph *directedWeightedGraph) string {
	result := new(strings.Builder)
	visited := make(map[int]bool, len(graph.adjacencyList))
	reversedAjacencyList := make(map[int]list.List, len(graph.adjacencyList))
	for i := 0; i < len(graph.adjacencyList); i++ {
		reversedAjacencyList[i] = list.NewSinglyLinkedList()
		visited[i] = false
	}

	for node, edges := range graph.adjacencyList {
		p := edges.Front()
		for p != nil {
			e := p.Value.(edge)
			reversedAjacencyList[e.destnation].PushBack(edge{e.destnation, node, -1})
			p = p.Next()
		}
	}

	for i := 0; i < len(graph.adjacencyList); i++ {
		dfs(i, reversedAjacencyList, visited, result)
	}

	if strings.HasPrefix(result.String(), " ") {
		return result.String()[1:]
	}
	return result.String()
}
