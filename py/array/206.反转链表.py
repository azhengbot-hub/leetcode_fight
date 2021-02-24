#
# @lc app=leetcode.cn id=206 lang=python3
#
# [206] 反转链表
#

# @lc code=start
# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: ListNode) -> Optional[ListNode]:
        pre = None
        cur = head
        while cur:
            next = cur.next  # 先把原来cur.next位置存起来
            cur.next = pre
            pre = cur
            cur = next
        return pre


# @lc code=end


l = ListNode(
    1,
    ListNode(
        2,
        ListNode(3, ListNode(4, ListNode(5, None))),
    ),
)

s = Solution()
print(s.reverseList(l))