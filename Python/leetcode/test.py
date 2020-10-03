from typing import Callable, List, Tuple, TypeVar

from rich.console import Console
from rich.panel import Panel

console = Console()


def test(
    function: Callable,
    test_cases: List[Tuple],
    *,
    args_func: Callable = None,
    expect_func: Callable = None,
    actual_func: Callable = None,
    map_func: Callable = None,
    equals_func: Callable = None,
):
    if args_func is None:

        def args_func(case):
            return case[:-1]

    if expect_func is None:

        def expect_func(case):
            return case[-1]

    if actual_func is None:

        def actual_func(_, prev_actual):
            return prev_actual

    if map_func is None:

        def map_func(item):
            return item

    if equals_func is None:

        def equals_func(lhs, rhs):
            return lhs == rhs

    for test_case in test_cases:
        args = args_func(test_case)
        expect = expect_func(test_case)
        actual = actual_func(test_case, function(*args))

        expect, actual = map_func(expect), map_func(actual)
        passed = equals_func(actual, expect)

        if not passed:
            console.print(f"Test for [blue]{function}[/] [red]Failed[/]!")
            console.print(Panel(str(args), title="args"))
            console.print(Panel(str(expect), title="expect"))
            console.print(Panel(str(actual), title="actual"))
            exit(1)

    console.print(f"Test for [blue]{function}[/] [green]Succeed[/]!")


T = TypeVar("T")


def sorted_list(src: List[T], **kwargs) -> List[T]:
    src.sort(**kwargs)
    return src


def sorted_2d_list(src: List[List[T]], **kwargs) -> List[List[T]]:
    for ll in src:
        ll.sort()
    src.sort(**kwargs)
    return src
