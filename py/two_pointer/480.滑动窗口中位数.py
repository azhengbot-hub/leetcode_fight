#
# @lc app=leetcode.cn id=480 lang=python3
#
# [480] 滑动窗口中位数
#

# @lc code=start
class Solution:
    def medianSlidingWindow(self, nums: list, k: int) -> list:

        left = 0
        res_list = []

        while left+k <= len(nums):
            aim_nums = sorted(nums[left:left+k])

            if k % 2 == 0:
                res = (aim_nums[int(k/2)] + aim_nums[int(k/2)-1])/2
            else:
                res = aim_nums[int((k-1)/2)]

            res_list.append(res)

            left += 1

        return res_list
# @lc code=end


s = Solution()

print(s.medianSlidingWindow([1, 3, -1, -3, 5, 3, 6, 7], 3))
