#
# @lc app=leetcode.cn id=4 lang=python3
#
# [4] 寻找两个正序数组的中位数
#

# @lc code=start
class Solution:
    def findMedianSortedArrays(self, nums1: list, nums2: list) -> float:
        nums = sorted(nums1 + nums2)
        n = len(nums)
        half = n / 2
        if n % 2 == 0:
            res = (nums[int(half) - 1] + nums[int(half)]) / 2

        else:
            res = nums[int(half)]

        return res

# @lc code=end


s = Solution()

print(s.findMedianSortedArrays([1, 2], [3, 4]))
