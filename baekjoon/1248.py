# https://www.acmicpc.net/problem/1248

n = int(input())
s = input()
sign = [[0] * n for _ in range(n)]
ans = [0] * n
cnt = 0


def go(index):
    if index == n:
        return True

    if sign[index][index] == 0:
        ans[index] = 0
        return check(index) and go(index + 1)

    for i in range(1, 11):
        ans[index] = sign[index][index] * i
        if check(index) and go(index + 1):
            return True
    return False


def check(index: int):
    _sum = 0
    for i in range(index, -1, -1):
        _sum += ans[i]
        if sign[i][index] == 0:
            if _sum != 0:
                return False
        elif sign[i][index] > 0:
            if _sum <= 0:
                return False
        elif sign[i][index] < 0:
            if _sum >= 0:
                return False

    return True


for i in range(n):
    for j in range(i, n):
        if s[cnt] == '0':
            sign[i][j] = 0
        elif s[cnt] == '+':
            sign[i][j] = 1
        else:
            sign[i][j] = -1
        cnt += 1


go(0)
print(ans)
