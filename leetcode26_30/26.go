package leetcode26_30

/**
 * Given a sorted array, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
It doesn't matter what you leave beyond the new length.

*/

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	for i, j := 0, 1; j < len(nums); {
		if nums[i] == nums[j] {
			nums = append(nums[:j], nums[j+1:]...)
		} else {
			j++
			i++
		}
	}
	return len(nums)
}

// 官方解答
func removeDuplicates2(nums []int) int {
	if nums == nil {
		return 0
	}
	rel := 1
	for i, j := 0, 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
			rel++
		}
	}
	return rel
}
