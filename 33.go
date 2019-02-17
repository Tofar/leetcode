func search(nums []int, target int) int {
	nL := len(nums)
	if nL == 0 {
		return -1
	} else if nL == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	p := findPivot(nums)
	var l, r int
	if p == -1 {
		l = 0
		r = nL - 1
	} else {
		if nums[0] > target {
			l = p + 1
			r = nL - 1
		} else if nums[0] == target {
			return 0
		} else {
			l = 1
			r = p
		}
	}

	for mid := (l + r) / 2; l <= r; {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = (l + r) / 2
	}
	return -1
}

// [4,5,6,7,0,1,2] -> index:3
func findPivot(nums []int) int {
	if len(nums) <= 1 {
		return -1
	}

	l, r := 0, len(nums)-1
	i := r / 2
	for l <= r {
		if i >= len(nums)-1 {
			break
		}
		if nums[i] <= nums[i+1] {
			if nums[i] < nums[0] {
				// turn left
				r = i - 1
			} else {
				// turn right
				l = i + 1
			}
			i = (l + r) / 2
		} else {
			return i
		}

	}
	return -1
}
