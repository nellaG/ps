# https://www.acmicpc.net/problem/10844


# 계단 수
# d[i][j] = d[i-1][j-1] + d[i-1][j+1]
#  8 >= j >= 1
# j = 0 ) d[i][j] = d[i-1][j+1]
# j = 9) d[i][j] = d[i -1][j -1]


n = int(input())

d = [[0] * 10 for _ in range(n + 1)]
mod = 1_000_000_000

def solution():
    for i in range(0, 10):
        d[1][i] = 1
    d[1][0] = 0
    for i in range(2, n + 1):
        # j = 0
        d[i][0] = d[i - 1][1]
        # j = 9
        d[i][9] = d[i - 1][8]
        for j in range(1, 9):
            d[i][j] = (d[i - 1][j - 1] + d[i - 1][j + 1])

    return sum(d[n]) % mod

print(solution())


