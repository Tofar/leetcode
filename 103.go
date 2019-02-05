package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type cntrl struct {
	Dir       bool // 方向：true -> 从左往右, false -> 从右往左
	CurNodes  []*TreeNode
	NextNodes []*TreeNode
	Zig       [][]int
	CurArray  []int
}

func newCntrl(curNodes []*TreeNode) *cntrl {
	return &cntrl{
		Dir:       true,
		CurNodes:  curNodes,
		NextNodes: []*TreeNode{},
		Zig:       [][]int{},
		CurArray:  []int{},
	}
}

func (this *cntrl) Pull() *TreeNode {
	if l := len(this.CurNodes); l != 0 {
		var node *TreeNode
		if this.Dir {
			node = this.CurNodes[0]
			this.CurNodes = this.CurNodes[1:]
		} else {
			node = this.CurNodes[l-1]
			this.CurNodes = this.CurNodes[:l-1]
		}
		this.CurArray = append(this.CurArray, node.Val)
		return node
	} else {
		this.Zig = append(this.Zig, this.CurArray)
		this.CurArray = []int{}

		if len(this.NextNodes) == 0 {
			return nil
		}
		this.CurNodes = this.NextNodes
		this.NextNodes = []*TreeNode{}
		this.Dir = !this.Dir
		return this.Pull()
	}
}

func (this *cntrl) Put(node *TreeNode) {
	if this.Dir {
		// 方向：true -> 从左往右遍历, 从左往右插入
		if node.Left != nil {
			this.NextNodes = append(this.NextNodes, node.Left)
		}
		if node.Right != nil {
			this.NextNodes = append(this.NextNodes, node.Right)
		}
	} else {
		if node.Right != nil {
			this.NextNodes = append([]*TreeNode{node.Right}, this.NextNodes...)
		}
		if node.Left != nil {
			this.NextNodes = append([]*TreeNode{node.Left}, this.NextNodes...)
		}
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ct := newCntrl([]*TreeNode{root})
	curNode := ct.Pull()
	for curNode != nil {
		ct.Put(curNode)
		curNode = ct.Pull()
	}
	return ct.Zig
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(zigzagLevelOrder(root))
}

