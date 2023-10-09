# https://www.acmicpc.net/problem/2133
# 3 * N 의 타일 채우기

n = int(input())
d = [0] * (n + 1)
d[0] = 1
# d[i] = 3 * i 를 채우는 방법수.
# d[i] = 3 * d[i-2] + 2 * d[i-4] + 2 * d[i-6] + ...

def solution():
    for i in range(2, n + 1, 2):
        d[i] = 3 * d[i - 2]
        for j in range(i - 4, -1, -2):
            d[i] +=  2 * d[j]

    return d[-1]


print(solution())

