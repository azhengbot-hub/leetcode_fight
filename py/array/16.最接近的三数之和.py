#
# @lc app=leetcode.cn id=16 lang=python3
#
# [16] 最接近的三数之和
#

# @lc code=start
from typing import List


class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        res = nums[0] + nums[1] + nums[2]
        nums = sorted(nums)
        for i in range(len(nums)):
            left = i + 1
            right = len(nums) - 1
            while left < right:
                val = nums[i] + nums[left] + nums[right]
                if abs(val - target) < abs(res - target):
                    res = val
                if val > target:
                    right -= 1
                elif val < target:
                    left += 1
                else:
                    return val

        return res


# @lc code=end

s = Solution()

print(s.threeSumClosest([1, 1, 1, 0], -100))
