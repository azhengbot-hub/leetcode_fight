from typing import List


class Solution:
    def exchange(self, nums: List[int]) -> List[int]:
        res_list: List[int] = []

        for i in nums:
            if i % 2 == 0:
                res_list.append(i)
            else:
                res_list.insert(0, i)

        return res_list