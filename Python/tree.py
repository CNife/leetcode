from typing import Optional


class TreeNode:
    def __init__(self, val: int):
        self.val: int = val
        self.left: Optional[TreeNode] = None
        self.right: Optional[TreeNode] = None

    def __eq__(self, other: object) -> bool:
        if self is other:
            return True
        return (isinstance(other, TreeNode)
                and self.val == other.val
                and self.left == other.left
                and self.right == other.right)
