# https://www.acmicpc.net/problem/11724
# 연결 요소

import sys

sys.setrecursionlimit(100000)

n, m = map(int, input().split())
a = [[] for _ in range(n)]
check = [False] * n


def dfs(x):
    check[x] = True
    for y in a[x]:
        if not check[y]:
            dfs(y)


for _ in range(m):
    u, v = map(int, input().split())
    a[u - 1].append(v - 1)
    a[v - 1].append(u - 1)


def solution():
    ans = 0
    for i in range(n):
        if not check[i]:
            dfs(i)
            ans += 1

    return ans


print(solution())
