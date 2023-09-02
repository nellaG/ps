# https://www.acmicpc.net/problem/11726


n = int(input())

memo = [0] * (n + 1)
memo[0] = 1
memo[1] = 1  # initial
# d[n] = d[n-1] + d[n-2]
def solution():
    for i in range(2, n + 1):
        memo[i] = memo[i - 1] + memo[i - 2]

    return memo[n]

print(solution() % 10007)
