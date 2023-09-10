# https://www.acmicpc.net/problem/11722


# d[i] = a[i] 에서 끝나는 가장 긴 감소 부분수열 길이
# j < i, a[j] > a[i]
# ...

# d[i] = a[i] 에서'시작하는' 가장 긴 감소하는 부분수열의 길이
# i < j and a[i] > a[j], d[i] = max(d[j]) + 1
# N -> 1 로 반대로 순회
n = int(input())
a = list(map(int, input().split()))
d = [0] * n


def solution():
    for i in range(n - 1, -1, -1):
        d[i] = 1
        for j in range(i + 1, n):
            if a[i] > a[j] and d[i] < d[j] + 1:
                d[i] = d[j] + 1
    return max(d)


print(solution())
