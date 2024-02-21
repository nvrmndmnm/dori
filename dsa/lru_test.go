package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {

	lru := (&LRUCache[string, int]{}).NewLRUCache(3)

	if value := lru.Get("foo"); value != nil {
		t.Errorf("Expected nil for key 'foo', got %v", value)
	}

	lru.Update("foo", 42)
	if value := lru.Get("foo"); *value != 42 {
		t.Errorf("Expected value 42 for key 'foo', got %v", *value)
	}

	lru.Update("bar", 69)
	lru.Update("baz", 1337)
	lru.Update("ball", 420)

	if value := lru.Get("foo"); value != nil {
		t.Errorf("Expected nil for key 'foo' after eviction, got %v", value)
	}

	lru.Update("baz", 69420)
	if value := lru.Get("baz"); *value != 69420 {
		t.Errorf("Expected value 69420 for key 'baz', got %v", *value)
	}
}
