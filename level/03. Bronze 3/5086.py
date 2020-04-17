import sys

while True:
    a, b = map(int, sys.stdin.readline().split())
    if a == b == 0:
        break
    if b % a == 0:
        print("factor")
    elif a % b == 0:
        print("multiple")
    else:
        print("neither")
