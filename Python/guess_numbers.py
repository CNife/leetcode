import unittest
from typing import List


def game(guess: List[int], answer: List[int]) -> int:
    return len([None for g, a in zip(guess, answer) if g == a])


class Test(unittest.TestCase):
    def test_game(self):
        cases = [
            ([1, 2, 3], [1, 2, 3], 3),
            ([2, 2, 3], [3, 2, 1], 1)
        ]
        for guess, answer, expected in cases:
            output = game(guess, answer)
            self.assertEqual(output, expected)
