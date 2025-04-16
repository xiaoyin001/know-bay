package numberofislands

import "fmt"

/*
岛屿数量

给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1

示例 2：
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
*/

func NumIslandsTest() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
}

func numIslands(grid [][]byte) int {
	// 遍历所有的点，在发现是1的时候进行递归，将所有上下左右相邻的1都置为0
	// 然后记录这样递归了多少次

	mCnt := 0
	for i, vi := range grid {
		for j, vj := range vi {
			if vj == byte('1') {
				mCnt++
				// 开始进行周围的递归，将所有是1的值都置为0
				check(grid, i, j)
			}
		}
	}

	return mCnt
}

func check(grid [][]byte, i, j int) {
	// 检查其上下左右，如果有事1的将其置为0

	// 这点的上边
	if i-1 >= 0 && grid[i-1][j] == byte('1') {
		grid[i-1][j] = byte('0')
		check(grid, i-1, j)
	}

	// 这点的下边
	if i+1 < len(grid) && grid[i+1][j] == byte('1') {
		grid[i+1][j] = byte('0')
		check(grid, i+1, j)
	}

	// 这点的左边
	if j-1 >= 0 && grid[i][j-1] == byte('1') {
		grid[i][j-1] = byte('0')
		check(grid, i, j-1)
	}

	// 这点的右边
	if j+1 < len(grid[i]) && grid[i][j+1] == byte('1') {
		grid[i][j+1] = byte('0')
		check(grid, i, j+1)
	}
}
