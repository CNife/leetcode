import unittest
from typing import Optional

from tree import TreeNode


def is_valid_bst(root: TreeNode) -> bool:
    def is_valid_bst_bounded(node: TreeNode, lower: Optional[int], higher: Optional[int]) -> bool:
        if not node:
            return True
        if lower is not None and node.val <= lower:
            return False
        if higher is not None and node.val >= higher:
            return False
        left_is_valid = (not node.left) or is_valid_bst_bounded(node.left, lower, node.val)
        right_is_valid = (not node.right) or is_valid_bst_bounded(node.right, node.val, higher)
        return left_is_valid and right_is_valid

    return is_valid_bst_bounded(root, None, None)


class Test(unittest.TestCase):
    def test_valid_bst_1(self):
        root = TreeNode(2)
        root.left = TreeNode(1)
        root.right = TreeNode(3)
        self.assertTrue(is_valid_bst(root))

    def test_invalid_bst_2(self):
        root = TreeNode(5)
        root.left = TreeNode(1)
        root.right = TreeNode(4)
        root.right.left = TreeNode(3)
        root.right.right = TreeNode(6)
        self.assertFalse(is_valid_bst(root))

    def test_invalid_bst_3(self):
        root = TreeNode(10)
        root.left = TreeNode(5)
        root.right = TreeNode(15)
        root.right.left = TreeNode(6)
        root.right.right = TreeNode(20)
        self.assertFalse(is_valid_bst(root))

    def test_invalid_bst_4(self):
        root = TreeNode(0)
        root.right = TreeNode(-1)
        self.assertFalse(is_valid_bst(root))
