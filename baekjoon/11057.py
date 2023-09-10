# https://www.acmicpc.net/problem/11057


n = int(input())

d = [[0] * 10 for _ in range(n + 1)]  # d[i][j] = j로 끝나는 i자리수 오르막수 개수
d[0][0] = 1
mod = 10007

def solution():
    for i in range(1, n + 1):
        for j in range(10):
            for k in range(j + 1):
                d[i][j] += d[i - 1][k]
            d[i][j] %= mod

    return sum(d[n]) % mod


print(solution())
