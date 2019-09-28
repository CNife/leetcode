from collections import namedtuple
from heapq import heappush, heapreplace, heappop
from typing import List, Optional
from unittest import TestCase

from list import ListNode, new_list


def merge_k_lists(lists: List[Optional[ListNode]]) -> Optional[ListNode]:
    tops: List[CmpListNode] = []
    for node in lists:
        if node is not None:
            heappush(tops, CmpListNode(node))

    head = ListNode(-1)
    tail = head
    while tops:
        tail.next = tops[0].node
        tail = tail.next
        if tail.next is not None:
            heapreplace(tops, CmpListNode(tail.next))
        else:
            heappop(tops)

    return head.next


CmpListNode = namedtuple('CmpListNode', ['node'])


def cmp_list_node_lt(lhs, rhs):
    return lhs.node.val < rhs.node.val


CmpListNode.__slots__ = ()
CmpListNode.__lt__ = cmp_list_node_lt


class TestMergeKLists(TestCase):
    def test(self):
        input_lists = [new_list([1, 4, 5]), new_list([1, 3, 4]), new_list([2, 6])]
        expected = new_list([1, 1, 2, 3, 4, 4, 5, 6])
        output = merge_k_lists(input_lists)
        self.assertEqual(output, expected)
