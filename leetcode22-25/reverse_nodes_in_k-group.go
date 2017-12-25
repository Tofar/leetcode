package leetcode22_25

/*
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

You may not alter the values in the nodes, only nodes itself may be changed.

Only constant memory is allowed.

For example,
Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	root := &ListNode{}
	root.Next = head
	tail := root
	currentNode := head
	for {
		// 判断是否提前中断（剩余没有 k 个节点
		ifBreak := false
		var reversingNodes []*ListNode
		// 利用堆栈让链表反向
		for i := 0; i < k; i++ {
			if currentNode == nil {
				ifBreak = true
				break
			}
			reversingNodes = append(reversingNodes, currentNode)
			currentNode = currentNode.Next
		}

		if ifBreak {
			if len(reversingNodes) > 0 {
				// 提前中断
				tail.Next = reversingNodes[0]
			} else {
				tail.Next = nil
			}
			break
		}

		// 反向连接
		for i := k - 1; i >= 1; i-- {
			reversingNodes[i].Next = reversingNodes[i-1]
		}

		tail.Next = reversingNodes[len(reversingNodes)-1]
		tail = reversingNodes[0]
	}
	return root.Next
}
