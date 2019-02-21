/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    // quick sort
    if head == nil || head.Next == nil {
        return head
    }
    if head.Next.Next == nil {
        if head.Next.Val < head.Val {
            head.Val, head.Next.Val = head.Next.Val, head.Val
        }
        return head
    }
    
    left := head
    for left.Next != nil && left.Next.Val < head.Val {
        left = left.Next
    }
    
    for right := left.Next; right != nil; right = right.Next {
        if right.Val < head.Val {
            left = left.Next
            left.Val, right.Val = right.Val, left.Val
        }
    }
    head.Val, left.Val = left.Val, head.Val
    
    temp := left.Next
    left.Next = nil
    sortList(head)
    sortList(temp)
    left.Next = temp
    return head
}

