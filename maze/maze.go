package main

import (
	"fmt"
	"os"
)

func createMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	var row, col int
	_, _ = fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, _ = fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}
func main() {
	maze := createMaze("maze/maze.in")
	for _, row := range maze {
		fmt.Println(row)
	}
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

type point struct {
	x, y int
}

func (p point) add(dir point) point {
	return point{
		p.x + dir.x,
		p.y + dir.y,
	}
}

func (p point) at(maze [][]int) (int, bool) {
	if p.x < 0 || p.x > len(maze)-1 {
		return 0, false
	}
	if p.y < 0 || p.y > len(maze[p.x])-1 {
		return 0, false
	}
	return maze[p.x][p.y], true
}

var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	queue := []point{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			if curStep, ok := cur.at(steps); ok {
				steps[next.x][next.y] = curStep + 1
				queue = append(queue, next)
			}
		}
	}
	return steps
}
