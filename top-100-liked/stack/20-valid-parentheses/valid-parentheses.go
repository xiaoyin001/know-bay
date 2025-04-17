package validparentheses

import "fmt"

/*
有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


示例 1：
输入：s = "()"
输出：true

示例 2：
输入：s = "()[]{}"
输出：true

示例 3：
输入：s = "(]"
输出：false

示例 4：
输入：s = "([])"
输出：true

提示：
= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

func IsValidTest() {
	fmt.Println(isValid("({[]})"))
}

func isValid(s string) bool {
	mLen := len(s)
	// 如果不成成对的直接不满足
	if mLen%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 作为栈，用来存放当前匹配的
	stack := []byte{}

	// 遍历字符串
	for i := 0; i < mLen; i++ {
		// 如果找到就说明是反括号
		if pairs[s[i]] > 0 {
			// 如果栈是空的，现在出来反括号，就是不满足的
			// 如果当前栈的最上面一位是不是当前反括号对应的，也是不满足的
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 如果就是对应的正反括号，就将栈最上面的进行出栈操作
			stack = stack[:len(stack)-1]
		} else {
			// 如果是正括号之类的就将其入栈
			stack = append(stack, s[i])
		}
	}

	// 遍历完字符串，如果栈是空的，就说明括号都是正反对应的
	return len(stack) == 0
}
