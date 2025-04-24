package kthlargestelementinanarray

import "fmt"

/*
数组中的第K个最大元素

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

提示：
1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/

func FindKthLargestTest() {
	// fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 5))
}

func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}

	// 这里我需要使用堆的概念进行处理

	// 首先是需要将当前的切片转换成堆的存放形式
	// 假设 根节点是 r 其对应的左右子节点分别是 r*2+1 和 r*2+2
	f3_1(nums, len(nums))
	fmt.Println("nums: ", nums)

	// 将切片转换成最大堆以后就需要开始将堆顶的元素推出，推出后整体的结构是需要调整的
	// 最大堆：父节点 >= 左右子节点
	// 推出的思路：先将顶点与最后一个子节点进行调换，然后再对调换后的顶点向下调整，直到满足堆的特性
	// 被调换到最后的节点（老的顶点）就可以直接去掉了

	// 然后就是看需要返回第几大的数，然后就推出顶点 k-1 次
	for i := 1; i < k; i++ {
		// 将堆顶部的数据与最后一个数据进行交换
		nums[0], nums[len(nums)-i] = nums[len(nums)-i], nums[0]
		// 然后将交换后的新的顶部数据进行调整，校验当前的数组结构是否满足最大堆特性
		f3_2(nums, 0, len(nums)-i)
	}

	return nums[0]
}

// 将数组置为最大堆
// a: 存放堆节点的数组
// heapSize: 堆的大小
func f3_1(a []int, heapSize int) {
	// 从最小的非叶子节点开始，往前遍历所有的节点 r
	// 最后一个节点节点的下标是 heapSize-1，那么其父节点的下标就是 (heapSize-1)/2
	for i := (heapSize - 1) / 2; i >= 0; i-- {
		// 这里i以及其之前的节点都是非叶子节点，也都可以是父节点

		// 这里就开始检查这个父节点，是否需要与其子节点进行调换，为了保证最大堆的特性
		f3_2(a, i, heapSize)
	}
}

// 将数组中的当前父节点进行调整，保证其满足最大堆的特性
// a: 存放堆节点的数组
// i: 当前父节点的下标
// heapSize: 堆的大小
func f3_2(a []int, i, heapSize int) {
	// 先找到这个节点对应的左右子节点的下标
	// 在使用的时候需要检查左右子节点是否存在，就是需要与 heapSize 进行比较
	// 只要给进来的 i不是小于0的 l 和 r 就不可能小于0的
	l := i*2 + 1
	r := i*2 + 2

	// 判断当前节点及其左右子节点，将最大值作为当前节点
	// 然后将变幻的节点继续进行递归，看其子节点是否满足，直接没有子节点或都满足

	// 弄一个临时变量，作为当前根节点 curI
	curI := i

	// 先与左子节点进行比较，如果左子节点大于当前节点，则将当前根节点与左子节点进行调换
	if l < heapSize && a[l] > a[curI] {
		curI = l
	}

	// 再与右子节点进行比较，如果右子节点大于当前节点，则将当前根节点与右子节点进行调换
	if r < heapSize && a[r] > a[curI] {
		curI = r
	}

	// 判断当前根节点是否发生了变化，如果发生了变化，将当前节点 i 与 替换后的新根节点 curI 进行交换
	// 交换后继续检查交换后子节点作为根节点与其子节点进行比较，是否满足堆的特性
	if curI != i {
		a[i], a[curI] = a[curI], a[i]
		f3_2(a, curI, heapSize)
	}

	// 如果没有进行交换，说明当前的数组是满足堆的特性
}

// ========================================================================
// 这是官方的答案

// findKthLargest 函数用于找到数组中第 k 大的元素
func findKthLargest2(nums []int, k int) int {
	// heapSize 初始化为数组的长度，表示堆的大小
	heapSize := len(nums)
	// 调用 buildMaxHeap 函数将数组转化为最大堆
	buildMaxHeap(nums, heapSize)
	// 从数组的末尾开始，逐步将堆顶元素（即当前最大值）与末尾元素交换
	// 每次交换后，堆的大小减1，并对堆顶元素进行下沉操作以维持最大堆性质
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		// 交换堆顶元素和当前末尾元素
		nums[0], nums[i] = nums[i], nums[0]
		// 堆的大小减1
		heapSize--
		// 对堆顶元素进行下沉操作，以维持最大堆性质
		maxHeapify(nums, 0, heapSize)
	}
	// 经过 k-1 次交换后，堆顶元素即为第 k 大的元素
	return nums[0]
}

// buildMaxHeap 构建最大堆
// 参数a是一个整数切片，表示待构建堆的数组
// 参数heapSize表示堆的大小，即数组中有效元素的个数
func buildMaxHeap(a []int, heapSize int) {
	// 从最后一个非叶子节点开始向前遍历，对每个节点进行最大堆化
	// 非叶子节点的索引为heapSize/2 - 1
	for i := heapSize/2 - 1; i >= 0; i-- {
		// 对当前节点进行最大堆化
		maxHeapify(a, i, heapSize)
	}
}

// maxHeapify 维护最大堆性质，确保以节点i为根的子树满足父节点值 >= 子节点值
// 参数:
//
//	a         : 待调整的数组（堆的底层存储）
//	i         : 当前需要调整的节点索引
//	heapSize  : 堆的有效大小（数组前 heapSize 个元素属于堆）
func maxHeapify(a []int, i, heapSize int) {
	// 计算左右子节点的索引（数组从0开始存储完全二叉树）
	l := i*2 + 1 // 左子节点索引 = 2*i + 1
	r := i*2 + 2 // 右子节点索引 = 2*i + 2
	largest := i // 记录当前节点、左子节点、右子节点中的最大值索引

	// 如果左子节点存在且大于当前最大值节点
	if l < heapSize && a[l] > a[largest] {
		largest = l // 更新最大值为左子节点
	}

	// 如果右子节点存在且大于当前最大值节点（可能是更新后的左子节点）
	if r < heapSize && a[r] > a[largest] {
		largest = r // 更新最大值为右子节点
	}

	// 如果最大值不是当前节点，说明需要调整
	if largest != i {
		// 交换当前节点与最大值节点的值
		a[i], a[largest] = a[largest], a[i]
		// 递归调整被交换后的子树（新的子节点可能违反堆性质）
		// 例如：原a[largest]被下移到i位置后，可能比它的新子节点小
		maxHeapify(a, largest, heapSize)
	}
	// 若 largest == i，说明当前子树已满足堆性质，递归终止
}

// ========================================================================
// 这种写法超时了

func findKthLargest_1(nums []int, k int) int {
	// 第一个想法就是将 nums 进行排序，然后返回

	// 有没有可能 k 是大于 nums的长度

	// 有一个新的思路就是创建一个固定长度的数组，数组长度就是 k ，然后遍历 nums ,开始的几个直接按照大小的顺序进行插入
	// 然后就是遍历完 nums 得到答案
	mArr := make([]int, 0, k) // 按照从大到小的顺序进行排序
	mMin := -104
	for _, v := range nums {
		if len(mArr) < k {
			// 还没塞满继续塞
			// fix: 没有满的时候直接往后插入的，导致排序出现问题，所以这里需要在前面进行插入
			mArr = append([]int{v}, mArr...)
		} else {
			// 塞满了，如果当前的值比最小的还小，可以直接忽略
			if mMin > v {
				continue
			}

			// 塞满了，将最小的代替，为下面的排序做准备
			mArr[0] = v
		}

		// 进行排序
		f1(mArr)

		// 排完以后重新设置最小数
		mMin = mArr[0]
	}

	return mMin
}

func f1(nums []int) {
	// 只需要将第一个数进行比较，一直往后交换，直到找到对应的位置

	for i := 1; i < len(nums); i++ {
		// 当前的跟前一个进行比较
		if nums[i-1] <= nums[i] {
			// 如果前面一个 <= 当前的，就不需要继续了，现在已经算是好的
			return
		} else {
			// 如果前面一个 > 当前的，就需要进行冒泡交换，然后继续下次比较，直接完成
			nums[i-1], nums[i] = nums[i], nums[i-1]
		}
	}
}

// ========================================================================
