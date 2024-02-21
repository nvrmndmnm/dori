package main

import "testing"

func TestPostOrderTraversal(t *testing.T) {
	expected := []int{7, 5, 15, 10, 21, 29, 49, 45, 30, 50, 20}
	result := PostOrderTraversal(tree2)

	for i := range expected {
		if i >= len(result) || expected[i] != result[i] {
			t.Errorf("Mismatch at index %d. Expected %d, got %d", i, expected[i], result[i])
		}
	}
}
