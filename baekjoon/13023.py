# https://www.acmicpc.net/problem/13023
# 그래프 문제
import sys

n, m = map(int, input().split())
edges = []
a = [[False] * n for _ in range(n)]  # 인접 행렬
g = [[] for _ in range(n)]  # 간선 리스트 (양방향)

def solution(m):
    for _ in range(m):
        u, v = map(int, input().split())
        edges.append((u, v))
        edges.append((v, u))
        a[u][v] = a[v][u] = True
        g[u].append(v)
        g[v].append(u)

    m *= 2

    for i in range(m):
        for j in range(m):
            # A -> B, C -> D 간선 찾기
            A, B = edges[i]
            C, D = edges[j]
            if A == B or A == C or A == D or B == C or B == D or C == D:
                continue

            # B -> C 간선 찾기
            if not a[B][C]:
                continue

            for E in g[D]: # 인접 리스트에서 모든 E
                if A == E or B == E or C == E or D == E:
                    continue
                print(1)
                sys.exit(0)


solution(m)
print(0)
