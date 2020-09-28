import functools
import itertools
import math
from typing import List, Tuple

from leetcode_python import test


def judge_24_points(nums: List[int]) -> bool:
    @functools.lru_cache(None)
    def inner(t: Tuple[int, ...]) -> bool:
        if len(t) == 1:
            return math.isclose(t[0], 24)
        return any(
            inner(tuple(rest) + (x,))
            for a, b, *rest in itertools.permutations(t)
            for x in {a + b, a - b, a * b, b and a / b}
        )

    return inner(tuple(nums))


test(
    judge_24_points,
    [
        ([4, 1, 8, 7], True),
        ([1, 2, 1, 2], False),
    ],
)
