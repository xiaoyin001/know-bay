package reverselinkedlist

import "fmt"

/*
反转链表

给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：
输入：head = [1,2]
输出：[2,1]

示例 3：
输入：head = []
输出：[]

提示：
链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000

进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseListTest() {
	mRootNode := &ListNode{
		Val:  0,
		Next: nil,
	}

	mNode := mRootNode
	for i := 1; i < 5; i++ {
		mTmpNone := &ListNode{Val: i}
		mNode.Next = mTmpNone
		mNode = mTmpNone
	}

	mCurNode := reverseList(mRootNode)

	for mCurNode != nil {
		fmt.Println(mCurNode.Val)

		mCurNode = mCurNode.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	// 使用一个变量，将列表进行反转

	//  0 1 2 3 4
	// 可以理解为反转后根节点的Next节点
	// var mOldNode *ListNode = nil
	mCurNode := head
	head = nil
	for mCurNode != nil {
		// 将当前节点置为根节点 例如：遍历到0的时候将其设置为跟节点
		// head = mCurNode
		// 记录当前根节点的Next节点 例如：遍历到0的时候这里记录1这个节点
		mNextNode := mCurNode.Next
		// 将当前节点的Next指向新的反转后的节点 例如：遍历到0的时候这里将 0.Next = mOldNode(nil)
		// mCurNode.Next = mOldNode
		mCurNode.Next = head
		// 标记当前反转后的节点
		// mOldNode = head
		// mOldNode = mCurNode
		head = mCurNode
		// 下一个节点的遍历
		if mNextNode == nil {
			return mCurNode
		}
		mCurNode = mNextNode
	}

	return head
}

func reverseList2(head *ListNode) *ListNode {
	var mFunc func(*ListNode) *ListNode

	var mResultNode *ListNode = nil

	mFunc = func(node *ListNode) *ListNode {
		// fmt.Println("222", node)

		if node == nil || node.Next == nil {
			// 这里可能是 最后一个节点的Next or 最后一个节点 也可以说是初始节点，在没有的情况下
			mResultNode = node

			return node
		}

		// 假设链表是按照 01234 的顺序传入
		// 这里返回的mNode就是 4321 可以理解传入什么返回什么 mNode == node.Next
		mNode := mFunc(node.Next)
		// fmt.Println("111 -- ", node)
		// 这里将开始进行链表顺序反转
		mNode.Next = node
		// 反转后将 node.Next 置空
		node.Next = nil

		// 然后就继续，传入什么返回什么
		return node
	}

	mFunc(head)

	return mResultNode
}
