# https://www.acmicpc.net/problem/1182


def solution():
    n, s = list(map(int, input().split()))
    inputs = list(map(int, input().split()))
    ans = 0
    for i in range(1, (1 << n)):  # 2^n - 1
        sum_ = 0
        for k in range(n):
            if (i & (1 << k)) > 0:  # if k is in i
                sum_ += inputs[k]
        if sum_ == s:
            ans += 1

    return ans

print(solution())
