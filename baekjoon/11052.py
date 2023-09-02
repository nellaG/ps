# https://www.acmicpc.net/problem/11052

n = int(input())
prices = list(map(int, input().split()))
total = [0] * (n + 1)


def solution():
    for i in range(1, n + 1):
        for j in range(1, i + 1):
            total[i] = max(total[i], total[i - j] + prices[j - 1])
    return total[-1]


print(solution())
