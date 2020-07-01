import sys


def test(function, test_cases, get_func=None, map_func=None, eq_func=None):
    for test_case in test_cases:
        args, expect = test_case[:-1], test_case[-1]
        actual = function(*args)
        if get_func:
            actual = get_func(test_case)
        if map_func:
            actual, expect = map_func(actual), map_func(expect)
        if (eq_func and not eq_func(actual, expect)) or actual != expect:
            print(f'Test failed\n'
                  f'\tfunction: {function!r}\n'
                  f'\targs: {", ".join(map(lambda arg: str(arg), args))}\n'
                  f'\tactual: {actual}\n'
                  f'\texpect: {expect}', file=sys.stderr)
            exit(1)
    print(f'Test for {function.__name__} passed')
