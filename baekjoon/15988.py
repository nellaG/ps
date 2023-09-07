# https://www.acmicpc.net/problem/15988


n = int(input())
mod = 1_000_000_009
d = [0] * (1000000 + 1)


def solution(d):
    d[0] = 1
    d[1] = d[0]
    d[2] = d[1] + d[0]
    d[3] = d[2] + d[1] + d[0]
    for i in range(4, 1000001):
        d[i] += d[i - 1] + d[i - 2] + d[i - 3]
        d[i] %= mod

solution(d)

for _ in range(n):
    num = int(input())
    print(d[num])
