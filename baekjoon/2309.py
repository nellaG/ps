# https://www.acmicpc.net/problem/2309


def solution():
    hobbit = [0] * 9
    for i in range(0, 9):
        height = int(input())
        hobbit[i] = height

    hobbit = pick(hobbit)
    hobbit = sorted(hobbit)
    for hob in hobbit:
        if hob != 0:
            print(hob)


def pick(hobbit):
    total = sum(hobbit)
    subt = total - 100
    for i in range(0, 9):
        for j in range(i + 1, 9):
            if hobbit[i] + hobbit[j] == subt:
                hobbit[i] = 0
                hobbit[j] = 0
                return hobbit
    return hobbit


solution()
