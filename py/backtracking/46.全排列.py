#
# @lc app=leetcode.cn id=46 lang=python3
#
# [46] 全排列
#

# @lc code=start

from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        length = len(nums)
        used = [False for _ in range(length)]

        def dfs(nums, path, depth=0, used=used):
            # path = []
            if depth == length:
                res.append(path[:])
                return

            for i in range(length):
                if not used[i]:
                    used[i] = True
                    path.append(nums[i])

                    dfs(nums, path, depth + 1, used)

                    used[i] = False
                    path.pop()

        dfs(nums, path=[])

        return res

        # def backtrack(first=0):
        #     if first == len(nums):
        #         res.append(nums[:])
        #         return

        #     for i in range(first, len(nums)):
        #         print(i, "=====", first)
        #         nums[first], nums[i] = nums[i], nums[first]
        #         print(nums)
        #         backtrack(first + 1)

        #         nums[i], nums[first] = nums[first], nums[i]

        # backtrack()
        # return res

        # def backtrack(nums, tmp):
        #     if not nums:
        #         res.append(tmp)
        #         return

        #     for i in range(len(nums)):
        #         backtrack(nums[:i] + nums[i + 1 :], tmp + [nums[i]])

        # backtrack(nums, [])

        # return res


# @lc code=end

s = Solution()

print(s.permute([1, 2, 3]))
