


def fibonacci(n):
    memo = [0] * (n + 1)
    if n == 0 or n == 1:
        return n
    memo[1] = 1

    for i in range(2, n + 1):
        memo[i] = memo[i - 2] + memo[i - 1]
        print(memo)
    return memo[-1]



print(fibonacci(10))
