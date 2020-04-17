import sys

for _ in range(int(sys.stdin.readline())):
    a = b = 0
    for _ in range(9):
        c, d = map(int, sys.stdin.readline().split())
        a += c
        b += d
    if a > b:
        print("Yonsei")
    elif a < b:
        print("Korea")
    else:
        print("Draw")
