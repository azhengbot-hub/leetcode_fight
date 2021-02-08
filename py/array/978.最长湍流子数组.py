#
# @lc app=leetcode.cn id=978 lang=python3
#
# [978] 最长湍流子数组
#

# @lc code=start
class Solution:
    def maxTurbulenceSize(self, arr: list) -> int:
        n = len(arr)
        res = 1
        left, right = 0, 0

        while right < n-1:
            if right == left:
                if arr[left] == arr[left+1]:
                    left += 1
                right += 1

            else:
                if arr[right-1] < arr[right] and arr[right] > arr[right+1]:
                    right += 1
                elif arr[right-1] > arr[right] and arr[right] < arr[right+1]:
                    right += 1
                else:
                    left = right
            res = max(res, right-left+1)

        return res
        # n = len(arr)
        # res_list = []

        # if len(set(arr)) == 1:
        #     return 1

        # for i in range(1, n):
        #     res = arr[i] - arr[i-1]
        #     if res > 0:
        #         res_list.append(1)
        #     elif res < 0:
        #         res_list.append(0)
        #     else:
        #         res_list.append(-1)

        # left = 0
        # right = 1
        # res = 0

        # while right < n-1:
        #     if res_list[left] == -1:

        #         left += 1
        #         right = left + 1
        #         continue
        #     if res_list[right] == res_list[right-1]:

        #         res = max(res, right-left-1)
        #         left += 1
        #         right = left+1
        #     elif res_list[right] == -1:
        #         res = max(res, right-left-1)
        #         left += 1
        #         right = left+1
        #     else:
        #         right += 1

        # res = max(res, right-left-1)

        # return res+2

        # @lc code=end


s = Solution()

print(s.maxTurbulenceSize([9, 4, 2, 10, 7, 8, 8, 1, 9]))
# print(s.maxTurbulenceSize([9, 9]))
# print(s.maxTurbulenceSize([0, 1, 1, 0, 1, 0, 1, 1, 0, 0]))
# print(s.maxTurbulenceSize([0, 8, 45, 88, 48, 68, 28, 55, 17, 24]))

# [9, 4, 2, 10, 7, 8, 8, 1, 9]
# [  0, 0, 1, 0, 1, -1]
