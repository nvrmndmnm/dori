package main

type WeightedEdge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]WeightedEdge
type WeightedAdjacencyMatrix [][]int

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
