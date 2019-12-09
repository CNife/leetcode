from typing import List
from unittest import TestCase

from tree import TreeNode


def zigzag_level_order(root: TreeNode) -> List[List[int]]:
    if not root:
        return []

    level = [root]
    result = []
    reverse = False

    while level:
        values = [node.val for node in level]
        if reverse:
            values.reverse()
        reverse = not reverse
        result.append(values)

        tmp = []
        for node in level:
            if node.left:
                tmp.append(node.left)
            if node.right:
                tmp.append(node.right)
        level = tmp

    return result


class Test(TestCase):
    def test_zigzag_level_order(self):
        root = TreeNode(3)
        root.left = TreeNode(9)
        root.right = TreeNode(20)
        root.right.left = TreeNode(15)
        root.right.right = TreeNode(7)
        expected = [[3], [20, 9], [15, 7]]

        output = zigzag_level_order(root)
        self.assertEqual(output, expected)
