# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def getKthFromEnd(self, head: ListNode, k: int) -> Optional[ListNode]:

        cnt = 1
        res_dic = {cnt: head}
        while head.next:
            cnt += 1
            head = head.next
            res_dic[cnt] = head

        return res_dic[cnt - k + 1]
