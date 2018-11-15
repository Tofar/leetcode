package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	mapNum := map[int]int{}
	for _, num := range nums {
		if _, ok := mapNum[num]; !ok {
			if leftLen, leftOk := mapNum[num-1]; leftOk {
				if rightLen, rightOk := mapNum[num+1]; rightOk {
					l := leftLen + 1 + rightLen
					mapNum[num-leftLen] = l
					mapNum[num+rightLen] = l
					mapNum[num] = l
				} else {
					l := leftLen + 1
					mapNum[num] = l
					mapNum[num-leftLen] = l
				}
				continue
			}
			if rightLen, rightOk := mapNum[num+1]; rightOk {
				if leftLen, leftOk := mapNum[num-1]; leftOk {
					l := leftLen + 1 + rightLen
					mapNum[num-leftLen] = l
					mapNum[num+rightLen] = l
					mapNum[num] = l
				} else {
					mapNum[num+rightLen] = rightLen + 1
					mapNum[num] = rightLen + 1
				}

				continue
			}
			mapNum[num] = 1
		}
	}

	max := 0
	for _, len := range mapNum {
		if len > max {
			max = len
		}
	}
	return max
}

func main() {
	test := []int{4, 0, -4, -2, 2, 5, 2, 0, -8, -8, -8, -8, -1, 7, 4, 5, 5, -4, 6, 6, -3}
	fmt.Println(longestConsecutive(test))
}

