# https://www.acmicpc.net/problem/16194

n = int(input())
prices = list(map(int, input().split()))
total = [100_000_000] * (n + 1)
total[0] = 0
total[1] = prices[0]  # initial


def solution():
    for i in range(1, n + 1):
        for j in range(1, i + 1):
            total[i] = min(total[i], total[i - j] + prices[j - 1])
    return total[-1]


print(solution())
