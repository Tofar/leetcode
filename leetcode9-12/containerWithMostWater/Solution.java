package containerWithMostWater;

/**
 * Given n non-negative integers a1, a2, ..., an, where each represents a point
 * at coordinate (i, ai). n vertical lines are drawn such that the two endpoints
 * of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis
 * forms a container, such that the container contains the most water.

 * Note: You may not slant the container and n is at least 2.
 */

// 暴力解超时
public class Solution {
    public int maxArea(int[] height) {
        int left = 0, right = height.length - 1, maxArea = 0;
        while(left < right){
            // 每次更新最大面积（盛水量）
            maxArea = Math.max(maxArea, Math.min(height[left], height[right]) * (right - left));
            // 如果左边较低，则将左边向中间移一点
            if(height[left] < height[right]){
                left++;
                // 如果右边较低，则将右边向中间移一点
            } else {
                right--;
            }
        }
        return maxArea;
    }
}
