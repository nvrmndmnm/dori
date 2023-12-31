package main

var list1 WeightedAdjacencyList = [][]WeightedEdge{
	{{1, 3}, {2, 1}},           // 0
	{{0, 3}, {2, 4}, {4, 1}},    // 1
	{{1, 4}, {3, 7}, {0, 1}},    // 2
	{{2, 7}, {4, 5}, {6, 1}},    // 3
	{{1, 1}, {3, 5}, {5, 2}},    // 4
	{{6, 1}, {4, 2}, {2, 18}},   // 5
	{{3, 1}, {5, 1}},            // 6
}

var list2 WeightedAdjacencyList = [][]WeightedEdge{
	{{1, 3}, {2, 1}},    // 0
	{{4, 1}},            // 1
	{{3, 7}},            // 2
	{},                  // 3
	{{1, 1}, {3, 5}, {5, 2}},    // 4
	{{2, 18}, {6, 1}},          // 5
	{{3, 1}},                   // 6
}

var matrix2 WeightedAdjacencyMatrix = [][]int{
	{0, 3, 1, 0, 0, 0, 0},    // 0
	{0, 0, 0, 0, 1, 0, 0},    // 1
	{0, 0, 7, 0, 0, 0, 0},    // 2
	{0, 0, 0, 0, 0, 0, 0},    // 3
	{0, 1, 0, 5, 0, 2, 0},    // 4
	{0, 0, 18, 0, 0, 0, 1},   // 5
	{0, 0, 0, 1, 0, 0, 1},    // 6
}

func intSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
