package _3Sum;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 *
 Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

 Note: The solution set must not contain duplicate triplets.

 For example, given array S = [-1, 0, 1, 2, -1, -4],

 A solution set is:
 [
 [-1, 0, 1],
 [-1, -1, 2]
 ]
 */
public class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        if (nums.length < 3) {
            return new ArrayList<>();
        } else {
            // 首先对数组进行排序
            Arrays.sort(nums);

            List<List<Integer>> result = new ArrayList<>();

            for (int i = 0; i < nums.length - 2; i++) {
                // 去重
                if (i > 0 && nums[i] == nums[i - 1]) {
                    continue;
                }

                int start = i + 1;
                int end = nums.length - 1;

                while (start < end) {
                    int sum = nums[i] + nums[start] + nums[end];
                    if (sum == 0) {
                        ArrayList<Integer> temp = new ArrayList<>();
                        temp.add(nums[i]);
                        temp.add(nums[start]);
                        temp.add(nums[end]);
                        result.add(temp);
                        start++;
                        end--;

                        // 去重
                        while (start < end && nums[start] == nums[start - 1]) {
                            start++;
                        }
                        while (start < end && nums[end] == nums[end + 1]) {
                            end--;
                        }

                    } else if (sum < 0) {
                        start++;
                    } else {
                        end--;
                    }

                }
            }
            return result;
        }
    }
}
