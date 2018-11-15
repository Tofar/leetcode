package main

import (
	"fmt"
)

func candy(ratings []int) int {
	nums := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		nums[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] < ratings[i-1] {
			if nums[i-1] <= nums[i] {
				nums[i-1] += 1
			}
		} else if ratings[i] > ratings[i-1] {
			nums[i] = nums[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && nums[i] <= nums[i+1] {
			nums[i] = nums[i+1] + 1
		}
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	ratings := []int{1, 2, 87, 87, 87, 2, 1}
	fmt.Println(candy(ratings))
}

