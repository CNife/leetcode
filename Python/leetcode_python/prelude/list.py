from typing import Optional


class ListNode:
    def __init__(self, val: int):
        self.val: int = val
        self.next: Optional[ListNode] = None

    def __eq__(self, o: object) -> bool:
        return isinstance(o, ListNode) and self.val == o.val and self.next == o.next

    def __str__(self) -> str:
        return f"{self.val}->{self.next}"

    def __repr__(self):
        return f"ListNode(val={self.val},next={self.next})"


def new_list(*nums: int) -> Optional[ListNode]:
    if nums:
        head = ListNode(nums[0])
        node = head
        for num in nums[1:]:
            node.next = ListNode(num)
            node = node.next
        return head
    else:
        return None
