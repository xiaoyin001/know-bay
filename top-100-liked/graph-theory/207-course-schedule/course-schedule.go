package courseschedule

import "fmt"

/*
课程表

你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

示例 2：
输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

提示：
1 <= numCourses <= 2000
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对 互不相同
*/

func CanFinishTest() {
	// fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	// fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{}))
}

// ================================================================================
// 虽然这种方法是可以通过官方的测试，但是这种解法好像不是那么那么的优秀

var studyMap map[int][]int

func canFinish(numCourses int, prerequisites [][]int) bool {
	studyMap = make(map[int][]int)

	// 先遍历 prerequisites 存map，然后遍历map，看是否可以全部修完
	for _, v := range prerequisites {
		mSlice, mOk := studyMap[v[0]]
		if !mOk {
			mSlice = make([]int, 0)
		}
		mSlice = append(mSlice, v[1])
		studyMap[v[0]] = mSlice
	}

	// 每次都需要判断map是否已经空了，如果空了就表示可以修完
	for i := 0; i < numCourses; i++ {
		if len(studyMap) == 0 {
			return true
		}

		// 当前需要学习的，第几次深度递归，最多只能递归map长度次
		if !f1(i, 0, numCourses) {
			return false
		} else {
			delete(studyMap, i)
		}
	}

	return len(studyMap) == 0
}

func f1(needStudy, cnt, max int) bool {
	if cnt >= max {
		return false
	}

	mSlice, mOk := studyMap[needStudy]
	if !mOk {
		return true
	}

	mIdx := 0
	mLen := len(mSlice)
	for mIdx < mLen {
		if !f1(mSlice[mIdx], cnt+1, max) {
			// 有问题的，直接返回false
			return false
		} else {
			mIdx++
		}
	}

	return true
}

// ================================================================================
