#
# @lc app=leetcode.cn id=424 lang=python3
#
# [424] 替换后的最长重复字符
#

# @lc code=start
from collections import Counter


class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        n = len(s)
        left, right = 0, 0
        counter = Counter()
        res = 0

        while right < n:
            counter[s[right]] += 1
            while right - left + 1 - counter.most_common(1)[0][1] > k:
                counter[s[left]] -= 1
                left += 1

            res = max(res, right-left+1)
            right += 1
        return res

        # left = 0
        # right = 1
        # sum_res = 0

        # if len(s) == 1:
        #     return 1

        # while right <= len(s):
        #     check_s = s[left: right]
        #     count_check_s = sorted(list(Counter(check_s).values()))[-1]
        #     len_check_s = len(check_s)
        #     if len_check_s - count_check_s <= k:
        #         right += 1
        #         if len_check_s > sum_res:
        #             sum_res = len_check_s
        #     else:
        #         left += 1

        # return sum_res

# @lc code=end


s = Solution()

print(s.characterReplacement("AABABBA", 1))
