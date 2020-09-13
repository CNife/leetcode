import sys

from typing import Callable, List, Tuple, TypeVar


def test(
    function: Callable,
    test_cases: List[Tuple],
    *,
    actual_func: Callable = None,
    map_func: Callable = None,
    eq_func: Callable = None,
):
    for test_case in test_cases:
        args, expect = test_case[:-1], test_case[-1]
        actual = function(*args)

        def error():
            message = (
                f"Test failed\n"
                f"function: {function}\n"
                f'args: {", ".join(map(lambda arg: str(arg), args))}\n'
                f"actual: {actual}\n"
                f"expect: {expect}\n"
            )
            print(message, file=sys.stderr)
            sys.exit(1)

        if actual_func:
            actual = actual_func(test_case)
        if map_func:
            actual, expect = map_func(actual), map_func(expect)
        if eq_func:
            if not eq_func(actual, expect):
                error()
        elif actual != expect:
            error()

    print(f"Test for {function.__name__} passed")


T = TypeVar("T")


def sorted_list(src: List[T], **kwargs) -> List[T]:
    src.sort(**kwargs)
    return src
