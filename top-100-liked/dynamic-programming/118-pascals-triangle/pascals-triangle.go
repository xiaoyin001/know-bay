package pascalstriangle

import "fmt"

/*
118-杨辉三角

给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。

示例 1:
输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

示例 2:
输入: numRows = 1
输出: [[1]]

提示:
1 <= numRows <= 30
*/

func GenerateTest() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	// 1、杨辉三角特点就是下一层的值等于上面层左右两边的和
	// 2、上面的1号特点是从第三层开始的，所以若干只需要返回2层，可以直接进行返回
	// 3、从第一层开始，每增加一层数组的长度也会随之增大
	// 4、所有行的第一个元素和最后一个元素都是1

	// 结合上面的情况需要将这个杨辉三角分成2部分进行
	// 第一部分：初始化杨辉三角
	// 第二部分：将初始化后的杨辉三角从第三层开始，计算数组中非收尾的值

	// 初始化杨辉三角
	mResult := make([][]int, 0, numRows)
	for i := 0; i < numRows; i++ {
		// 每一层的数组长度
		mLen := i + 1
		mResult = append(mResult, make([]int, mLen))
		// 初始化每一层的数组收尾值
		mResult[i][0] = 1
		mResult[i][mLen-1] = 1

		// 为了避免再次重复遍历一次，所以这里可以根据 i 进行判断，我们从第三行开始可以填充中间部分的内容
		// 前面1号特性总结出来就是：mResult[i][j] = mResult[i-1][j-1] + mResult[i-1][j]
		if i >= 2 {
			for j := 1; j < mLen-1; j++ {
				mResult[i][j] = mResult[i-1][j-1] + mResult[i-1][j]
			}
		}
	}

	return mResult
}
