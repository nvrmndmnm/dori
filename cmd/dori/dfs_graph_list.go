package main

func walkGraph(graph WeightedAdjacencyList, curr int, needle int, seen []bool, path []int) (bool, []int) {

	if seen[curr] {
		return false, nil
	}

	seen[curr] = true

	path = append(path, curr)
	if curr == needle {
		return true, path
	}

	list := graph[curr]
	for i := 0; i < len(list); i++ {
		edge := list[i]

		pass, path := walkGraph(graph, edge.To, needle, seen, path)
		if pass {
			return true, path
		}
	}

	return false, path[1:]
}

func DFSGraphList(graph WeightedAdjacencyList, source int, needle int) []int {
	var seen []bool = make([]bool, len(graph))
	var path []int

	_, path = walkGraph(graph, source, needle, seen, path)

	if len(path) == 0 {
		return nil
	}

	return path
}
