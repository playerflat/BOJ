import sys

for _ in range(int(sys.stdin.readline())):
    r = 0
    for _ in range(int(sys.stdin.readline())):
        a = max(map(int, sys.stdin.readline().split()))
        if a > 0:
            r += a
    print(r)
