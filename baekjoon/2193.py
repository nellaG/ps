# https://www.acmicpc.net/problem/2193


n = int(input())

# d[i][0] = d[i-1][0] + d[i-1][1]
# d[i][1] = d[i-1][0]
# d[0][j] = 0
# d[1][0] = 0
# d[1][1] = 1
d = [[0] * 2 for _ in range(n + 1)]

def solution():
    d[1][1] = 1

    for i in range(2, n + 1):
        d[i][0] = d[i-1][0] + d[i-1][1]
        d[i][1] = d[i-1][0]

    return sum(d[n])


print(solution())
