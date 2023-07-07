# https://www.acmicpc.net/problem/1978


def solution():
    _ = int(input())
    xs = list(map(int, input().split(' ')))
    answers = 0
    for x in xs:
        answers += (1 if prime(x) else 0)

    return answers


def prime(x: int) -> bool:
    if x < 2:
        return False
    i = 2
    while i ** 2 <= x:
        if x % i == 0:
            return False
        i += 1

    return True


print(solution())
