#
# @lc app=leetcode.cn id=665 lang=python3
#
# [665] 非递减数列
#

# @lc code=start
class Solution:
    def checkPossibility(self, nums: list) -> bool:

        res = 0
        for i in range(1, len(nums)):

            if nums[i] < nums[i-1]:
                res += 1
                if i == 1 or nums[i] >= nums[i-2]:
                    nums[i-1] = nums[i]

                else:
                    nums[i] = nums[i-1]

        return res <= 1

# @lc code=end


s = Solution()
print(s.checkPossibility([4, 2, 3]))
