package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := math.MinInt64
	helpFunc(root, &res)
	return res
}

func helpFunc(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	l := helpFunc(root.Left, res)
	r := helpFunc(root.Right, res)
	*res = max(*res, root.Val+max(0, l)+max(0, r))
	return max(0, max(l, r)+root.Val)
}

func max(array ...int) int {
	if len(array) == 0 {
		return 0
	}
	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func main() {

}

