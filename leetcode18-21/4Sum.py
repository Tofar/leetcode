"""
Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note: The solution set must not contain duplicate quadruplets.

For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
"""


class Solution:

    def fourSum(self, num, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        # two_sum_dict 存储two_sum(两个数的和)
        # result 为set，去重
        num_len, result, two_sum_dict = len(num), set(), {}
        if num_len < 4:
            return []
        # 排序
        num.sort()
        for p in range(num_len):
            for q in range(p + 1, num_len):
                if num[p] + num[q] not in two_sum_dict:
                    two_sum_dict[num[p] + num[q]] = [(p, q)]
                else:
                    two_sum_dict[num[p] + num[q]].append((p, q))

        for i in range(num_len):
            for j in range(i + 1, num_len - 2):
                T = target - num[i] - num[j]
                if T in two_sum_dict:
                    for k in two_sum_dict[T]:
                        if k[0] > j:
                            result.add((num[i], num[j], num[k[0]], num[k[1]]))
        return list(result)
