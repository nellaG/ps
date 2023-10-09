# https://www.acmicpc.net/problem/10845
# use sys.stdin.readline (input() is slow)
# also use sys.stdout.write

import sys

n = int(sys.stdin.readline())
queue = []
begin = 0
while n > 0:
    command, *x = sys.stdin.readline().split()
    if x != []:
        x = int(x[0])
        if command == 'push':
            queue.append(x)
    else:
        if command == 'pop':
            if begin >= len(queue):
                print(-1)
            else:
                print(queue[begin])
                queue[begin] = 0
                begin += 1
        elif command == 'size':
            print(len(queue) - begin)
        elif command == 'empty':
            print(int(begin >= len(queue)))
        elif command == 'front':
            if begin >= len(queue):
                print(-1)
            else:
                print(queue[begin])
        elif command == 'back':
            if begin >= len(queue):
                print(-1)
            else:
                print(queue[-1])
    n -= 1
