# https://www.acmicpc.net/problem/11054

n = int(input())
a = list(map(int, input().split()))

d = [1] * n  # 증가
d2 = [1] * n  # 감소
d3 = [0] * n


def solution():
    # 증가
    for i in range(n):
        for j in range(i):
            if a[j] < a[i] and d[j] + 1 > d[i]:
                d[i] = d[j] + 1

    # 감소
    for i in range(n - 1, -1, -1):
        d2[i] = 1
        for j in range(i, n):
            if a[i] > a[j] and d2[i] < d2[j] + 1:
                d2[i] = d2[j] + 1

    for i in range(n):
        d3[i] = d[i] + d2[i] - 1

    return max(d3)


print(solution())
