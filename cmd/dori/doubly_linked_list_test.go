package main

import "testing"

func TestListOperations(t *testing.T) {
	list := NewDoublDoublyLinkedList()

	list.Append(5)
	list.Append(7)
	list.Append(9)

	if value := list.Get(2); value != 9 {
		t.Errorf("Got %d, expected 9", value)
	}

	if removed := list.RemoveAt(1); removed != 7 {
		t.Errorf("Got %d, expected 7", removed)
	}

	if length := list.Length; length != 2 {
		t.Errorf("Got %d, expected 2", length)
	}

	list.Append(11)

	if removed := list.RemoveAt(1); removed != 9 {
		t.Errorf("Got %d, expected 9", removed)
	}

	if removed := list.Remove(9); removed != nil {
		t.Errorf("Got %v, expected nil", removed)
	}

	if removed := list.RemoveAt(0); removed != 5 {
		t.Errorf("Got %d, expected 5", removed)
	}

	if removed := list.RemoveAt(0); removed != 11 {
		t.Errorf("Got %d, expected 11", removed)
	}

	if length := list.Length; length != 0 {
		t.Errorf("Got %d, expected 0", length)
	}

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	if value := list.Get(2); value != 5 {
		t.Errorf("Got %d, expected 5", value)
	}

	if value := list.Get(0); value != 9 {
		t.Errorf("Got %d, expected 9", value)
	}

	if removed := list.Remove(9); removed != 9 {
		t.Errorf("Got %d, expected 9", removed)
	}

	if length := list.Length; length != 2 {
		t.Errorf("Got %d, expected 2", length)
	}

	if value := list.Get(0); value != 7 {
		t.Errorf("Got %d, expected 7", value)
	}
}
