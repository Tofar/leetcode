package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Cur  *TreeNode
	Next *Node
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	cur := root
	head := &Node{
		Cur: cur,
	}

	curNode := head

	if cur.Left != nil && cur.Right == nil {
		curNode.Next = &Node{
			Cur: cur.Left,
		}
	}

	if cur.Right != nil && cur.Left == nil {
		curNode.Next = &Node{
			Cur: cur.Right,
		}
	}

	if cur.Left != nil && cur.Right != nil {
		curNode.Next = &Node{
			Cur: cur.Left,
			Next: &Node{
				Cur: cur.Right,
			},
		}
	}
	curNode = head.Next

	for curNode != nil {
		cur = curNode.Cur
		var temp *Node
		if cur.Left != nil && cur.Right == nil {
			temp = &Node{
				Cur: cur.Left,
			}
		}

		if cur.Right != nil && cur.Left == nil {
			temp = &Node{
				Cur: cur.Right,
			}
		}

		if cur.Left != nil && cur.Right != nil {
			temp = &Node{
				Cur: cur.Left,
				Next: &Node{
					Cur: cur.Right,
				},
			}
		}

		if temp != nil {
			if curNode.Next != nil {
				if temp.Next != nil {
					temp.Next.Next = curNode.Next
				} else {
					temp.Next = curNode.Next
				}
			}
			curNode.Next = temp
		}

		curNode = curNode.Next
	}

	curNode = head
	for curNode != nil {
		result = append(result, curNode.Cur.Val)
		curNode = curNode.Next
	}

	return result
}

func main() {
	n := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(preorderTraversal(n))
}

