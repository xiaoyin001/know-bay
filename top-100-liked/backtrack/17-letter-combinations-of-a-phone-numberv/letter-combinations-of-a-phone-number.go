package lettercombinationsofaphonenumberv

import (
	"fmt"
)

/*
电话号码的字母组合

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：
输入：digits = ""
输出：[]

示例 3：
输入：digits = "2"
输出：["a","b","c"]

提示：
0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。
*/

func LetterCombinationsTest() {
	fmt.Println(letterCombinations("34"))
}

var phoneNumMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var result []string

func letterCombinations(digits string) []string {
	result = []string{}

	// 根据数字的固定长度进行排列出来字母组合
	f1(digits, "", 0)

	return result
}

// 数字组合递归
func f1(digits, prefix string, idx int) {
	// 确保idx是在传入字符串的范围内的
	if idx < 0 || idx >= len(digits) {
		return
	}

	// 拿到当前数字对应的字母组合
	mStrs := phoneNumMap[string(digits[idx])]
	// 遍历对应的字母组合
	for _, v := range mStrs {
		// 如果当前不是最后一个数字的字母组合，就需要将当前数字的字母作为前缀
		// 将这个前缀传入给后面，直到最后一个数字的字母组合进行拼接
		if idx < len(digits)-1 {
			// 当前还不是最后一个数字的字母组合，继续递归到下一个数字对应的字母组
			// 需要拼接上当前的字符的座位前缀，并且告知下一组是那个数字
			f1(digits, prefix+string(v), idx+1)
		} else {
			// 这是最后一个数字对应的字母组，将这个字母拼接就是我们需要的结果之一
			result = append(result, prefix+string(v))
		}
	}
}
