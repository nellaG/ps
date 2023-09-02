# https://www.acmicpc.net/problem/1463

# avoid using recursion with python (recursion limit)
# use bottom-up instead

n = int(input())

memo = [0] * (n + 1)
memo[1] = 0


def solution(n: int):
    for i in range(2, n + 1):
        memo[i] = memo[i - 1] + 1  # temporary minimum
        if i % 3 == 0 and memo[i] > memo[i // 3] + 1:
            memo[i] = memo[i // 3] + 1
        elif i % 2 == 0 and memo[i] > memo[i // 2] + 1:
            memo[i] = memo[i // 2] + 1


solution(n)

print(memo[n])
