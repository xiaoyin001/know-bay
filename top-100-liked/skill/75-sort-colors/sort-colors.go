package sortcolors

import "fmt"

/*
75-颜色分类

给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
必须在不使用库内置的 sort 函数的情况下解决这个问题。

示例 1：
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

示例 2：
输入：nums = [2,0,1]
输出：[0,1,2]

提示：
n == nums.length
1 <= n <= 300
nums[i] 为 0、1 或 2

进阶：
你能想出一个仅使用常数空间的一趟扫描算法吗？
*/

func SortColorsTest() {
	mNums := []int{2, 0, 2, 1, 1, 0}
	sortColors(mNums)
	fmt.Println(mNums)
}

func sortColors(nums []int) {
	// 我最初的想法就是用额外的切片才存放不同的颜色，遍历完了再重新将数组拼接起来

	// 然后看到题目后面进阶的提示，又想到了直接使用常亮计数又多少个0，1，2，然后重新赋值即可

	// 白色数量
	mWhiteCount := 0
	// 蓝色数量
	mBlueCount := 0
	// 红色不用计数，剩下的都是红色的
	for i, v := range nums {
		if v == 1 {
			mWhiteCount++
		} else if v == 2 {
			mBlueCount++
		}

		// 统计过的就可以给数组重置了，这样也省去了后面赋值红色0
		nums[i] = 0
	}

	for i := 0; i < mBlueCount; i++ {
		// 填充蓝色，在最后几位，只需要将数组后面的数填充2即可
		nums[len(nums)-1-i] = 2
	}

	for i := 0; i < mWhiteCount; i++ {
		// 填充白色，在蓝色前面几位，只需要将数组中间的数填充1即可
		nums[len(nums)-1-mBlueCount-i] = 1
	}
}
