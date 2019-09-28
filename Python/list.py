from typing import Optional, List


class ListNode:
    def __init__(self, val: int):
        self.val: int = val
        self.next: Optional[ListNode] = None

    def __eq__(self, o: object) -> bool:
        if object.__eq__(self, o):
            return True
        elif not isinstance(o, ListNode):
            return False
        else:
            return self.val == o.val and self.next == o.next


def new_list(nums: List[int]) -> Optional[ListNode]:
    if len(nums) == 0:
        return None
    else:
        head = ListNode(nums[0])
        head.next = new_list(nums[1:])
        return head
