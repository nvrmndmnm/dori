package main

import "testing"

func TestPreOrderTraversal(t *testing.T) {
	expected := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	result := PreOrderTraversal(tree)

	for i := range expected {
		if i >= len(result) || expected[i] != result[i] {
			t.Errorf("Mismatch at index %d. Expected %d, got %d", i, expected[i], result[i])
		}
	}
}