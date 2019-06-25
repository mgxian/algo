package graph

import (
	"strconv"
	"strings"
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
