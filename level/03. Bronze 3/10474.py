import sys

while 1:
    a, b = map(int, sys.stdin.readline().split())
    if a == b == 0:
        break
    print(f'{a // b} {a % b} / {b}')
