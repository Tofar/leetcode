/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := l1
    if l2 != nil && (l1 == nil || l1.Val > l2.Val) {
        head = l2
        l2 = l2.Next
    } else if l1 != nil {
        l1 = l1.Next
    }
    
    for cur := head; l1 != nil || l2 != nil; {
        if l2 != nil && (l1 == nil || l1.Val > l2.Val) {
            cur.Next = l2
            l2 = l2.Next
        } else {
            cur.Next = l1
            l1 = l1.Next
        }
        cur = cur.Next
    }
    
    return head
}
