package main

import "testing"

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap()
	if heap.Length != 0 {
		t.Errorf("Expected length: %d, got: %d", 0, heap.Length)
	}

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	if length := heap.Length; length != 8 {
		t.Errorf("Expected length: %d, got: %d", 8, length)
	}
	if value := heap.Delete(); value != 1 {
		t.Errorf("Expected value: %d, got: %d", 1, value)
	}
	if value := heap.Delete(); value != 3 {
		t.Errorf("Expected value: %d, got: %d", 3, value)
	}
	if value := heap.Delete(); value != 4 {
		t.Errorf("Expected value: %d, got: %d", 4, value)
	}
	if value := heap.Delete(); value != 5 {
		t.Errorf("Expected value: %d, got: %d", 5, value)
	}
	if length := heap.Length; length != 4 {
		t.Errorf("Expected length: %d, got: %d", 4, length)
	}
	if value := heap.Delete(); value != 7 {
		t.Errorf("Expected value: %d, got: %d", 7, value)
	}
	if value := heap.Delete(); value != 8 {
		t.Errorf("Expected value: %d, got: %d", 8, value)
	}
	if value := heap.Delete(); value != 69 {
		t.Errorf("Expected value: %d, got: %d", 69, value)
	}
	if value := heap.Delete(); value != 420 {
		t.Errorf("Expected value: %d, got: %d", 420, value)
	}
	if length := heap.Length; length != 0 {
		t.Errorf("Expected length: %d, got: %d", 0, length)
	}
}
