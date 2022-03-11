package main

import (
	"fmt"
	"os"
)

func main() {
	// Récupère les paramètres start et end
	var (
		s     = os.Args[1]
		e     = os.Args[2]
		start = s[8:]
		end   = e[6:]
	)

	// Création du métro de Lyon
	graph := NewGraph()

	// Dijkstra
	fmt.Println("----- Algorithme de Dijkstra -----")
	costDijkstra, dijkstraPath := graph.Dijkstra(start, end)
	if len(dijkstraPath) < 1 {
		fmt.Println("La source ou destination ne peuvent pas s'atteindre")
	} else {
		fmt.Println("Le chemin avec Dijkstra coûte ", costDijkstra)
		for i, s := range dijkstraPath {
			fmt.Print(s)
			if i != len(dijkstraPath)-1 {
				fmt.Print("  ->  ")
			}
		}
		fmt.Println()
	}
	// Prims
	fmt.Println("----- Algorithme de Prims -----")
	costPrim, primsPath := graph.Prims(start, end)
	if len(primsPath) < 1 {
		fmt.Println("La source ou destination ne peuvent pas s'atteindre")
	} else {
		fmt.Println("Le chemin avec Prims coûte ", costPrim)
		for _, p := range primsPath {
			fmt.Println("Départ de", p.From, "vers", p.To, "avec un poid de", p.weight)
		}
	}
	// Kruskal
	fmt.Println("----- Algorithme de Kruskal -----")
	KMST.Vertices = len(graph.nodes)
	resultantVertex := make([]string, 0)
	check := make(map[string]int)
	for _, val := range arrayVertex {
		check[val] = 1
	}
	for letter := range check {
		resultantVertex = append(resultantVertex, letter)
	}
	kruskalResult := graph.Kruskal(start, end, resultantVertex)
	var minimumCost = 0
	for _, v := range kruskalResult {
		minimumCost += v.weight
		fmt.Println(v.From, " -> ", v.To, " == ", v.weight)
	}
	fmt.Println("Le chemin minimum avec Kruskal est de", minimumCost)
}
