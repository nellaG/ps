# https://www.acmicpc.net/problem/15652

# count; 사용한 수의 개수
t = int(input())

values = [int(input()) for _ in range(t)]


# 호출시점에서 사용한 각 수의 개수, 합, 수
def solution(count, sum_, n):
    if sum_ == n:
        return 1
    elif sum_ > n:
        return 0
    now = 0
    now += solution(count + 1, sum_ + 1, n)
    now += solution(count + 1, sum_ + 2, n)
    now += solution(count + 1, sum_ + 3, n)

    return now


for v in values:
    print(solution(0, 0, v))
