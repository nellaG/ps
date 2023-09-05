# https://www.acmicpc.net/problem/2225


# d[k][n] : k개까지 더해서 n이 되는 경우의 수
# d[k][n] = sigma(d[k-1][n - l]) 0 <= l <= N

mod = 1_000_000_000
n, k = list(map(int, input().split()))
d = [[0] * (n + 1) for _ in range(k + 1)]
d[0][0] = 1


def solution():
    for i in range(k + 1):
        for j in range(n + 1):
            for l in range(j + 1):
                d[i][j] += d[i - 1][j - l]
                d[i][j] %= mod

    return d[k][n]


print(solution())
