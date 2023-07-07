# https://www.acmicpc.net/problem/17427

def solution():
    num = int(input())
    result = 0
    for i in range(1, num + 1):
        result += (num // i) * i
    return result


print(solution())
