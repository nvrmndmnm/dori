package main

import (
	"strings"
	"testing"
)

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeResult := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9,  y: 4},
		{x: 8,  y: 4},
		{x: 7,  y: 4},
		{x: 6,  y: 4},
		{x: 5,  y: 4},
		{x: 4,  y: 4},
		{x: 3,  y: 4},
		{x: 2,  y: 4},
		{x: 1,  y: 4},
		{x: 1,  y: 5},
	}

	result := Solve(maze, "x", Point{x: 10, y: 0}, Point{x: 1, y: 5})

	drawnResult := drawPath(maze, result)

	expectedResult := drawPath(maze, mazeResult)

	if drawnResult != expectedResult {
		t.Errorf("Expected \n%s \nbut got \n%s", expectedResult, drawnResult)
	}
}

func drawPath(data []string, path []Point) string {
	data2 := make([][]rune, len(data))
	for i, row := range data {
		data2[i] = []rune(row)
	}

	for _, p := range path {
		if p.y >= 0 && p.y < len(data2) && p.x >= 0 && p.x < len(data2[p.y]) {
			data2[p.y][p.x] = '*'
		}
	}

	output := make([]string, len(data2))
	for i, row := range data2 {
		output[i] = string(row)
	}

	return strings.Join(output, "\n")
}