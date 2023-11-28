package main

func hasUnvisited(seen []bool, dists []int) bool {
	for i, s := range seen {
		if !s && dists[i] < int(^uint(0) >> 1) {
			return true
		}
	}

	return false
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	lowestDistance := int(^uint(0) >> 1)

	for i := 0; i < len(seen); i++ {
		if seen[i] {
			continue
		}

		if lowestDistance > dists[i] {
			lowestDistance = dists[i]
			idx = i
		}
	}

	return idx
}

func DijkstraList(source int, sink int, arr WeightedAdjacencyList) []int {
	seen := make([]bool, len(arr))
	prev := make([]int, len(arr))
	dists := make([]int, len(arr))

	for i := range dists {
		prev[i] = -1
		dists[i] = int(^uint(0) >> 1)
	}

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		
		seen[curr] = true
		adjs := arr[curr]
	
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]

			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + edge.Weight
			
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	var out []int
	curr := sink

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}
	
	out = append(out, source)

	return reverse(out)
}
