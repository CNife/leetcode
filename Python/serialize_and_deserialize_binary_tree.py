from collections import deque
from typing import Optional
from unittest import TestCase

from tree import TreeNode


def serialize(root: TreeNode) -> str:
    result, queue = [], deque([root])
    while queue:
        node = queue.popleft()
        if node is not None:
            queue.extend([node.left, node.right])
            result.append(str(node.val))
        else:
            result.append('null')
    while result and result[-1] == 'null':
        result.pop()
    return ','.join(result)


def deserialize(data: str) -> Optional[TreeNode]:
    nodes = map(lambda t: None if t in ('null', '') else TreeNode(int(t)), data.split(','))
    root = next(nodes)
    if root is None:
        return None

    queue = deque([root])
    is_left = True
    node = None
    for next_node in nodes:
        if next_node:
            queue.append(next_node)
        if is_left:
            node = queue.popleft()
            node.left = next_node
        else:
            node.right = next_node
        is_left = not is_left
    return root


class Test(TestCase):
    data = "1,2,3,null,null,4,5"
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)
    root.right.left = TreeNode(4)
    root.right.right = TreeNode(5)

    def test_codec(self):
        self.assertEqual(deserialize(self.data), self.root)

    def test_serialize(self):
        self.assertEqual(serialize(self.root), self.data)

    def test_serialize_2(self):
        root = TreeNode(1)
        root.left = TreeNode(2)
        root.left.left = TreeNode(3)
        root.left.left.left = TreeNode(4)
        root.left.left.left.left = TreeNode(5)
        self.assertEqual(deserialize(serialize(root)), root)
