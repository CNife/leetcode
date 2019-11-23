from typing import List, Optional
from unittest import TestCase

from tree import TreeNode


def build_tree(preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
    if not inorder:
        return None
    mid = inorder.index(preorder[0])
    root = TreeNode(preorder[0])
    if len(inorder) > 1:
        root.left = build_tree(preorder[1:mid + 1], inorder[:mid])
        root.right = build_tree(preorder[mid + 1:], inorder[mid + 1:])
    return root


class Test(TestCase):
    def test_build_tree(self):
        root = TreeNode(3)
        root.left = TreeNode(9)
        root.right = TreeNode(20)
        root.right.left = TreeNode(15)
        root.right.right = TreeNode(7)

        preorder = [3, 9, 20, 15, 7]
        inorder = [9, 3, 15, 20, 7]
        output = build_tree(preorder, inorder)
        self.assertEqual(output, root)
