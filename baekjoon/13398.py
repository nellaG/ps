# https://www.acmicpc.net/problem/13398
n = int(input())
a = list(map(int, input().split()))

# 바이토닉이랑 같음
d = [0] * n  # d[i] = i번째에서 끝나는 최대 연속합
d2 = [0] * n # d2[i] = i번째에서 시작하는 최대 연속합. d[i-1] + d2[i + 1]로 계산



def solution():
    for i in range(n):
        d[i] = a[i]
        if i == 0:
            continue
        if d[i] < d[i - 1] + a[i]:
            d[i] = d[i - 1] + a[i]

    for i in range(n - 1, -1, -1):
        d2[i] = a[i]
        if i == n - 1:
            continue
        if d2[i] < d2[i + 1] + a[i]:
            d2[i] = d2[i + 1] + a[i]

    ans = max(d)

    for i in range(1, n - 1):
        ans = max(ans, d[i - 1] + d2[i + 1])

    return ans


print(solution())
