# https://www.acmicpc.net/problem/11723
# use sys.stdin.readline (input() is slow)
# also use sys.stdout.write

import sys

s = 0
n = int(sys.stdin.readline())
while n > 0:
    command, *x = sys.stdin.readline().split()
    if x != []:
        x = int(x[0]) - 1
        if command == 'add':
            s = (s | (1 << x))
        elif command == 'remove':
            s = (s & ~(1 << x))
        elif command == 'check':
            sys.stdout.write(f'{int(bool((s & (1 << x))))}\n')
        elif command == 'toggle':
            s = (s ^ (1 << x))
    else:
        if command == 'all':
            s = (1 << 20) - 1
        elif command == 'empty':
            s = 0
    n -= 1
