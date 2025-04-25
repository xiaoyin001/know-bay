package besttimetobuyandsellstock

import "fmt"

/*
买卖股票的最佳时机

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

示例 1：
输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

示例 2：
输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。

提示：
1 <= prices.length <= 105
0 <= prices[i] <= 104
*/

func MaxProfitTest() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	// 记录当前的最大收益
	// 然后一直找最小值，然后计算最大的收益
	// 这里有一个潜规则就是只能买入在前面，卖出在后面

	mMaxIn := 0 // 当前的最大收益

	mMinVal := prices[0] // 股票的最低点，初始假设第一天是最低的

	// 开始遍历每天的行情，从第二天开始，看是否能够卖出
	for i := 1; i < len(prices); i++ {
		// 如果发现更低的价格，就更新最低点，然后继续后面的遍历
		if prices[i] < mMinVal {
			mMinVal = prices[i]
			continue
		}

		// 如果当前的价格比最低点高，就可以考虑是否能出手，与当前最大收益比较是否更高
		if mMaxIn < prices[i]-mMinVal {
			mMaxIn = prices[i] - mMinVal
		}
	}

	return mMaxIn
}
