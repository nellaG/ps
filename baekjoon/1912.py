# https://www.acmicpc.net/problem/1912

# d[i] = i번째 수로 끝나는 가장 큰 연속합
# d[i] = d[i-1] + a[i] # a[i-1]과 연속하는경우
# d[i] = a[i] 일수도 .이때는 불연속. 그럼 둘중에 max

n = int(input())
a = list(map(int, input().split()))
d = [0] * n
d[0] = a[0]


def solution():
    for i in range(1, n):
        d[i] = a[i]
        d[i] = max(a[i], d[i-1] + a[i])

    return max(d)


print(solution())
