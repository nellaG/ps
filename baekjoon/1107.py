
# https://www.acmicpc.net/problem/1107 강의


channel = int(input())
num_broken = int(input())
if num_broken > 0:
    brokens = list(map(int, input().split(' ')))
else:
    brokens = []
broken = [False] * 10
for x in brokens:
    broken[x] = True   # marking broken number

# bruteforce
def possible(c):
    if c == 0:
        if broken[0]:
            return 0
        return 1

    len_ = 0
    while c > 0:
        if broken[c % 10]:
            return 0
        len_ += 1
        c //= 10
    return len_


def solution():
    ans = abs(channel - 100)

    for i in range(1000000 + 1):
        c = i
        len_ = possible(c)
        if len_ > 0:
            press = abs(c - channel)
            if ans > len_ + press:
                ans = len_ + press

    return ans

print(solution())