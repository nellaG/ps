# https://www.acmicpc.net/problem/3085

def solution():
    num = int(input())
    # 행 체크 한번, 열 체크 한번
    a = [list(input()) for _ in range(num)]

    answer = 0
    for i in range(num):
        for j in range(num):
            # 행
            if j + 1 < num:
                a[i][j], a[i][j+1] = a[i][j+1], a[i][j] # swap
                temp = check(a, i, i, j, j + 1)
                if answer < temp:
                    answer = temp
                a[i][j], a[i][j+1] = a[i][j+1], a[i][j] # reswap
            # 열
            if i + 1 < num:
                a[i][j], a[i+1][j] = a[i+1][j], a[i][j] # swap
                temp = check(a, i, i+1, j, j)
                if answer < temp:
                    answer = temp
                a[i][j], a[i+1][j] = a[i+1][j], a[i][j] # reswap
    return answer


def check(a, start_row, end_row, start_col, end_col):
    n = len(a)
    ans = 1
    for i in range(start_row, end_row + 1):
        cnt = 1
        for j in range(1, n):
            if a[i][j] == a[i][j - 1]:
                cnt += 1  # 같은 색상
            else:
                cnt = 1
            if ans < cnt:
                ans = cnt
    for i in range(start_col, end_col + 1):
        cnt = 1
        for j in range(1, n):
            if a[j][i] == a[j - 1][i]:
                cnt += 1
            else:
                cnt = 1
            if ans < cnt:
                ans = cnt
    return ans


print(solution())
