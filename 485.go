func findMaxConsecutiveOnes(nums []int) int {
	num := 0
	tempNum := 0
	for _, v := range nums {
		if v == 1 {
			tempNum++
		} else {
			if tempNum > num {
				num = tempNum
			}
			tempNum = 0
		}
	}
	if tempNum > num {
		num = tempNum
	}
	return num
}
