


def sum_rec(i: int):
    if not isinstance(i, int):
        raise Exception(f'must input int value: {i}')
    if i < 0:
        raise Exception(f'must input int value upper 0: {i}')
    if 0 <= i <= 1:
        return i
    else:
        return sum_rec(i - 1) + i


def gcd(x, y):
    print(f'gcd({y}, {x} % {y})')
    if y == 0:
        return x
    else:
        return gcd(y, x % y)
