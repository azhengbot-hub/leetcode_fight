package leetcode_array_test

import (
	"fmt"
	common "leetcode/go/common"
	"testing"
)

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *common.ListNode) *common.ListNode {
	var pre *common.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// @lc code=en

func TestFunc206(t *testing.T) {
	res := reverseList(common.TestListNode)
	fmt.Println(res)
}
