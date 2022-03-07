package Search

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func shortestBridge(grid [][]int) int {
	queue := make([][]int, 0)
	findFirst := false
	for i := 0; i < len(grid) && !findFirst; i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				findOneIsland(grid, i, j, &queue)
				findFirst = true
				break
			}
		}
	}
	res := findBridge(grid, &queue)
	return res
}

func isInGrid(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

// 使用深度优先搜索找到一个岛
func findOneIsland(grid [][]int, x, y int, queue *[][]int) {
	if !isInGrid(grid, x, y) {
		return
	}
	if grid[x][y] == 0 {
		*queue = append(*queue, []int{x, y})
	} else if grid[x][y] == 1 {
		grid[x][y] = 2
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			findOneIsland(grid, nx, ny, queue)
		}
	}
}

// 使用广度优先搜索找到最短的搭桥距离
func findBridge(grid [][]int, queue *[][]int) int {
	var res int
	for len(*queue) > 0 {
		res++
		for j := len(*queue); j > 0; j-- {
			x, y := (*queue)[0][0], (*queue)[0][1]
			*queue = (*queue)[1:]
			for k := 0; k < 4; k++ {
				nx, ny := x+dir[k][0], y+dir[k][1]
				if isInGrid(grid, nx, ny) {
					if grid[nx][ny] == 1 {
						return res
					}
					if grid[nx][ny] == 2 {
						continue
					}
					*queue = append(*queue, []int{nx, ny})
					grid[nx][ny] = 2
				}
			}
		}
	}
	return res
}

//func main() {
//	grid := [][]int{
//		{0,1,0,0,0,0},
//		{0,1,1,1,0,0},
//		{0,0,0,0,0,0},
//		{0,0,0,0,0,0},
//		{0,0,0,0,0,0},
//		{1,1,0,0,0,0},
//	}
//	res := shortestBridge(grid)
//	println(res)
//}
