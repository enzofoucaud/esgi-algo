package main

import (
	hp "container/heap"
	"sort"
)

func NewGraph() *graph {
	graph := newGraph()

	// Metro A
	graph.addEdge("Bellecour", "Perrache", 2)
	graph.addEdge("Perrache", "Bellecour", 2)
	graph.addEdge("Cordeliers", "Bellecour", 1)
	graph.addEdge("Bellecour", "Cordeliers", 1)
	graph.addEdge("Hôtel de Ville", "Cordeliers", 1)
	graph.addEdge("Cordeliers", "Hôtel de Ville", 1)
	graph.addEdge("Charpennes", "Hôtel de Ville", 3)
	graph.addEdge("Hôtel de Ville", "Charpennes", 3)
	graph.addEdge("Laurent Bonnevay", "Charpennes", 5)
	graph.addEdge("Charpennes", "Laurent Bonnevay", 5)
	graph.addEdge("Vaulx-en-Velin", "Laurent Bonnevay", 1)
	graph.addEdge("Laurent Bonnevay", "Vaulx-en-Velin", 1)
	// Metro B
	graph.addEdge("Debourg", "Gare d’Oullins", 2)
	graph.addEdge("Gare d’Oullins", "Debourg", 2)
	graph.addEdge("Jean Macé", "Debourg", 2)
	graph.addEdge("Debourg", "Jean Macé", 2)
	graph.addEdge("Saxe Gambetta", "Jean Macé", 1)
	graph.addEdge("Jean Macé", "Saxe Gambetta", 1)
	graph.addEdge("Gare Part-Dieu", "Saxe Gambetta", 2)
	graph.addEdge("Saxe Gambetta", "Gare Part-Dieu", 2)
	graph.addEdge("Brotteaux", "Gare Part-Dieu", 1)
	graph.addEdge("Gare Part-Dieu", "Brotteaux", 1)
	graph.addEdge("Charpennes", "Brotteaux", 1)
	graph.addEdge("Brotteaux", "Charpennes", 1)
	// Metro C
	graph.addEdge("Croix-Paquet", "Cuire", 4)
	graph.addEdge("Cuire", "Croix-Paquet", 4)
	// Metro D
	graph.addEdge("Gare de Vaise", "Vieux Lyon", 3)
	graph.addEdge("Vieux Lyon", "Gare de Vaise", 3)
	graph.addEdge("Bellecour", "Vieux Lyon", 1)
	graph.addEdge("Vieux Lyon", "Bellecour", 1)
	graph.addEdge("Bellecour", "Guillotière", 1)
	graph.addEdge("Guillotière", "Bellecour", 1)
	graph.addEdge("Saxe Gambetta", "Guillotière", 1)
	graph.addEdge("Guillotière", "Saxe Gambetta", 1)
	graph.addEdge("Saxe Gambetta", "Grange Blanche", 4)
	graph.addEdge("Grange Blanche", "Saxe Gambetta", 4)
	graph.addEdge("Mermoz-Pinel", "Grange Blanche", 2)
	graph.addEdge("Grange Blanche", "Mermoz-Pinel", 2)
	graph.addEdge("Mermoz-Pinel", "Gare de Vénissieux", 2)
	graph.addEdge("Gare de Vénissieux", "Mermoz-Pinel", 2)

	return graph
}

type graph struct {
	nodes map[string][]destinationsedge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]destinationsedge)}
}

type vertex struct {
	cost  int
	edged []string
}

type kruskal struct {
	Vertices int
	graph    []MST
}

type MST struct {
	From   string
	To     string
	weight int
}

var (
	KMST        kruskal
	arrayVertex []string
)

type minimunPath []vertex

type priorityHeap struct {
	values *minimunPath
}

func newHeap() *priorityHeap {
	return &priorityHeap{values: &minimunPath{}}
}
func (pheap *priorityHeap) push(p vertex) {
	hp.Push(pheap.values, p)
}
func (pheap *priorityHeap) pop() vertex {
	temp := hp.Pop(pheap.values)
	return temp.(vertex)
}
func (pheap minimunPath) Len() int {
	return len(pheap)
}
func (pheap minimunPath) Less(i, j int) bool {
	return pheap[i].cost < pheap[j].cost
}
func (pheap minimunPath) Swap(i, j int) {
	pheap[i], pheap[j] = pheap[j], pheap[i]
}
func (pheap *minimunPath) Push(x interface{}) {
	*pheap = append(*pheap, x.(vertex))
}
func (pheap *minimunPath) Pop() interface{} {
	oldOne := *pheap
	news := len(oldOne)
	temp := oldOne[news-1]
	*pheap = oldOne[0 : news-1]
	return temp
}

type destinationsedge struct {
	destination string
	edgeweight  int
}

func (grph *graph) Dijkstra(origin, destiny string) (int, []string) {
	newHeap := newHeap()
	newHeap.push(vertex{cost: 0, edged: []string{origin}})
	visited := make(map[string]bool)
	for len(*newHeap.values) > 0 {
		pHeap := newHeap.pop()
		edge := pHeap.edged[len(pHeap.edged)-1]
		if visited[edge] {
			continue
		}
		if edge == destiny {
			return pHeap.cost, pHeap.edged
		}
		for _, e := range grph.getEdges(edge) {
			if !visited[e.destination] {
				newHeap.push(vertex{cost: pHeap.cost + e.edgeweight, edged: append([]string{}, append(pHeap.edged, e.destination)...)})
			}
		}
		visited[edge] = true
	}
	return 0, nil
}

func (grph *graph) Prims(origin, destiny string) (int, []MST) {
	var current = origin
	var list []MST
	var nodes []MST
	var cost = 0
	visited := make(map[string]bool)
	var i = 0
	for i < len(grph.nodes) && current != destiny {
		if !visited[current] {
			visited[current] = true
			for _, e := range grph.getEdges(current) {
				list = append(list, MST{From: current, To: e.destination, weight: e.edgeweight})
			}
			sort.SliceStable(list, func(i, j int) bool {
				return list[i].weight < list[j].weight
			})
			var popelement = list[0]
			list = list[1:]
			current = popelement.To
			if !visited[current] {
				nodes = append(nodes, popelement)
				cost += popelement.weight
			}
		} else {
			var popelement = list[0]
			list = list[1:]
			current = popelement.To
			if !visited[current] {
				nodes = append(nodes, popelement)
				cost += popelement.weight
			}
		}
		i++
	}
	return cost, nodes
}

func (grph *graph) addEdge(origin, destiny string, weight int) {
	grph.nodes[origin] = append(grph.nodes[origin], destinationsedge{destination: destiny, edgeweight: weight})
	grph.nodes[destiny] = append(grph.nodes[destiny], destinationsedge{destination: origin, edgeweight: weight})
	KMST.graph = append(KMST.graph, MST{From: origin, To: destiny, weight: weight})
	arrayVertex = append(arrayVertex, origin)
	arrayVertex = append(arrayVertex, destiny)
}

func (grph *graph) getEdges(node string) []destinationsedge {
	return grph.nodes[node]
}

func (grph *graph) Kruskal(origin string, destiny string, resultant []string) []MST {
	var (
		result []MST
		i      = 0
		e      = 0
		node   = 0
		parent = make(map[string]string)
		rank   = make(map[string]int)
	)
	sort.SliceStable(KMST.graph, func(i, j int) bool {
		return KMST.graph[i].weight < KMST.graph[j].weight
	})
	for node < KMST.Vertices {
		parent[resultant[node]] = resultant[node]
		rank[resultant[node]] = 0
		node++
	}
	for e < KMST.Vertices-1 {
		var (
			u = KMST.graph[i].From
			v = KMST.graph[i].To
			w = KMST.graph[i].weight
			x = findParent(parent, u)
			y = findParent(parent, v)
		)
		if x != y {
			e++
			result = append(result, MST{From: u, To: v, weight: w})
			union(parent, rank, x, y)
		}
		i++
		e++
	}
	return result
}

func union(parent map[string]string, rank map[string]int, x string, y string) {
	var xRoot = findParent(parent, x)
	var yRoot = findParent(parent, y)
	if rank[xRoot] < rank[yRoot] {
		parent[xRoot] = yRoot
	} else if rank[xRoot] > rank[yRoot] {
		parent[yRoot] = xRoot
	} else {
		parent[yRoot] = xRoot
		rank[xRoot] += 1
	}
}

func findParent(parent map[string]string, i string) string {
	if parent[i] == i {
		return i
	} else {
		return findParent(parent, parent[i])
	}
}
