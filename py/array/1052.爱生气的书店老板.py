from typing import List

#
# @lc app=leetcode.cn id=1052 lang=python3
#
# [1052] 爱生气的书店老板
#

# @lc code=start
class Solution:
    def maxSatisfied(self, customers: List[int], grumpy: List[int], X: int) -> int:
        left: int = 0
        right: int = left + X - 1

        res = 0
        for i in range(len(customers)):
            if left <= i <= right:
                res += customers[i]
                continue
            if grumpy[i] == 0:
                res += customers[i]

        max_res = res

        while right < len(customers) - 1:
            left += 1
            right += 1

            if grumpy[left - 1] == 1:
                res = res - customers[left - 1]

            if grumpy[right] == 1:
                res = res + customers[right]

            max_res = max(res, max_res)
        return max_res or res


# @lc code=end

s = Solution()

print(
    s.maxSatisfied(
        customers=[1, 0, 1, 2, 1, 1, 7, 5],
        ######### [0, 1, 0, 1, 0, 1, 0, 1]
        grumpy=[0, 1, 0, 1, 0, 1, 0, 1],
        X=3,
    )
)

print(s.maxSatisfied([10, 1, 7], [0, 0, 0], 2))
print(s.maxSatisfied([1], [0], 1))
print(s.maxSatisfied([2, 6, 6, 9], [0, 0, 1, 1], 1))  # 17
# assert (
#     s.maxSatisfied(
#         customers=[1, 0, 1, 2, 1, 1, 7, 5], grumpy=[0, 1, 0, 1, 0, 1, 0, 1], X=3
#     )
#     == 16
# )
