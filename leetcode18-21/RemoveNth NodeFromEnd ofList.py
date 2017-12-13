# -*- coding : utf-8 -*-
# Definition for singly-linked list.


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        num = self.remove(head, head.next, n)
        if num < n:
            head = head.next
        return head

    def remove(self, pre, head, n):
        if head.next is None:
            return 1
        else:
            num = self.remove(head, head.next, n) + 1
            if num == n:
                pre.next = head.next
            return num
