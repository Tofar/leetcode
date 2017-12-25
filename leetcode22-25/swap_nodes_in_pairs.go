package leetcode22_25

/*
Given a linked list, swap every two adjacent nodes and return its head.

For example,
Given 1->2->3->4, you should return the list as 2->1->4->3.

Your algorithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.


*/
import (
	"fmt"
)

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head.Next != nil {
		var pre = head
		var back = pre.Next
        var result
        if back != nil {
            result = back
        } else {
            result = pre
        }
		for pre != nil && back != nil {
			// swap
			temp := back.Next
			back.Next = pre
			pre.Next = temp

			pre = temp
			if pre != nil {
				back = temp.Next
			}
		}
		return result
	} else {
		return head
	}

}

}
