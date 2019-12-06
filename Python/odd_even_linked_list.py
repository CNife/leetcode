from typing import Optional
from unittest import TestCase

from list import ListNode, new_list


def odd_even_list(head: Optional[ListNode]) -> Optional[ListNode]:
    if head is None or head.next is None or head.next.next is None:
        return head

    flag = True
    even_head = head.next
    odds = head
    evens = even_head
    ptr = even_head.next
    while ptr is not None:
        if flag:
            odds.next = ptr
            odds = ptr
        else:
            evens.next = ptr
            evens = ptr
        ptr = ptr.next
        flag = not flag

    odds.next = even_head
    evens.next = None
    return head


class Test(TestCase):
    def test_odd_even_list(self):
        cases = [
            (new_list(1, 2, 3, 4, 5), new_list(1, 3, 5, 2, 4)),
            (new_list(2, 1, 3, 5, 6, 4, 7), new_list(2, 3, 6, 7, 1, 5, 4))
        ]
        for head, expected in cases:
            output = odd_even_list(head)
            self.assertEqual(output, expected)
