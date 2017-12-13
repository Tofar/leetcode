# -*- coding : utf-8 -*-
# Definition for singly-linked list.


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        if l1 is None:
            return l2
        elif l2 is None:
            return None

        linked_list = ListNode(0)
        head = linked_list
        while not (l1 is None and l2 is None):
            if l1.val < l2.val:
                linked_list.next = ListNode(l1.val)
                linked_list = linked_list.next
                l1 = l1.next
            else:
                linked_list.next = ListNode(l2.val)
                linked_list = linked_list.next
                l2 = l2.next
        return head.next

