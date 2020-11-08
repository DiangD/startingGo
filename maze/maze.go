package main

import (
	"fmt"
	"os"
)

//BFS迷宫算法

//createMaze 生成迷宫
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

//add 计算可探索范围内的点的地址
func (p point) add(dir point) point {
	return point{
		p.x + dir.x,
		p.y + dir.y,
	}
}

//at 判断是否在迷宫内
func (p point) at(maze [][]int) (int, bool) {
	if p.x < 0 || p.x > len(maze)-1 {
		return 0, false
	}
	if p.y < 0 || p.y > len(maze[p.x])-1 {
		return 0, false
	}
	return maze[p.x][p.y], true
}

//dirs 4个方向
var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

//walk 探索
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	queue := []point{start}

	for len(queue) > 0 {
		//当前位置
		cur := queue[0]
		queue = queue[1:]
		//探索到终点
		if cur == end {
			break
		}
		//探索四个方向
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)
			//遇到墙
			if !ok || val == 1 {
				continue
			}
			//已经探索过
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			//探索回到起点
			if next == start {
				continue
			}
			//获取到上一个探索的点在steps的值，++
			if curStep, ok := cur.at(steps); ok {
				steps[next.x][next.y] = curStep + 1
				//push进队列
				queue = append(queue, next)
			}
		}
	}
	return steps
}
