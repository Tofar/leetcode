# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        num1 = 0
        num2 = 0
        temp = 1
        while l1 is not None:
            num1 = num1 + l1.val * temp
            l1 = l1.next
            temp = temp * 10
        temp = 1
        while l2 is not None:
            num2 = num2 + l2.val * temp
            l2 = l2.next
            temp = temp * 10
        num = num1 + num2
        remainder = num % 10
        twonum = ListNode(remainder)
        head = twonum

        num = num/10
        remainder = num % 10
        while num != 0:
            twonum.next = ListNode(remainder)
            twonum = twonum.next
            num = num/10
            remainder = num % 10
        return head
