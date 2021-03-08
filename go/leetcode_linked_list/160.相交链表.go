package leetcode_linked_list_test

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB

	for nodeA != nodeB {
		if nodeA != nil {
			nodeA = nodeA.Next
		} else {
			nodeA = headA
		}

		if nodeB != nil {
			nodeB = nodeB.Next
		} else {
			nodeB = headB
		}
	}
	return nodeA
}

// @lc code=end
