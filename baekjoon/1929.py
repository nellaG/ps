# https://www.acmicpc.net/problem/1929


def solution():
    floor, ceil = list(map(int, input().split(' ')))
    check = [0] * (ceil + 1)
    check[0] = check[1] = 1  # 1: erased

    # erasing
    for i in range(2, ceil + 1):
        if check[i] == 0:
            j = i + i
            while j <= ceil:
                check[j] = 1
                j += i

    for i in range(floor, ceil + 1):
        if check[i] == 0:
            print(i)


solution()
