package main

type WeightedEdge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]WeightedEdge
type WeightedAdjacencyMatrix [][]int
