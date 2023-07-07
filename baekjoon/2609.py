# https://www.acmicpc.net/problem/2609

def solution():
    x, y = map(int, input().split(' '))

    gcd_val = gcd(x, y)
    lcm_val = int(x * y / gcd_val)

    return f'{gcd_val}\n{lcm_val}'


def gcd(x, y):
    if y == 0:
        return x
    return gcd(y, x % y)

print(solution())
