import sys


def test(function, test_cases, get_func=None, map_func=None, eq_func=None):
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
            exit(1)

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
