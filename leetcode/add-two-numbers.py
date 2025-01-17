# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


from typing import Optional


class Solution:
    def addTwoNumbers(
        self, l1: Optional[ListNode], l2: Optional[ListNode]
    ) -> Optional[ListNode]:
        carry = 0
        carry2 = 0
        answer = ListNode()
        head = answer
        while l1 or l2 or carry != 0:
            l1val = l1.val if l1 else 0
            l2val = l2.val if l2 else 0
            digit = l1val + l2val
            if digit + carry >= 10:
                digit -= 10
                carry2 = 1
            answer.next = ListNode(val=digit + carry)
            carry = carry2
            carry2 = 0
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
            answer = answer.next

        return head.next
