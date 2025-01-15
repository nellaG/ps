# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeTwoLists(
        self, list1: Optional[ListNode], list2: Optional[ListNode]
    ) -> Optional[ListNode]:
        answer = ListNode()
        head = answer
        while list1 and list2:
            if list1.val <= list2.val:
                answer.next = list1
                answer = list1
                list1 = list1.next
            else:
                answer.next = list2
                answer = list2
                list2 = list2.next

        remain = list1 or list2

        answer.next = remain
        return head.next
