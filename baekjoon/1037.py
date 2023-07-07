# https://www.acmicpc.net/problem/1037


def solution():
    _ = input()
    numbers = input()
    numbers = map(int, numbers.split(' '))
    sorted_numbers = sorted(numbers)

    return sorted_numbers[0] * sorted_numbers[-1]


print(solution())
