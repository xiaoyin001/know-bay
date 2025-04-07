package twosum

import "fmt"

/*
两数之和

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]


提示：
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
*/

func TwoSumTest() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum2(nums, target))
}

func twoSum(nums []int, target int) []int {
	mResult := make([]int, 2)

	// 用时间换空间，用map先存一份，然后遍历查找
	mNumsMap := make(map[int]int)
	for k, v := range nums {
		mNumsMap[v] = k
	}

	for k, v := range nums {
		mNum := target - v
		if vk, ok := mNumsMap[mNum]; ok {
			if k == vk {
				continue
			}

			mResult[0] = k
			mResult[1] = vk
			break
		}

	}

	return mResult
}

// 改进，应该一边遍历一边检查是否有合适的，而不是先存一份再遍历
func twoSum2(nums []int, target int) []int {
	mResult := make([]int, 2)

	mNumsMap := make(map[int]int)
	for k, v := range nums {
		mNum := target - v
		if vk, ok := mNumsMap[mNum]; ok {

			mResult[0] = k
			mResult[1] = vk
			break
		}

		mNumsMap[v] = k
	}

	return mResult
}
