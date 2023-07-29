# https://www.acmicpc.net/problem/2529
N = int(input())

signs = input().split()
check = [False] * 10
answers = []
def go(index, num):
    if index == N + 1:
        answers.append(num)
        return

    for i in range(10):
        if check[i]:
            continue
        if index != 0 and (not good(num[index - 1], str(i), signs[index - 1])):
            continue
        check[i] = True
        go(index + 1, num + str(i))
        check[i] = False


def ok(num: str):
    for i in range(N):
        if signs[i] == '<':
            if num[i] > num[i + 1]:
                return False
        else:
            if num[i] < num[i + 1]:
                return False

    return True

def good(x, y, op):
    return (op == '<' and x < y) or (op == '>' and x > y)

go(0, '')

print(max(answers))
print(min(answers))
