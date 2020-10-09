from typing import List

from leetcode import ListNode, test, new_list


def has_cycle(head: ListNode) -> bool:
    if not head:
        return False

    slow, fast = head, head.next
    while slow and fast:
        if slow is fast:
            return True
        slow, fast = slow.next, fast.next
        if fast:
            fast = fast.next
    return False


def _new_cycle_list(nums: List[int], cycle_end: int) -> ListNode:
    src_list = new_list(*nums)
    start_node, end_node = src_list[len(nums) - 1], src_list[cycle_end]
    start_node.next = end_node
    return src_list


test(
    has_cycle,
    [
        (_new_cycle_list([3, 2, 0, -4], 1), True),
        (_new_cycle_list([1, 2], 0), True),
        (new_list(1), False),
        (new_list(1, 2), False),
    ],
)
