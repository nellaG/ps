# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        idx = 0
        llist = head
        while head:
            idx += 1
            head = head.next
        if idx == n:
            return llist.next
        x = idx - n
        head = llist
        while x != 1:
            x -= 1
            head = head.next if head.next else None
        head.next = head.next.next

        return llist
