#
# @lc app=leetcode.cn id=1208 lang=python3
#
# [1208] 尽可能使字符串相等
#

# @lc code=start


class Solution:
    def equalSubstring(self, s: str, t: str, maxCost: int) -> int:
        n = len(s)
        diff_list = []

        start = 0
        end = 0
        res = 0

        for i in range(n):
            diff = abs(ord(s[i]) - ord(t[i]))
            diff_list.append(diff)
        while end <= n:
            diff_slice = diff_list[start: end]

            sum_diff_slice = sum(diff_slice)

            if sum_diff_slice <= maxCost:
                end += 1
                res = max(res, len(diff_slice))
            else:
                start += 1
                end += 1
                res = max(res, len(diff_slice)-1)

        return res


# @lc code=end
s = Solution()
print(s.equalSubstring("krpgjbjjznpzdfy", "nxargkbydxmsgby", 14))
