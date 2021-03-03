#
# @lc app=leetcode.cn id=17 lang=python3
#
# [17] 电话号码的字母组合
#

# @lc code=start
from typing import List


class Solution:
    res_map = {
        1: [],
        2: ["a", "b", "c"],
        3: ["d", "e", "f"],
        4: ["g", "h", "i"],
        5: ["j", "k", "l"],
        6: ["m", "n", "o"],
        7: ["p", "q", "r", "s"],
        8: ["t", "u", "v"],
        9: ["w", "x", "y", "z"],
    }

    def letterCombinations(self, digits: str) -> List[str]:

        if not digits:
            return []

        res: List[str] = []
        combination = []

        def backTrack(index):
            if index == len(digits):
                res.append("".join(combination))
            else:
                digit = digits[index]
                for i in self.res_map[int(digit)]:
                    combination.append(i)
                    backTrack(index + 1)
                    combination.pop()

        backTrack(0)
        return res


# @lc code=end

s = Solution()

print(s.letterCombinations("23"))