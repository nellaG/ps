# https://www.acmicpc.net/problem/2156


# a[i] = i번 포도주의 양
# d[i][j] = i번째있고 i번 포도주는 j번 연속해서 마신 포도주
# 0번 연속 -> 안 마심, 1번 연속 -> 직전 포도주를 안 마심
# 2번 연속 -> 직전 포도주를 마시고, 그 전의 포도주는 안 마심

n = int(input())

a = [int(input()) for _ in range(n)]
d = [[0] * 3 for _ in range(n + 1)]

def solution():

    for i in range(1, n + 1):
        d[i][0] = max(d[i-1][0], d[i-1][1], d[i-1][2])
        d[i][1] = d[i-1][0] + a[i-1]
        d[i][2] = d[i-1][1] + a[i-1]
    return max(d[n])


print(solution())
