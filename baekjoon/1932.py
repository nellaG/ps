# https://www.acmicpc.net/problem/1932

# d[i][j] = (i, j)에 도착했을때 합의 최대값
#    (i+1, j),  (i+1, j+1)
# d[i][j] = max(d[i-1][j], d[i-1][j-1]) + a[i][j]
# !값들을 배열 중앙에 정렬하지 않아도 됨
n = int(input())
a = [list(map(int, input().split())) for _ in range(n)]
d = [[0] * n for _ in range(n)]

d[0][0] = a[0][0]

def solution():
    for i in range(1, n):
        for j in range(0, i + 1):
            d[i][j] = d[i-1][j] + a[i][j]
            if j - 1 >= 0:
                d[i][j] = max(d[i][j], d[i-1][j-1] + a[i][j])

    return max(d[n - 1])


print(solution())
