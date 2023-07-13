# https://www.acmicpc.net/problem/1748


def solution():
    n = int(input())
    start = 1
    len_ = 1
    ans = 0
    while start <= n:
        end = start * 10 - 1
        if end > n:
            end = n
        ans += (end - start + 1) * len_
        len_ += 1
        start *= 10
    return ans

print(solution())

