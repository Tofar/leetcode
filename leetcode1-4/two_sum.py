class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        new_nums = sorted(nums)
        length = len(nums)
        target2 = target/2
        i = 0
        a = b = 0
        while new_nums[i] <= target2:
            j = length - 1
            check = 0
            while new_nums[j] >= target2:
                if (new_nums[i] + new_nums[j]) == target:
                    a = new_nums[i]
                    b = new_nums[j]
                    check = 1
                    break
                j = j - 1
            if check:
                break
            i = i + 1
        i = 0
        if a == b:
            check2 = 0
            temp1 = 0
            while i < length:
                if check2 and (nums[i] == a):
                    b = i
                    break
                if (check2 == 0) and (nums[i] == a):
                    check2 = 1
                    temp1 = i
                i = i + 1
            a = temp1
        else:
            check3 = 0
            while i < length:
                if nums[i] == a:
                    a = i
                    check3 = check3 + 1
                if nums[i] == b:
                    b = i
                    check3 = check3 + 1
                if check3 == 2:
                    break
                i = i + 1
        return [a, b]
