package _3SumClosest;

import java.util.Arrays;

/**
 * Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

 For example, given array S = {-1 2 1 -4}, and target = 1.

 The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 */

public class Solution {
    public int threeSumClosest(int[] nums, int target) {
        if (nums.length < 3) {
            return target;
        } else {
            // 首先对数组进行排序
            Arrays.sort(nums);
            int min_target = nums[0] + nums[1] + nums[2];

            for (int i = 0; i < nums.length - 2; i++) {
                // 去重
                if (i > 0 && nums[i] == nums[i - 1]) {
                    continue;
                }

                int start = i + 1;
                int end = nums.length - 1;

                while (start < end) {
                    int sum = nums[i] + nums[start] + nums[end];

                    if (Math.abs(target - min_target) > Math.abs(target - sum)) {
                        min_target = sum;
                    }
                    if (sum == target) {
                        return target;
                    } else if (sum < target) {
                        start++;
                    } else {
                        end--;
                    }

                }
            }
            return min_target;
        }
    }
}