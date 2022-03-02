package Search

func solveNQueens(n int) [][]string {
	var res [][]string
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	col, dia1, dia2 := map[int]bool{}, map[int]bool{}, map[int]bool{}
	var backtrack func(queens []int, n, row int, col, dia1, dia2 map[int]bool)
	backtrack = func(queens []int, n, row int, col, dia1, dia2 map[int]bool) {
		if row == n {
			board := generateBoard(queens, n)
			res = append(res, board)
			return
		}
		for i := 0; i < n; i++ {
			if col[i] {
				continue
			}
			diagonal1 := row + i
			if dia1[diagonal1] {
				continue
			}
			diagonal2 := row - i
			if dia2[diagonal2] {
				continue
			}
			queens[row] = i
			col[i] = true
			dia1[diagonal1], dia2[diagonal2] = true, true
			backtrack(queens, n, row+1, col, dia1, dia2)
			queens[row] = -1
			delete(col, i)
			delete(dia1, diagonal1)
			delete(dia2, diagonal2)
		}
	}
	backtrack(queens, n, 0, col, dia1, dia2)
	return res
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
