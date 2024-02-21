package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(5)
	queue.Enqueue(7)
	queue.Enqueue(9)

	if value := queue.Deque(); value != 5 {
		t.Errorf("Expected 5 but got %d", value)
	}
	if length := queue.Length; length != 2 {
		t.Errorf("Expected length 2 but got %d", length)
	}

	queue.Enqueue(11)

	if value := queue.Deque(); value != 7 {
		t.Errorf("Expected 7 but got %d", value)
	}
	if value := queue.Deque(); value != 9 {
		t.Errorf("Expected 9 but got %d", value)
	}
	if value := queue.Peek(); value != 11 {
		t.Errorf("Expected 11 at the front but got %d", value)
	}
	if value := queue.Deque(); value != 11 {
		t.Errorf("Expected 11 but got %d", value)
	}
	if value := queue.Deque(); value != nil {
		t.Errorf("Expected nil but got %v", value)
	}
	if length := queue.Length; length != 0 {
		t.Errorf("Expected length 0 but got %d", length)
	}

	queue.Enqueue(69)

	if value := queue.Peek(); value != 69 {
		t.Errorf("Expected 69 at the front but got %d", value)
	}
	if length := queue.Length; length != 1 {
		t.Errorf("Expected length 1 but got %d", length)
	}
}