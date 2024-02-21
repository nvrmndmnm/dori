package main

import "testing"

func TestDijkstraList(t *testing.T) {
	// list1
	// {{1, 3}, {2, 1}}     
	// {{0, 3}, {2, 4}, {4, 1}}
	// {{1, 4}, {3, 7}, {0, 1}}
	// {{2, 7}, {4, 5}, {6, 1}}
	// {{1, 1}, {3, 5}, {5, 2}}
	// {{6, 1}, {4, 2}, {2, 18}}
	// {{3, 1}, {5, 1}}

	expectedPath := []int{0, 1, 4, 5, 6}
	result := DijkstraList(0, 6, list1)

	if !intSlicesEqual(result, expectedPath) {
		t.Errorf("Expected path %v, but got %v", expectedPath, result)
	}
}
