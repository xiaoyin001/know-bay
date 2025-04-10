package findallanagramsinastring

import "fmt"

/*
找到字符串中所有字母异位词
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词：字母异位词是通过重新排列不同单词或短语的字母而形成的单词或短语，并使用所有原字母一次

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

提示:
1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母
*/

func FindAnagramsTest() {
	// fmt.Println(findAnagrams("baa", "aa"))
	fmt.Println(findAnagrams2("baa", "aa"))
}

// 暴力算法，过于暴力，执行超时，需要进行优化
func findAnagrams(s string, p string) []int {
	mResult := make([]int, 0)
	if len(p) > len(s) {
		return mResult
	}

	// 将s拥有的参数进行计数，看每个字母出现过几次，方便后续检查每组是否满足
	mPMap := make(map[byte]uint16, 0)
	for i := 0; i < len(p); i++ {
		mV, mOk := mPMap[p[i]]
		if !mOk {
			mV = 0
		}

		mV++
		mPMap[p[i]] = mV
	}

	// 使用窗口的方式进行遍历所有满足长度的字符串
	mPLen := len(p)
	mCurIdx := 0
	for mCurIdx+mPLen <= len(s) {
		mStr := s[mCurIdx : mCurIdx+mPLen]
		mTmpPMap := make(map[byte]uint16, 0)
		for k, v := range mPMap {
			mTmpPMap[k] = v
		}

		for i := 0; i < len(mStr); i++ {
			mV, mOk := mTmpPMap[mStr[i]]
			if !mOk {
				// 已经存在不匹配的了
				break
			}

			if mV > 1 {
				mV--
				mTmpPMap[mStr[i]] = mV
			} else {
				delete(mTmpPMap, mStr[i])
			}
		}

		if len(mTmpPMap) == 0 {
			mResult = append(mResult, mCurIdx)
		}

		mCurIdx++
	}

	return mResult
}

// 看了官方的方式一，改进一下实现
func findAnagrams2(s string, p string) []int {
	mResult := make([]int, 0)
	mPLen := len(p)
	mSLen := len(s)
	if mPLen > mSLen {
		return mResult
	}

	// 将s拥有的参数进行计数，看每个字母出现过几次，方便后面比较
	// 为了更加翻遍的比较内容，就需要使用数组进行填充
	// 这里为什么要用26这个长度，因为英文字母是26个，这里可能出现的只会是这些
	mPArr := [26]int16{}
	mSArr := [26]int16{}
	// 这里算是给定滑动窗口的初始大小，同时也是为了方便后续的滑动的时候进行比较
	for i := 0; i < mPLen; i++ {
		// 这里 -'a' 是为了让26个字母能有在数组中找到对应的地方
		mPArr[p[i]-'a']++
		// 这里直接使用s[i]是因为前面已经判断过了，在这里字符串p的长度一定不会大于字符串s的长度
		mSArr[s[i]-'a']++
	}

	// 上面在进行初始滑动窗口大小的时候就已经算是将最左侧的窗口数据填充可，这里可以直接进行比较一下
	if mPArr == mSArr {
		mResult = append(mResult, 0)
	}

	// 这里移动窗口，看后面满足的字符串是否满足要求
	// 下面这个for里面的操作就是将左侧即将移出滑动窗口的数据从 mSArr 中去掉，然后将右侧即将进入的数据加入 mSArr 中
	// 然后再与p对应的 mPArr 进行比较，如果数组相同，就表示是满足的
	// mSLen-mPLen 是为了确保边界
	for i := 0; i < mSLen-mPLen; i++ {
		mSArr[s[i]-'a']--
		mSArr[s[i+mPLen]-'a']++

		// 上面已经移动成功，这里就可以判断滑动窗口内的数据是否满足了
		if mSArr == mPArr {
			mResult = append(mResult, i+1)
		}
	}

	return mResult
}
