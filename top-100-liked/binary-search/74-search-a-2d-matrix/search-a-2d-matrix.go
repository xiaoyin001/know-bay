package searcha2dmatrix

import "fmt"

/*
搜索二维矩阵

给你一个满足下述两条属性的 m x n 整数矩阵：
每行中的整数从左到右按非严格递增顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

示例 1：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true

示例 2：
输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104
*/

func SearchMatrixTest() {
	// fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	// fmt.Println(searchMatrix([][]int{{1}}, 1))
	fmt.Println(searchMatrix([][]int{{1, 3}}, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	// 先遍历行，然后行的第一个元素与 target 进行比较
	mU, mD := 0, len(matrix)-1 // Up Down 行的表示
	mIdx := (mD-mU)/2 + mU
	for mD-mU > 0 {
		mNums := matrix[mIdx]
		if mNums[0] == target {
			return true
		} else if matrix[mIdx][0] > target {
			// 将区间设置为 [mU, mIdx-1]
			mD = mIdx - 1
		} else {
			// 这里需要将当前行的最后一个位数进行比较，如果 最后一位数大于 target 说明这个数就在这行内，否则就继续进行二分查找
			if mNums[len(mNums)-1] == target {
				return true
			} else if mNums[len(mNums)-1] > target {
				// 对应的数字就是该行区间内，跳出当前循环，进入行的二分查找
				break
			} else {
				// 说明不在这行 mU 所在行，将当前区间调整为 [mIdx+1, mD]
				mU = mIdx + 1
			}
		}

		mIdx = (mD-mU)/2 + mU
	}

	mNums := matrix[mIdx]
	// fix：1 begin 出现 [][]int{{1}}, 1 的情况无法进行，这里做一手只有1个参数的情况判断
	if len(mNums) == 1 {
		return mNums[0] == target
	}
	// fix: 1 end
	mL, mR := 0, len(mNums)-1 // Left Right 列的表示
	// fix: 2 begin [][]int{{1, 3}}, 3 ，原来这里是 mR-mL > 0
	for mR-mL >= 0 {
		// fix: 2 end
		mIdx = (mR-mL)/2 + mL
		if mNums[mIdx] == target {
			return true
		} else if mNums[mIdx] > target {
			// 将区间设置为 [mL, mIdx-1]
			mR = mIdx - 1
		} else {
			// 将区间设置为 [mIdx+1, mR]
			mL = mIdx + 1
		}
	}

	return false
}
