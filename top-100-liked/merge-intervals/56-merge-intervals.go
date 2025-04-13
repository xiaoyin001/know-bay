package mergeintervals

import "fmt"

/*
合并区间

以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

提示：
1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/

func MergeTest() {
	// fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	// fmt.Println(merge([][]int{{1, 4}, {1, 5}}))
	// fmt.Println(merge([][]int{{1, 4}, {1, 4}}))
	// fmt.Println(merge([][]int{{1, 4}, {0, 4}}))
	// fmt.Println(merge([][]int{{1, 4}, {0, 0}}))
	// fmt.Println(merge([][]int{{1, 4}, {0, 0}, {5, 5}, {8, 8}}))
	fmt.Println(merge([][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}))
}

func merge(intervals [][]int) [][]int {
	mResult := make([][]int, 0)

	mMax := 0
	for _, v := range intervals {
		if mMax < v[0] {
			mMax = v[0]
		}
		if mMax < v[1] {
			mMax = v[1]
		}
	}

	mTmpMap := make(map[int]struct{}) // k: v[0]*1000000 + v[1]
	mArr := make([]int, mMax+1)
	for _, v := range intervals {
		// fix：发现有完全重叠的区间的情况处理会有问题
		// 在这里需要先去重一次，如果这个区间已经存在了，就需要给忽略
		// 因为这里的数最大是 10的4次方
		mKey := v[0]*100000 + v[1]
		_, mOk := mTmpMap[mKey]
		if mOk {
			continue
		}
		mTmpMap[mKey] = struct{}{}

		mArr[v[0]] += 1
		mArr[v[1]] -= 1
	}

	mCurIdx := -1
	mState := 0
	for idx, v := range mArr {

		if v > 0 && mState == 0 {
			mResult = append(mResult, []int{idx})
			mCurIdx++
		} else if v < 0 {
			if mState == 1 {
				mTmpV := mResult[mCurIdx]
				mTmpV = append(mTmpV, idx)
				mResult[mCurIdx] = mTmpV
			} else if mState+v == 0 {
				mTmpV := mResult[mCurIdx]
				mTmpV = append(mTmpV, idx)
				mResult[mCurIdx] = mTmpV
			}
		} else {
			// v == 0
			// 判断是否有这组，如果有就加入返回
			mKey := idx*100000 + idx
			_, mOk := mTmpMap[mKey]
			if mOk {
				// FIX: 修复中间存在的 start stop相同的情况，需要跳过
				if mCurIdx == -1 || len(mResult[mCurIdx]) == 2 {
					mResult = append(mResult, []int{idx, idx})
					mCurIdx++
				}
			}
		}

		mState += v
	}

	return mResult
}
