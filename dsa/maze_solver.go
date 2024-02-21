package main

type Point struct {
	x int
	y int
}

var dir [][]int = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func walk(maze []string,
	wall string,
	curr Point,
	end Point,
	seen [][]bool,
	path []Point,
) (bool, []Point) {
	if curr.x < 0 || curr.x >= len(maze[0]) ||
		curr.y < 0 || curr.y >= len(maze) {
		return false, nil
	}

	if string(maze[curr.y][curr.x]) == wall {
		return false, nil
	}

	if curr.x == end.x && curr.y == end.y {
		return true, append(path, end)
	}

	if seen[curr.y][curr.x] {
		return false, nil
	}

	seen[curr.y][curr.x] = true
	path = append(path, curr)

	for i := 0; i < len(dir); i++ {
		x, y := dir[i][0], dir[i][1]

		pass, path := walk(maze, wall, Point{
			x: curr.x + x,
			y: curr.y + y,
		}, end, seen, path)
		
		if pass {
			return true, path
		}
	}

	return false, path[1:]
}

func Solve(maze []string,
	wall string,
	start Point,
	end Point,
) []Point {
	var seen [][]bool = make([][]bool, 0)
	var path []Point

	for i := 0; i < len(maze); i++ {
		seen = append(seen, make([]bool, len(maze[0])))
	}

	_, path = walk(maze, wall, start, end, seen, path)

	return path
}
