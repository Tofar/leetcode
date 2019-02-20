func PredictTheWinner(nums []int) bool {
	return winNum(nums) >= 0
}

func winNum(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		if nums[0] > nums[1] {
			return nums[0] - nums[1]
		}
		return nums[1] - nums[0]
	default:
		num1 := nums[0] - winNum(nums[1:])
		num2 := nums[len(nums)-1] - winNum(nums[:(len(nums)-1)])
		if num1 > num2 {
			return num1
		}
		return num2
	}
}
