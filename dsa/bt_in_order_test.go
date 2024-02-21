package main

import "testing"

func TestInOrderTraversal(t *testing.T) {
	expected := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	result := InOrderTraversal(tree)

	for i := range expected {
		if i >= len(result) || expected[i] != result[i] {
			t.Errorf("Mismatch at index %d. Expected %d, got %d", i, expected[i], result[i])
		}
	}
}
