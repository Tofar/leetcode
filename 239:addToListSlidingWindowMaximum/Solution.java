package addToListSlidingWindowMaximum;

import org.w3c.dom.ranges.Range;

/**
  Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

 For example,
 Given nums = [1,3,-1,-3,5,3,6,7], and k = 3.

 Window position                Max
 ---------------               -----
 [1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
 Therefore, return the max sliding window as [3,3,5,5,6,7].

 Note:
 You may assume k is always valid, ie: 1 ≤ k ≤ input array's size for non-empty array.

 Follow up:
 Could you solve it in linear time?

 */

public class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {
        if (k > nums.length) {
            return new int[]{};
        } else {
            if (nums.length == 0) {
                return nums;
            }
        }

        int[] result = new int[nums.length - k + 1];
        for (int i = 0; i < nums.length - k + 1; i++) {
            int[] input = new int[k];
            for (int j = 0; j < k; j++) {
                input[j] = nums[i + j];
            }
            result[i] = getMax(input);
        }

        return result;
    }

    public int getMax(int[] nums) {
        int max = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (max < nums[i]) {
                max = nums[i];
            }
        }

        return max;
    }
}
