#
# @lc app=leetcode.cn id=11 lang=python3
#
# [11] 盛最多水的容器
#

# @lc code=start
from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        right = len(height) - 1
        left = 0
        max_res = 0

        while right != left:
            if height[left] < height[right]:
                res = (right - left) * height[left]
                left += 1
            else:
                res = (right - left) * height[right]
                right -= 1

            max_res = max(max_res, res)

        return max_res


# @lc code=end

s = Solution()

print(s.maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]))  # 49
