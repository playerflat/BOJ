import math
import sys

count = 1
while 1:
    a, b, c = map(int, sys.stdin.readline().split())
    if a == b == c == 0:
        break
    print(f'Triangle #{count}')

    if c == -1:
        print(f'c = {math.sqrt(a ** 2 + b ** 2):.3f}')
    elif a >= c or b >= c:
        print('Impossible.')
    elif a == -1:
        print(f'a = {math.sqrt(c ** 2 - b ** 2):.3f}')
    else:
        print(f'b = {math.sqrt(c ** 2 - a ** 2):.3f}')

    print()
    count += 1
