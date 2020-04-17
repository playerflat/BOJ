import sys

x = int(sys.stdin.readline())
while True:
    a = int(sys.stdin.readline())
    if a == 0: break
    if a % x == 0:
        print(f'{a} is a multiple of {x}.')
    else:
        print(f'{a} is NOT a multiple of {x}.')
