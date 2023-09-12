# https://www.acmicpc.net/problem/2565
n = int(input())

a = [list(map(int, input().split())) for _ in range(n)]
# d[i] = i번째까지 왔을떄 교차 안하는 (증가하는) 길이
d = [1] * n
a = sorted(a)

def solution():
    for i in range(n):
        for j in range(i): # j < i
            if a[j][1] < a[i][1]:
                if d[j] + 1 > d[i]:
                    d[i] = d[j] + 1
    return n - max(d)


print(solution())
