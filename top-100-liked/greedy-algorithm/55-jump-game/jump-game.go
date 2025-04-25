package jumpgame

import "fmt"

/*
跳跃游戏

给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

示例 2：
输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。

提示：
1 <= nums.length <= 104
0 <= nums[i] <= 105
*/

func CanJumpTest() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	// fmt.Println(canJump([]int{3, 2, 1, 0, 0, 1, 4}))
}

func canJump(nums []int) bool {
	// 如果数组内容为空或只有1个元素，可以直接表示为可以到达最后一个
	if len(nums) <= 1 {
		return true
	}

	// 遍历数组，记录当前数组可以到达的最大位置
	// 默认先从第一个数开始，然后记录当前的最大数
	mMaxIdx := nums[0]

	// 如果第一个就已经是0了，则表示无法到达
	if mMaxIdx == 0 {
		return false
	}

	// 可以提前判断一手如果当前第一个就已经超过了数组长度，就表示可以到达最后一个
	if mMaxIdx >= len(nums)-1 {
		return true
	}

	for i := 1; i < len(nums); i++ {
		// 如果遍历的下标大于当前可以到达的最大下标位置，则表示当前无法到达这个目标，可以直接判定失败了
		if i > mMaxIdx {
			return false
		}

		// 判断是否需要更新当前的最大值
		if i+nums[i] > mMaxIdx {
			mMaxIdx = i + nums[i]
		}

		// 判断当前的最大值是否已经可以到达数组的最后一位，如果可以表示可以顺利到达
		if mMaxIdx >= len(nums)-1 {
			return true
		}
	}

	return false
}
