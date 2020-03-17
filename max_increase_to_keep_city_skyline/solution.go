package max_increase_to_keep_city_skyline

func maxIncreaseKeepingSkyline(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	maxItemsInRow := make([]int, row)
	maxItemsInColumn := make([]int, column)
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			//找出每行最大的
			if grid[i][j] > maxItemsInRow[i] {
				maxItemsInRow[i] = grid[i][j]
			}
			//找出每列最大的
			if grid[i][j] > maxItemsInColumn[j] {
				maxItemsInColumn[j] = grid[i][j]
			}
		}
	}

	increment := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			//根据某个位置的行和列，取到最小的，否则天际线会超出
			min := maxItemsInRow[i]
			if maxItemsInColumn[j] < min {
				min = maxItemsInColumn[j]
			}
			increment += min - grid[i][j]
		}
	}

	return increment
}
