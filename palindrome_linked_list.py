# Definition for singly-linked list.
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        nodes_array = []
        current_node = head
        while current_node != None:
            nodes_array.append(current_node.val)
            current_node = current_node.next

        for i in range((len(nodes_array) // 2) + 1):
            if nodes_array[i] != nodes_array[-1-i]:
                return False

        return True


print(Solution().isPalindrome(None))
