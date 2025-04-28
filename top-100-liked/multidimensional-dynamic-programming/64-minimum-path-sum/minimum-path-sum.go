package minimumpathsum

import "fmt"

/*
64-最小路径和

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

示例 1：
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。

示例 2：
输入：grid = [[1,2,3],[4,5,6]]
输出：12

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 200
*/

func MinPathSumTest() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}

func minPathSum(grid [][]int) int {
	// 没看解题思路前在想是不是能通过一条路径选择最优的，但是好像不行，会有极端情况让这条前面较优路径变得很差
	// 看了解题思路后，这里就是需要根据走完所有的路线找到最终最小的
	// 就相当于是要遍历所有的数据，由于只能往右和往下走，所以这里就可以一边遍历一边计算到当前点位的最小路径值
	// 计算当前点位的最小路径值就是根据左边和上面点位的最小路径值加上当前点位的路径
	// 然后取较小的，就是当前当前路径的最小路径值

	for i, rows := range grid {
		for j, v := range rows {
			// 在计算最小路径值的时候需要考虑边界问题
			// 只需要考虑当前点位是不是顶部顶点和最左边第一列的值即可，其余的就可以直接计算

			// 如果是第一行的数据 i == 0 如果是顶部数据的情况就只用考虑左边的值
			// 如果是第一列的数据 j == 0 如果是最左边第一列的数据就只用考虑上面的值
			// 如果 i和j 都是0，直接跳过，因为这里就是初始值
			// 总结下来就是 grid[i][j] == min(grid[i-1][j], grid[i][j-1])

			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				// 第一行的数据
				grid[i][j] = grid[i][j-1] + v
				continue
			}

			if j == 0 {
				// 第一列的数据
				grid[i][j] = grid[i-1][j] + v
				continue
			}

			// 非第一行和第一列的数据
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + v
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}
