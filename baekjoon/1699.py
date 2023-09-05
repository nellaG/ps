# https://www.acmicpc.net/problem/1699

# d[n] = min(d[N - i2]) + 1, 1 <= i^2 <= N 필요한 항의 최소 개수


n = int(input())

d = [0] * (n + 1)


def solution():
    for i in range(n + 1):
        d[i] = i
        j = 1
        while j * j <= i:
            if d[i] > d[i - j*j] + 1:  # find min
                d[i] = d[i - j*j] + 1
            j += 1
    return d[-1]


print(solution())
