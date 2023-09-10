# https://www.acmicpc.net/problem/11055

n = int(input())

a = [0] + list(map(int, input().split()))

d = [0] * (n + 1)
d[1] = a[1]
# d[i] = i번째수로 끝나는 부분수열의 합

def solution():
    for i in range(n + 1):
        for j in range(i):
            if a[j] < a[i] and d[j] + a[i] > d[i]:
                d[i] = d[j] + a[i]

    return max(d)


print(solution())
