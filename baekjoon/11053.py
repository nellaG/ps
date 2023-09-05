# https://www.acmicpc.net/problem/11053

n = int(input())

a = list(map(int, input().split()))

d = [1] * n

def solution():
    for i in range(n):
        for j in range(i):
            if a[j] < a[i] and d[j] + 1 > d[i]:
                d[i] = d[j] + 1

    return max(d)

print(solution())
