import unittest

from tree import TreeNode


def rob(root: TreeNode) -> int:
    def helper(node: TreeNode) -> (int, int):
        if not node:
            return 0, 0
        left_max, left_no_rob = helper(node.left)
        right_max, right_no_rob = helper(node.right)
        return max(left_no_rob + right_no_rob + node.val, left_max + right_max), left_max + right_max

    return helper(root)[0]


class Test(unittest.TestCase):
    def test(self):
        tree = TreeNode(3)
        tree.left = TreeNode(4)
        tree.left.left = TreeNode(1)
        tree.left.right = TreeNode(3)
        tree.right = TreeNode(5)
        tree.right.right = TreeNode(1)

        output = rob(tree)
        self.assertEqual(output, 9)
