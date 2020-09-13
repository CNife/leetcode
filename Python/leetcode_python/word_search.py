from typing import List

from leetcode_python import test


def exist(board: List[List[str]], word: str) -> bool:
    pass


# noinspection SpellCheckingInspection
test(
    exist,
    [
        ([['A', 'B', 'C', 'E'], ['S', 'F', 'C', 'S'], ['A', 'D', 'E', 'E']],
         "ABCCED", True),
        ([['A', 'B', 'C', 'E'], ['S', 'F', 'C', 'S'], ['A', 'D', 'E', 'E']],
         "SEE", True),
        ([['A', 'B', 'C', 'E'], ['S', 'F', 'C', 'S'], ['A', 'D', 'E', 'E']],
         "ABCB", False),
    ]
)
