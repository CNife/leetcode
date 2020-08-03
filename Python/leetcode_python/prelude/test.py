import sys

from typing import Callable, List, Tuple


def test(
        function: Callable,
        test_cases: List[Tuple],
        get_func: Callable = None,
        map_func: Callable = None,
        eq_func: Callable = None
):
    for test_case in test_cases:
        args, expect = test_case[:-1], test_case[-1]
        actual = function(*args)

        def error():
            message = (f'Test failed\n'
                       f'function: {function}\n'
                       f'args: {", ".join(map(lambda arg: str(arg), args))}\n'
                       f'actual: {actual}\n'
                       f'expect: {expect}\n')
            print(message, file=sys.stderr)
            sys.exit(1)

        if get_func:
            actual = get_func(test_case)
        if map_func:
            actual, expect = map_func(actual), map_func(expect)
        if eq_func:
            if not eq_func(actual, expect):
                error()
        elif actual != expect:
            error()

    print(f'Test for {function.__name__} passed')
