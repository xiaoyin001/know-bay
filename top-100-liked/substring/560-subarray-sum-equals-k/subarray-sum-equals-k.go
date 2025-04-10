package subarraysumequalsk

import "fmt"

/*
和为 K 的子数组

给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
子数组是数组中元素的连续非空序列。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2

示例 2：
输入：nums = [1,2,3], k = 3
输出：2


提示：
1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/

func SubarraySumTest() {
	fmt.Println(subarraySum([]int{-100, -1, 0, -1, 1, 0, 1, 100}, 0))
}

func subarraySum(nums []int, k int) int {
	mResult := 0

	// mTotalMap := make(map[int]int) // k: nums的idx               v: nums[0]~nums[i]的总和
	mPreMap := make(map[int]int) // k: nums[0]~nums[i]的总和   v: 出现的次数
	mTotle := 0
	for _, v := range nums {
		mTotle += v
		// mTotalMap[i] = mTotle
		// fmt.Println("mTotle=", mTotle)
		// map中的某一个值 = k 这就是区间内可行的，需要+1
		// map中的相同的数，所有两两之间相差=k都可以算一次

		// ================================
		// 检查一下有没有两两相差 k 的存在，如果有就表示可以有满足的次数
		mKey := mTotle - k
		mCnt, mOk := mPreMap[mKey]
		if mOk {
			// fmt.Println("可以加的次数：", mCnt)
			mResult += mCnt
		}

		// 将本次 nums[0]~nums[i]的总和 在 mPreMap 中进行更新
		mCnt = mPreMap[mTotle]
		// 这里不存在map中 mCnt = 0 所以这里可以忽略是否存在
		mCnt++
		mPreMap[mTotle] = mCnt
	}
	// fmt.Println("可以加的次数：", mPreMap[k])
	mResult += mPreMap[k]

	// fmt.Println("mPreMap=", mPreMap)

	return mResult
}
