from collections import deque
from typing import Optional, List


class TreeNode:
    def __init__(self, val: int):
        self.val: int = val
        self.left: Optional[TreeNode] = None
        self.right: Optional[TreeNode] = None

    def __eq__(self, other: object) -> bool:
        return (
            isinstance(other, TreeNode)
            and self.val == other.val
            and self.left == other.left
            and self.right == other.right
        )

    def __str__(self) -> str:
        return f"({self.val} {self.left} {self.right})"

    def __repr__(self) -> str:
        return f"TreeNode(val={self.val}, left={self.left}, right={self.right})"


def new_tree(*nums: Optional[int]) -> Optional[TreeNode]:
    if not nums:
        return None

    root = TreeNode(nums[0])
    queue = deque()
    queue.append(root)

    i = 1
    while i < len(nums):
        node = queue.popleft()
        if nums[i] is not None:
            node.left = TreeNode(nums[i])
            queue.append(node.left)
        i += 1
        if i < len(nums):
            if nums[i] is not None:
                node.right = TreeNode(nums[i])
                queue.append(node.right)
            i += 1
    return root


def is_valid_bst(root: TreeNode) -> bool:
    if not root:
        return True
    if root.left and root.left.val > root.val:
        return False
    if root.right and root.right.val < root.val:
        return False
    return is_valid_bst(root.left) and is_valid_bst(root.right)


def height(root: TreeNode) -> int:
    if not root:
        return 0
    left_height = height(root.left)
    right_height = height(root.right)
    root.height = max(left_height, right_height) + 1
    return root.height


def is_valid_avl(root: TreeNode) -> bool:
    if not root:
        return True
    if not is_valid_bst(root):
        return False
    return abs(height(root.left) - height(root.right)) <= 1


def inorder_traverse(root: TreeNode) -> List[int]:
    result = []

    def inorder(node: TreeNode) -> None:
        if node:
            inorder(node.left)
            result.append(node.val)
            inorder(node.right)

    inorder(root)
    return result


def preorder_traverse(root: TreeNode) -> List[int]:
    result = []

    def preorder(node: TreeNode) -> None:
        if node:
            result.append(node.val)
            preorder(node.left)
            preorder(node.right)

    preorder(root)
    return result


def postorder_traverse(root: TreeNode) -> List[int]:
    result = []

    def postorder(node: TreeNode) -> None:
        if node:
            postorder(node.left)
            postorder(node.right)
            result.append(node.val)

    postorder(root)
    return result


def level_order_traverse(root: TreeNode) -> List[int]:
    queue, result = deque(), []
    if root:
        queue.append(root)

    while queue:
        n = len(queue)
        for _ in range(n):
            node = queue.popleft()
            result.append(node.val)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)

    return result
