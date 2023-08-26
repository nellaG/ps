# https://www.acmicpc.net/problem/1463

# avoid using recursion with python (recursion limit)
# use bottom-up instead

n = int(input())
ans = 0

memo = [0] * 1000001

def solution(n: int):
    if n == 1:
        return 0
    if memo[n] > 0:
        return memo[n]
    memo[n] = solution(n - 1) + 1  # temporary min

    if n % 3 == 0:
        temp = solution(n // 3) + 1
        if memo[n] > temp:
            memo[n] = temp
    elif n % 2 == 0:
        temp = solution(n // 2) + 1
        if memo[n] > temp:
            memo[n] = temp
    return memo[n]


print(solution(n))
