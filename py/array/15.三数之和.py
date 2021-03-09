#
# @lc app=leetcode.cn id=15 lang=python3
#
# [15] 三数之和
#
from typing import List

# @lc code=start
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:

        res = []
        len_nums = len(nums)
        nums = sorted(nums)

        if (not nums) or (len_nums < 3):
            return res

        for i in range(len_nums):
            if nums[i] > 0:
                break

            if i > 0 and nums[i] == nums[i - 1]:
                continue

            left = i + 1
            right = len_nums - 1

            while left < right:
                if nums[i] + nums[left] + nums[right] == 0:
                    res.append([nums[i], nums[left], nums[right]])
                    while left < right and nums[left] == nums[left + 1]:
                        left += 1
                    while left < right and nums[right] == nums[right - 1]:
                        right -= 1
                    left += 1
                    right -= 1

                elif nums[i] + nums[left] + nums[right] < 0:
                    left += 1

                elif nums[i] + nums[left] + nums[right] > 0:
                    right -= 1
        return res


# @lc code=end

s = Solution()

print(s.threeSum([-1, 0, 1, 2, -1, -4]))
