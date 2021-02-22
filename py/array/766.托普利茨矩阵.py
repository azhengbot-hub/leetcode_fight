from typing import List

#
# @lc app=leetcode.cn id=766 lang=python3
#
# [766] 托普利茨矩阵
#

# @lc code=start


class Solution:
    def isToeplitzMatrix(self, matrix: List[List[int]]) -> bool:
        if len(matrix) == 1 or len(matrix[0]) == 1:
            return True

        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                if i != 0 and j != 0:
                    if matrix[i][j] != matrix[i - 1][j - 1]:
                        return False

        return True


# @lc code=end


s = Solution()
# print(s.isToeplitzMatrix([[1,2,3,4],[5,1,2,3],[9,5,1,2]]))
# print(s.isToeplitzMatrix([[1,2],[2,2]]))
assert s.isToeplitzMatrix([[1, 2, 3, 4], [5, 2, 1, 3]]) == False
