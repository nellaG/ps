# https://www.acmicpc.net/problem/14501

N = int(input())
pair = [list(map(int, (input().split()))) for _ in range(N)]
answer = 0


def go(index, money):

    if (index == N):
        global answer
        if money > answer:
            answer = money
        return
    # exceeds day limit
    if (index > N):
        return 0
    day_cost = pair[index][0]
    pay = pair[index][1]

    # 상담을 한다면
    go(index + day_cost, money + pay)
    # 상담을 안한다면
    go(index + 1, money)


go(0, 0)
print(answer)
