package setmatrixzeroes

import "fmt"

/*
矩阵置零

给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

示例 1：
输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]

示例 2：
输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

提示：
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1
*/

func SetZeroesTest() {
	// mParam := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	mParam := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(mParam)
	fmt.Println(mParam)
}

func setZeroes(matrix [][]int) {
	// 使用一个map存横纵坐标为0的
	// 横坐标K: i   中坐标K: j+1000
	mMap := make(map[int]struct{})
	for i, vs := range matrix {
		for j, v := range vs {
			if v == 0 {
				mMap[i] = struct{}{}
				mMap[j+1000] = struct{}{}
			}
		}
	}

	// 遍历map将满足的横纵对应的值设置为0
	for k := range mMap {
		if k >= 1000 {
			j := k - 1000
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		} else {
			i := k
			// 下面这种写法等同于上面 for i := 0; i < len(matrix); i++ { 的写法
			for j := range matrix[i] {
				matrix[i][j] = 0
			}
		}
	}

	// 我怀疑这里面会有坑比如切片的长度会不一样？
}
