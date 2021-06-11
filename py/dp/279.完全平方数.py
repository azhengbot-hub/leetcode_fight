#
# @lc app=leetcode.cn id=279 lang=python3
#
# [279] 完全平方数
#

# @lc code=start
class Solution:
    def numSquares(self, n: int) -> int:
        dp = [n + 1] * (n + 1)

        dp[0] = 0
        for i in range(1, n + 1):
            for j in range(i * i, n + 1):
                dp[j] = min(dp[j - i * i] + 1, dp[j])

        return dp[-1]


# @lc code=end

s = Solution()

print(s.numSquares(12))
