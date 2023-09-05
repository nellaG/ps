# https://www.acmicpc.net/problem/14002

n = int(input())
a = list(map(int, input().split()))
d = [1] * n
v = [-1] * n  # indices


def solution():
    for i in range(n):
        for j in range(i):
            if a[j] < a[i] and d[j] + 1 > d[i]:
                d[i] = d[j] + 1
                v[i] = j

    return max(d)
length = solution()
print(length)
idx = d.index(length)
ans = []
while idx >= 0:
    ans.append(a[idx])
    idx = v[idx]
ans.reverse()
print(' '.join(list(map(str, ans))))
