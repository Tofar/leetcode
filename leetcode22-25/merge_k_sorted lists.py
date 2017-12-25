# Definition for singly-linked list.
class ListNode(object):
     def __init__(self, x):
         self.val = x
         self.next = None

class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """

        def merge(listnode1, listnode2):
            header = ListNode(0)
            temp = header
            while listnode1 and listnode2:
                if listnode1.val < listnode2.val:
                    temp.next = listnode1
                    listnode1 = listnode1.next
                    temp = temp.next
                else:
                    temp.next = listnode2
                    listnode2 = listnode2.next
                    temp = temp.next

            if listnode1:
                temp.next = listnode1
            elif listnode2:
                temp.next = listnode2
            return header.next

        def back(alist):
            if len(alist) == 1:
                return alist[0]
            else:
                mid = int(len(alist)/2)
                return merge(back(alist[:mid]), back(alist[mid:len(alist)]))
                             
        length = len(lists)
        if length == 0:
            return None
        return back(lists)
        
