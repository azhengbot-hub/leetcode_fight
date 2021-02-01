#
# @lc app=leetcode.cn id=888 lang=python3
#
# [888] 公平的糖果交换
#

# @lc code=start
class Solution:
    def fairCandySwap(self, A: list, B: list) -> list:
        sum_a, sum_b = sum(A), sum(B)
        half_sum = (sum_a - sum_b) / 2

        print(half_sum)

        for ai in A:
            bi = ai - half_sum
            if bi in B:
                return [ai, int(bi)]

        # sum_A = sum(A)
        # sum_B = sum(B)
        # for ai in A:
        #     for bi in B:
        #         if sum_A - ai + bi == sum_B - bi + ai:
        #             return [ai, bi]

                # i = 0
                # j = 0

                # while i < len(A):
                #     while j < len(B):
                #         # print(i, j)
                #         A[i], B[j] = B[j], A[i]
                #         sum_a = sum(A)
                #         sum_b = sum(B)
                #         # print(A)
                #         # print(B)
                #         # print(sum_a)
                #         # print(sum_b)
                #         if sum_a == sum_b:
                #             return [B[j], A[i]]
                #         else:
                #             A[i], B[j] = B[j], A[i]
                #             j += 1
                #     i += 1
                #     j = 0

                # @lc code=end


s = Solution()

res = s.fairCandySwap([1, 2], [2, 3])

print(res)
