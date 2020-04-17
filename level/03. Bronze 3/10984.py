import sys

for _ in range(int(sys.stdin.readline())):
    c = 0
    r = 0
    for _ in range(int(sys.stdin.readline())):
        a, b = map(float, sys.stdin.readline().split())
        c += a * b
        r += a
    print(int(r), round(c / r, 1))
