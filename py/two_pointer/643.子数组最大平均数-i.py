#
# @lc app=leetcode.cn id=643 lang=python3
#
# [643] 子数组最大平均数 I
#

# @lc code=start
class Solution:
    def findMaxAverage(self, nums: list, k: int) -> float:
        res = total = sum(nums[:k])

        n = len(nums)

        for i in range(n - k):
            total = total + nums[i + k] - nums[i]
            res = max(total, res)
        return res / k
        # while right <= len(nums):
        #     slice_res = sum(nums[left:right])

        #     if slice_res > res:
        #         res = slice_res

        #     left += 1
        #     right = left + k
        # return res / k


# @lc code=end

s = Solution()
print(s.findMaxAverage([1, 12, -5, -6, 50, 3], 4))
