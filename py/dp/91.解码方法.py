class Solution:
    def numDecodings(self, s: str) -> int:

        if s.startswith("0"):
            return 0

        dp = [1 for i in range(len(s) + 1)]

        for i in range(2, len(s) + 1):
            if s[i - 1] == "0" and (not s[i - 2] in "12"):
                return 0
            elif s[i - 1] == "0" and s[i - 2] in "12":
                dp[i] = dp[i - 2]
            elif "11" <= s[i - 2 : i] <= "26":
                dp[i] = dp[i - 2] + dp[i - 1]
            else:
                dp[i] = dp[i - 1]

        return dp[len(s)]
