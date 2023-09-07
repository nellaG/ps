# https://www.acmicpc.net/problem/1149

# d[i][j] = i번 집을 j번 색으로 칠하려고 할때 1~i번 집을 칠하는 비용의 최소
# a[i][j]  i번 집을 j 번 색으로 칠하는 비용
# d[i][0] =  min(d[i-1][1], d[i-1][2]) + a[i][0]

n = int(input())
a = [list(map(int, input().split())) for _ in range(n)]
d = [[0] * 3 for _ in range(n)]  # not [[0] * 3] * n (reference)

def solution():
    d[0][0] = a[0][0]
    d[0][1] = a[0][1]
    d[0][2] = a[0][2]
    for i in range(1, n):
        d[i][0] = min(d[i - 1][1], d[i - 1][2]) + a[i][0]
        d[i][1] = min(d[i - 1][2], d[i - 1][0]) + a[i][1]
        d[i][2] = min(d[i - 1][1], d[i - 1][0]) + a[i][2]

    return min(d[n - 1][0], d[n - 1][1], d[n - 1][2])


print(solution())
