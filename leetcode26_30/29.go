package leetcode26_30

import (
	"fmt"
)

/*
Divide two integers without using multiplication, division and mod operator.

If it is overflow, return MAX_INT.
*/

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 2147483647
	}
	isNeg := (divisor^divisor)>>31 == 1
	result := 0
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}
	if divisor == -2147483648 {
		return 0
	}
	if divisor < 0 {
		divisor = ^divisor + 1
	}
	if dividend < 0 {
		dividend = ^dividend + 1
	}
	loc := 0
	for divisor <= dividend>>1 {
		divisor = divisor << 1
		loc++
	}
	for loc >= 0 {
		if dividend >= divisor {
			result += 1 << uint(loc)
			dividend -= divisor
		}
		divisor = divisor >> 1
		loc--
	}
	fmt.Println("result: ", result)
	if isNeg {
		result = -result
	}
	fmt.Println("result: ", result)

	return result
}
