from typing import Optional


class ListNode:
    def __init__(self, val: int):
        self.val: int = val
        self.next: Optional[ListNode] = None

    def __eq__(self, o: object) -> bool:
        return isinstance(o, ListNode) and self.val == o.val and self.next == o.next

    def __str__(self):
        result = str(self.val)
        if self.next is not None:
            result += f'->{self.next.__str__()}'
        return result

    def __repr__(self):
        next_repr = None if self.next is None else self.next.__repr__()
        return f'ListNode({self.val}, {next_repr}'


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
