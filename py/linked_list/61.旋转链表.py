#
# @lc app=leetcode.cn id=61 lang=python3
#
# [61] 旋转链表
#

# @lc code=start
# Definition for singly-linked list.

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def rotateRight(self, head: ListNode, k: int) -> Optional[ListNode]:
        cnt = 1
        if not head:
            return None
        if not head.next:
            return head

        old_tail = head

        while old_tail.next:
            old_tail = old_tail.next
            cnt += 1

        old_tail.next = head

        new_tail = head
        for i in range(cnt - k % cnt - 1):
            new_tail = new_tail.next

        new_head = new_tail.next

        new_tail.next = None

        return new_head


# @lc code=end
