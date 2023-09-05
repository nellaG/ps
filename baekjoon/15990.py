# https://www.acmicpc.net/problem/15990


t = int(input())
limit = 100000

d = [[0] * 4 for _ in range(limit + 1)]

mod = 1000000009

def solution():
    for i in range(1, limit + 1):
        d[1][1] = 1
        d[2][2] = 1
        d[3][3] = 1
        if i - 1 >= 0:
            d[i][1] = d[i-1][2] + d[i-1][3]
        if i - 2 >= 0:
            d[i][2] = d[i-2][1] + d[i-2][3]
        if i - 3 >= 0:
            d[i][3] = d[i-3][1] + d[i-3][2]

        d[i][1] %= mod
        d[i][2] %= mod
        d[i][3] %= mod


solution()
for _ in range(t):
    n = int(input())
    print(sum(d[n]) % mod)
