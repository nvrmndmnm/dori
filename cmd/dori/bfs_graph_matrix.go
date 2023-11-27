package main

func BFSGraphMatrix(graph WeightedAdjacencyMatrix, source int, needle int) []int {
	var seen []bool = make([]bool, len(graph))
	var prev []int = make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	var q []int
	q = append(q, source)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := 0; i < len(graph); i++ {
			if adjs[i] == 0 {
				continue
			}

			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q = append(q, i)
		}
		seen[curr] = true
	}

	if prev[needle] == -1 {
		return nil
	}

	curr := needle
	out := []int{}

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}
	out = reverse(out)
	res := append([]int{source}, out...)

	return res
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
