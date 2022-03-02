package Search

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func shortestBridge(grid [][]int) int {

}

func isInGrid(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

// 使用深度优先搜索找到一个岛
func findOneIsland(queue, grid [][]int, x, y int) {
	if !isInGrid(grid, x, y) {
		return
	}
	if grid[x][y] == 0 {
		queue = append(queue, []int{x, y})
	} else if grid[x][y] == 1 {
		grid[x][y] = 2
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			findOneIsland(queue, grid, nx, ny)
		}
	}
}

func findBridge(queue, grid [][]int) int {

}
