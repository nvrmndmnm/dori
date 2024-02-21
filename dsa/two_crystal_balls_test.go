package main

import (
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	idx := rand.Intn(10000)

	data := make([]bool, 10000)
	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	t.Run("TwoCrystalBalls", func(t *testing.T) {
		result := twoCrystalBalls(data)
		if result != idx {
			t.Errorf("Expected %v but got %v", idx, result)
		}
	})

	t.Run("TwoCrystalBalls with length 821", func(t *testing.T) {
		result := twoCrystalBalls(make([]bool, 821))
		if result != -1 {
			t.Errorf("Expected -1 but got %v", result)
		}
	})
}
