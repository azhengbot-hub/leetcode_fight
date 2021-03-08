#
# @lc app=leetcode.cn id=39 lang=python3
#
# [39] 组合总和
#

# @lc code=start
from typing import List


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        res = []

        def back_track(path=[], new_target=target, begin=0):
            if new_target == 0:
                res.append(path[:])
            if new_target < min(candidates):
                return

            for index in range(begin, len(candidates)):
                path.append(candidates[index])
                other_sum = new_target - candidates[index]

                back_track(path, other_sum, begin=index)

                path.pop()

        back_track([], target, 0)

        return res


# @lc code=end


s = Solution()

print(s.combinationSum([2, 3, 6, 7], 7))
