package binarytreelevelordertraversal

import "fmt"

/*
二叉树的层序遍历

给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]

示例 2：
输入：root = [1]
输出：[[1]]

示例 3：
输入：root = []
输出：[]

提示：
树中节点数目在范围 [0, 2000] 内
-1000 <= Node.val <= 1000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrderTest() {
	mRootNode := &TreeNode{
		Val:   0,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(levelOrder(mRootNode))
}

var uResult [][]int

func levelOrder(root *TreeNode) [][]int {
	uResult = make([][]int, 0)

	rangeTree(0, root)

	return uResult
}

func rangeTree(curIdx int, node *TreeNode) {
	if node == nil {
		return
	}

	if len(uResult) <= curIdx {
		uResult = append(uResult, []int{})
	}

	uResult[curIdx] = append(uResult[curIdx], node.Val)

	rangeTree(curIdx+1, node.Left)
	rangeTree(curIdx+1, node.Right)
}
