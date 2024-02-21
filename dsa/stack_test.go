package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	stack.Push(5)
	stack.Push(7)
	stack.Push(9)

	if value := stack.Pop(); value != 9 {
		t.Errorf("Expected 9 but got %v", value)
	}
	if length := stack.Length; length != 2 {
		t.Errorf("Expected length 2 but got %d", length)
	}

	stack.Push(11)

	if value := stack.Pop(); value != 11 {
		t.Errorf("Expected 11 but got %v", value)
	}
	if value := stack.Pop(); value != 7 {
		t.Errorf("Expected 7 but got %v", value)
	}
	if value := stack.Peek(); value != 5 {
		t.Errorf("Expected 5 at the top but got %v", value)
	}
	if value := stack.Pop(); value != 5 {
		t.Errorf("Expected 5 but got %v", value)
	}
	if value := stack.Pop(); value != nil {
		t.Errorf("Expected nil but got %v", value)
	}
	if length := stack.Length; length != 0 {
		t.Errorf("Expected length 0 but got %d", length)
	}

	stack.Push(69)
	if value := stack.Peek(); value != 69 {
		t.Errorf("Expected 69 at the top but got %v", value)
	}
	if length := stack.Length; length != 1 {
		t.Errorf("Expected length 1 but got %d", length)
	}
}
