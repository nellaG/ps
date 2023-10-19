# https://www.acmicpc.net/problem/1707
# 이분 그래프


import sys
sys.setrecursionlimit(1_000_000)
k = int(sys.stdin.readline())


def solution():
    n, m = map(int, sys.stdin.readline().split())
    a = [[] for _ in range(n)]
    color = [0] * n  # 0(방문x), 1(a), 2(b)
    for _ in range(m):
        u, v = map(int, sys.stdin.readline().split())
        a[u - 1].append(v - 1)
        a[v - 1].append(u - 1)

    def dfs(x, c):
        color[x] = c
        for y in a[x]:
            if color[y] == 0:
                if not dfs(y, 3 - c):
                    return False
            elif color[y] == color[x]:
                return False
        return True

    ans = True
    for i in range(n):
        if color[i] == 0:
            if not dfs(i, 1):
                ans = False

    print ('YES' if ans else 'NO')


for _ in range(k):
    solution()


